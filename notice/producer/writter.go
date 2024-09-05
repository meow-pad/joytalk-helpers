package producer

import (
	"context"
	"github.com/meow-pad/joytalk-helpers/notice/producer/msg"
	"github.com/meow-pad/persian/frame/plog"
	"github.com/segmentio/kafka-go"
	"time"
)

func NewWriter(opts ...Option) (*Writer, error) {
	options := newOptions()
	for _, opt := range opts {
		opt(options)
	}
	writer := new(Writer)
	if err := writer.init(options); err != nil {
		return nil, err
	}
	return writer, nil
}

type Writer struct {
	inner *kafka.Writer
}

func (writer *Writer) init(options *Options) error {
	err := checkOptions(options)
	if err != nil {
		return err
	}
	writer.inner = &kafka.Writer{
		Addr:         kafka.TCP(options.KafkaAddress...),
		Topic:        options.KafkaTopic,
		Balancer:     &kafka.Hash{},
		MaxAttempts:  options.KafkaMaxAttempts,
		WriteTimeout: options.KafkaWriteTimeout,
		RequiredAcks: options.KafkaRequiredAcks,
		Async:        true,
		ErrorLogger:  plog.SugarLogger(),
	}
	return nil
}

func (writer *Writer) WriteMessage(ctx context.Context, notice msg.Message) error {
	value, err := notice.ToNoticeBytes()
	if err != nil {
		return err
	}
	kMsg := kafka.Message{
		Topic:         notice.Topic(),
		Partition:     0,
		Offset:        0,
		HighWaterMark: 0,
		Key:           []byte(notice.Key()),
		Value:         value,
		Time:          time.Now(),
	}
	err = writer.inner.WriteMessages(ctx, kMsg)
	return err
}
