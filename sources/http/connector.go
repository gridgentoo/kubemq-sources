package http

import "github.com/kubemq-hub/builder/common"

func Connector() *common.Connector {
	return common.NewConnector().
		SetKind("source.http").
		SetDescription("HTTP/REST source properties").
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("methods").
				SetDescription("list of supported methods separated by a comma").
				SetMust(true),
		).
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("path").
				SetDescription("http endpoint path").
				SetMust(true).
				SetDefault("/"),
		)
}
