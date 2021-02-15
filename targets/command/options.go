package command

import (
	"fmt"
	"github.com/kubemq-hub/kubemq-sources/config"
	"github.com/nats-io/nuid"
	"math"
)

const (
	defaultTimeoutSeconds = 600
	defaultAddress        = "localhost:50000"
)

type options struct {
	host           string
	port           int
	clientId       string
	authToken      string
	channel        string
	timeoutSeconds int
}

func parseOptions(cfg config.Spec) (options, error) {
	o := options{}
	var err error
	o.host, o.port, err = cfg.Properties.MustParseAddress("address", defaultAddress)
	if err != nil {
		return options{}, fmt.Errorf("error parsing address value, %w", err)
	}
	o.channel, err = cfg.Properties.MustParseString("channel")
	if err != nil {
		return options{}, fmt.Errorf("error parsing channel value, %w", err)
	}
	o.authToken = cfg.Properties.ParseString("auth_token", "")
	o.clientId = cfg.Properties.ParseString("client_id", nuid.Next())
	o.timeoutSeconds, err = cfg.Properties.ParseIntWithRange("timeout_seconds", defaultTimeoutSeconds, 1, math.MaxInt32)
	if err != nil {
		return options{}, fmt.Errorf("error parsing default timeout seconds value, %w", err)
	}
	return o, nil
}
