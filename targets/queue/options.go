package queue

import (
	"fmt"
	"math"

	"github.com/kubemq-io/kubemq-sources/config"
	"github.com/nats-io/nuid"
)

const (
	defaultAddress = "localhost:50000"
)

type options struct {
	host              string
	port              int
	clientId          string
	authToken         string
	channel           string
	expirationSeconds int
	delaySeconds      int
	maxReceiveCount   int
	deadLetterQueue   string
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
	o.expirationSeconds, err = cfg.Properties.ParseIntWithRange("expiration_seconds", 0, 0, math.MaxInt32)
	if err != nil {
		return options{}, fmt.Errorf("error parsing expiration seconds, %w", err)
	}
	o.delaySeconds, err = cfg.Properties.ParseIntWithRange("delay_seconds", 0, 0, math.MaxInt32)
	if err != nil {
		return options{}, fmt.Errorf("error parsing delay seconds, %w", err)
	}
	o.maxReceiveCount, err = cfg.Properties.ParseIntWithRange("max_receive_count", 0, 0, math.MaxInt32)
	if err != nil {
		return options{}, fmt.Errorf("error max receive count seconds")
	}
	o.deadLetterQueue = cfg.Properties.ParseString("dead_letter_queue", "")
	return o, nil
}
