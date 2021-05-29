package activemq

import (
	"context"
	"fmt"
	"github.com/go-stomp/stomp"
	jsoniter "github.com/json-iterator/go"
	"github.com/kubemq-hub/builder/connector/common"
	"github.com/kubemq-hub/kubemq-sources/config"
	"github.com/kubemq-hub/kubemq-sources/middleware"
	"github.com/kubemq-hub/kubemq-sources/pkg/logger"
	"github.com/kubemq-hub/kubemq-sources/types"
	"time"
)

const (
	defaultSubTimeout = 5 * time.Second
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Client struct {
	opts   options
	conn   *stomp.Conn
	log    *logger.Logger
	target middleware.Middleware
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
	var options []func(*stomp.Conn) error = []func(*stomp.Conn) error{
		stomp.ConnOpt.Login(c.opts.username, c.opts.password),
		stomp.ConnOpt.Host("/"),
	}
	c.conn, err = stomp.Dial("tcp", c.opts.host, options...)
	if err != nil {
		return fmt.Errorf("error connecting to activemq broker, %w", err)
	}
	return nil
}

func (c *Client) createMetadataString(msg *stomp.Message) string {
	md := map[string]string{}
	md["destination"] = msg.Destination
	md["content_type"] = msg.ContentType
	if msg.Err != nil {
		md["error"] = msg.Err.Error()
	}
	str, err := json.MarshalToString(md)
	if err != nil {
		return fmt.Sprintf("error parsing stomp metadata, %s", err.Error())
	}
	return str
}

func (c *Client) Start(ctx context.Context, target middleware.Middleware) error {
	if target == nil {
		return fmt.Errorf("invalid target received, cannot be nil")
	} else {
		c.target = target
	}
	errCh := make(chan error, 1)
	go func() {
		subscription, err := c.conn.Subscribe(c.opts.destination, stomp.AckAuto)
		if err != nil {
			errCh <- fmt.Errorf("error subscribing to activemq destination, %w", err)
			return
		}
		errCh <- nil
		defer func() {
			_ = subscription.Unsubscribe()
		}()
		for {
			select {
			case msg := <-subscription.C:
				req := types.NewRequest().SetMetadata(c.createMetadataString(msg)).SetData(msg.Body)
				if c.opts.dynamicMapping {
					req.SetChannel(msg.Destination)
				}
				_, err := c.target.Do(ctx, req)
				if err != nil {
					c.log.Errorf("error processing activemq message, %s", err.Error())
				}
			case <-ctx.Done():
				return
			}
		}

	}()

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
		return nil
	case <-time.After(defaultSubTimeout):
		return fmt.Errorf("activemq subscription timeout")
	}
}

func (c *Client) Stop() error {
	return c.conn.Disconnect()
}
