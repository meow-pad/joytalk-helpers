package msg

import (
	"errors"
	"github.com/meow-pad/persian/frame/plog"
	"github.com/meow-pad/persian/frame/plog/pfield"
	"github.com/meow-pad/persian/utils/json"
)

type Message interface {
	Notice() string
}

var builders = map[string]messageBuilder{}

type messageBuilder func() Message

func registerBuilder(name string, builder messageBuilder) {
	_, exist := builders[name]
	if exist {
		plog.Panic("message builder already exists", pfield.String("name", name))
	}
	builders[name] = builder
}

type noticeMsg struct {
	Notice string `json:"notice"`
}

func BuildMessage(data []byte, concernedNotices map[string]struct{}) (Message, error) {
	nMsg := noticeMsg{}
	if err := json.Unmarshal(data, &nMsg); err != nil {
		return nil, err
	}
	if concernedNotices != nil {
		if _, exist := concernedNotices[nMsg.Notice]; !exist {
			// 并不关注的消息，无需构建
			return nil, nil
		}
	}
	builder, _ := builders[nMsg.Notice]
	if builder == nil {
		return nil, errors.New("message builder not found")
	}
	msg := builder()
	if msg == nil {
		return nil, errors.New("message build failed")
	}
	return msg, nil
}
