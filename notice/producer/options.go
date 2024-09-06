package producer

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func newOptions() *Options {
	return &Options{
		KafkaMaxAttempts:       5,
		KafkaWriteTimeout:      2 * time.Second,
		KafkaRequiredAcks:      kafka.RequireNone,
		AsyncWrite:             false,
		AllowAutoTopicCreation: false,
	}
}

type Options struct {
	KafkaAddress           []string
	KafkaTopic             string
	KafkaMaxAttempts       int                // 最大写尝试次数
	KafkaWriteTimeout      time.Duration      // 超时后放弃写入后重试
	KafkaRequiredAcks      kafka.RequiredAcks // 写确认机制
	AsyncWrite             bool               // 异步写
	AllowAutoTopicCreation bool               // 允许topic缺失时自动创建
}

func checkOptions(opts *Options) error {
	if len(opts.KafkaAddress) <= 0 {
		return fmt.Errorf("missing kafka address")
	}
	return nil
}

type Option func(*Options)

func WithKafkaAddress(addr []string) Option {
	return func(o *Options) {
		o.KafkaAddress = addr
	}
}

func WithKafkaTopic(topic string) Option {
	return func(o *Options) {
		o.KafkaTopic = topic
	}
}

func WithKafkaMaxAttempts(maxAttempts int) Option {
	return func(o *Options) {
		o.KafkaMaxAttempts = maxAttempts
	}
}

func WithKafkaWriteTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.KafkaWriteTimeout = timeout
	}
}

func WithKafkaRequiredAcks(acks kafka.RequiredAcks) Option {
	return func(o *Options) {
		o.KafkaRequiredAcks = acks
	}
}

func WithAsyncWrite(async bool) Option {
	return func(o *Options) {
		o.AsyncWrite = async
	}
}

func WithAllowAutoTopicCreation(allowAutoTopicCreation bool) Option {
	return func(o *Options) {
		o.AllowAutoTopicCreation = allowAutoTopicCreation
	}
}
