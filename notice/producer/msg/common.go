package msg

import "github.com/meow-pad/persian/utils/json"

const (
	topicPrefix = "monopoly."
)

type Message interface {
	Key() string
	Topic() string
	Notice() string
	Data() any
	ToNoticeBytes() ([]byte, error)
}

type Notice struct {
	Notice string `json:"notice"`
	Data   any    `json:"data,omitempty"`
}

func toNotice(msg Message) *Notice {
	return &Notice{
		Notice: msg.Notice(),
		Data:   msg.Data(),
	}
}

func toNoticeBytes(msg Message) ([]byte, error) {
	mBytes, err := json.Marshal(toNotice(msg))
	if err != nil {
		return nil, err
	}
	return mBytes, nil
}
