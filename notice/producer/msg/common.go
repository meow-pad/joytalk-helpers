package msg

import "github.com/meow-pad/persian/utils/json"

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

func ToNotice(msg Message) *Notice {
	return &Notice{
		Notice: msg.Notice(),
		Data:   msg.Data(),
	}
}

func ToNoticeBytes(msg Message) ([]byte, error) {
	mBytes, err := json.Marshal(ToNotice(msg))
	if err != nil {
		return nil, err
	}
	return mBytes, nil
}
