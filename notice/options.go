package notice

import (
	"errors"
	"github.com/meow-pad/joytalk-helpers/notice/msg"
	"github.com/meow-pad/persian/frame/plog"
	"github.com/meow-pad/persian/frame/plog/pfield"
	"time"
)

func newOptions() *Options {
	return &Options{
		KafkaMaxBytes: 10e6,
		RetryInterval: 20 * time.Second,
	}
}

func checkOptions(options *Options) error {
	if options.Handler == nil {
		return errors.New("handler is nil")
	}
	if options.KafkaMaxBytes <= 0 {
		return errors.New("kafkaMaxBytes must be greater than zero")
	}
	if len(options.KafkaBrokers) <= 0 {
		return errors.New("kafkaBrokers is empty")
	}
	if len(options.KafkaGroupId) > 0 {
		if len(options.KafkaTopic) <= 0 && len(options.KafkaGroupTopics) <= 0 {
			return errors.New("KafkaTopic and kafkaGroupTopics is empty")
		}
	} else {
		if len(options.KafkaTopic) <= 0 {
			return errors.New("kafkaTopic is empty")
		}
	}
	if options.RetryInterval <= 0 {
		return errors.New("retryInterval is zero")
	}
	return nil
}

type Options struct {
	KafkaBrokers     []string
	KafkaGroupId     string
	KafkaTopic       string
	KafkaGroupTopics []string
	KafkaMaxBytes    int
	Handler          func(msgData []byte) error
	RetryInterval    time.Duration // 连接出错时的重试间隔
	ConcernedNotices map[string]struct{}
}

type Option func(options *Options)

func WithKafkaBrokers(brokers []string) Option {
	return func(options *Options) {
		options.KafkaBrokers = brokers
	}
}

func WithKafkaGroupId(kafkaGroupId string) Option {
	return func(options *Options) {
		options.KafkaGroupId = kafkaGroupId
	}
}

func WithKafkaGroupTopics(topics ...string) Option {
	return func(options *Options) {
		options.KafkaGroupTopics = topics
	}
}

func WithKafkaTopic(kafkaTopic string) Option {
	return func(options *Options) {
		options.KafkaTopic = kafkaTopic
	}
}

func WithKafkaMaxBytes(kafkaMaxBytes int) Option {
	return func(options *Options) {
		options.KafkaMaxBytes = kafkaMaxBytes
	}
}

func WithHandler(handler func(msgData []byte) error) Option {
	return func(options *Options) {
		options.Handler = handler
	}
}

func WithConcernedNotices(notices ...string) Option {
	return func(options *Options) {
		options.ConcernedNotices = make(map[string]struct{})
		options.KafkaGroupTopics = nil
		for _, notice := range notices {
			topic, exist := msg.GetNoticeTypeTopic(notice)
			if !exist {
				plog.Error("notice type topic not found", pfield.String("notice", notice))
				continue
			}
			options.ConcernedNotices[notice] = struct{}{}
			options.KafkaGroupTopics = append(options.KafkaGroupTopics, topic)
		}
	}
}

func WithRetryInterval(retryInterval time.Duration) Option {
	return func(options *Options) {
		options.RetryInterval = retryInterval
	}
}
