package types

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/kubemq-io/kubemq-go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Request struct {
	Metadata string `json:"metadata,omitempty"`
	Data     []byte `json:"data,omitempty"`
	Channel  string `json:"channel"`
}

func NewRequest() *Request {
	return &Request{
		Metadata: "",
		Data:     nil,
	}
}

func (r *Request) SetMetadata(value string) *Request {
	r.Metadata = value
	return r
}

func (r *Request) SetChannel(value string) *Request {
	r.Channel = value
	return r
}

func (r *Request) SetData(value []byte) *Request {
	r.Data = value
	return r
}

func (r *Request) Size() float64 {
	return float64(len(r.Data))
}

func ParseRequest(body []byte) (*Request, error) {
	if body == nil {
		return nil, fmt.Errorf("empty request")
	}
	req := &Request{}
	err := json.Unmarshal(body, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (r *Request) MarshalBinary() []byte {
	data, _ := json.Marshal(r)
	return data
}

func (r *Request) ToEvent() *kubemq.Event {
	return kubemq.NewEvent().
		SetBody(r.MarshalBinary())
}

func (r *Request) ToEventStore() *kubemq.EventStore {
	return kubemq.NewEventStore().
		SetBody(r.MarshalBinary())
}

func (r *Request) ToCommand() *kubemq.Command {
	return kubemq.NewCommand().
		SetBody(r.MarshalBinary())
}

func (r *Request) ToQuery() *kubemq.Query {
	return kubemq.NewQuery().
		SetBody(r.MarshalBinary())
}

func (r *Request) ToQueueMessage() *kubemq.QueueMessage {
	return kubemq.NewQueueMessage().
		SetBody(r.MarshalBinary())
}

func (r *Request) String() string {
	str, err := json.MarshalToString(r)
	if err != nil {
		return ""
	}
	return str
}

func (r *Request) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
