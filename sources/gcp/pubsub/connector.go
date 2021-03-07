package pubsub

import (
	"github.com/kubemq-hub/builder/connector/common"
)

func Connector() *common.Connector {
	return common.NewConnector().
		SetKind("gcp.pubsub").
		SetDescription("AWS Pubsub source properties").
		SetName("PubSub").
		SetProvider("GCP").
		SetCategory("Messaging").
		SetTags("streaming","cloud","managed").
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("project_id").
				SetDescription("Set Project Id").
				SetMust(true),
		).
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("subscriber_id").
				SetDescription("Set Subscriber Id").
				SetMust(true).
				SetDefault(""),
		).
		AddProperty(
			common.NewProperty().
				SetKind("multilines").
				SetName("credentials").
				SetDescription("Set gcp Credentials").
				SetMust(true).
				SetDefault(""),
		)
}
