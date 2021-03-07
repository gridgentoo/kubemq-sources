package msk

import (
	"github.com/kubemq-hub/builder/connector/common"
)

func Connector() *common.Connector {
	return common.NewConnector().
		SetKind("aws.msk").
		SetDescription("AWS MSK source properties").
		SetName("MSK").
		SetProvider("AWS").
		SetCategory("Messaging").
		SetTags("streaming","kafka","cloud","managed").
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("brokers").
				SetDescription("Set Brokers list").
				SetMust(true).
				SetDefault(""),
		).
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("topics").
				SetDescription("Set Topics list").
				SetMust(true).
				SetDefault(""),
		).
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("consumer_group").
				SetDescription("Set Consumer Group").
				SetMust(true).
				SetDefault(""),
		).
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("sasl_username").
				SetDescription("Set Saal Username").
				SetMust(false).
				SetDefault(""),
		).
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("sasl_password").
				SetDescription("Set Saal Password").
				SetMust(false).
				SetDefault(""),
		).
		AddProperty(
			common.NewProperty().
				SetKind("bool").
				SetName("dynamic_mapping").
				SetDescription("Set Topic/Channel dynamic mapping").
				SetMust(true).
				SetDefault("true"),
		)
}
