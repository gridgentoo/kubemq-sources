package query

import (
	"context"
	"fmt"
	"github.com/kubemq-hub/builder/connector/common"
	"github.com/kubemq-hub/kubemq-sources/config"
	"github.com/kubemq-hub/kubemq-sources/pkg/logger"
	"github.com/kubemq-hub/kubemq-sources/types"
	"github.com/kubemq-io/kubemq-go"
	"time"
)

type Client struct {
	log    *logger.Logger
	opts   options
	client *kubemq.Client
}

func New() *Client {
	return &Client{}

}
func (c *Client) Connector() *common.Connector {
	return Connector()
}

func (c *Client) Init(ctx context.Context, cfg config.Spec, log *logger.Logger) error {
	c.log = log
	if c.log == nil {
		c.log = logger.NewLogger(cfg.Kind)
	}

	var err error
	c.opts, err = parseOptions(cfg)
	if err != nil {
		return err
	}
	c.client, err = kubemq.NewClient(ctx,
		kubemq.WithAddress(c.opts.host, c.opts.port),
		kubemq.WithClientId(c.opts.clientId),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC),
		kubemq.WithAuthToken(c.opts.authToken),
		// making sure that this stays false in order the http source will work correctly
		kubemq.WithCheckConnection(false))
	if err != nil {
		return err
	}
	return nil
}
func (c *Client) Stop() error {
	if c.client != nil {
		return c.client.Close()
	}
	return nil
}
func (c *Client) getChannel(request *types.Request) string {
	if request.Channel != "" {
		return fmt.Sprintf("%s%s", c.opts.channel, request.Channel)
	}
	return c.opts.channel
}

func (c *Client) Do(ctx context.Context, request *types.Request) (*types.Response, error) {

	queryResponse, err := c.client.Q().
		SetTimeout(time.Duration(c.opts.timeoutSeconds) * time.Second).
		SetChannel(c.getChannel(request)).
		SetMetadata(request.Metadata).
		SetBody(request.Data).
		Send(ctx)
	if err != nil {
		return types.NewResponse().SetError(err), nil
	}
	if !queryResponse.Executed {
		return types.NewResponse().SetError(fmt.Errorf(queryResponse.Error)), nil
	}
	return types.NewResponse().
			SetMetadata(queryResponse.Metadata).
			SetData(queryResponse.Body),
		nil
}
