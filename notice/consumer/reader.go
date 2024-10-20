package consumer

import (
	"context"
	"errors"
	"github.com/meow-pad/persian/frame/plog"
	"github.com/meow-pad/persian/frame/plog/pfield"
	"github.com/meow-pad/persian/utils/coding"
	"github.com/segmentio/kafka-go"
	"sync/atomic"
	"time"
)

func NewReader(opts ...Option) (*Reader, error) {
	options := newOptions()
	for _, opt := range opts {
		opt(options)
	}
	if err := checkOptions(options); err != nil {
		return nil, err
	}
	kReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     options.KafkaBrokers,
		GroupID:     options.KafkaGroupId,
		GroupTopics: options.KafkaGroupTopics,
		Topic:       options.KafkaTopic,
		MaxBytes:    options.KafkaMaxBytes,
	})
	return &Reader{
		options: options,
		inner:   kReader,
		closeCh: make(chan struct{}),
	}, nil
}

type Reader struct {
	inner   *kafka.Reader
	options *Options
	closed  atomic.Bool
	closeCh chan struct{}
}

func (reader *Reader) Start() error {
	if reader.closed.Load() {
		return errors.New("reader is closed")
	}
	go reader.read()
	return nil
}

func (reader *Reader) Close() error {
	if !reader.closed.CompareAndSwap(false, true) {
		return nil
	}
	close(reader.closeCh)
	return reader.inner.Close()
}

func (reader *Reader) read() {
	defer coding.CatchPanicError("", func() {
		if reader.closed.Load() {
			return
		}
		go reader.read()
	})
	ctx := context.Background()
	for {
		if reader.closed.Load() {
			break
		}
		msg, err := reader.inner.FetchMessage(ctx)
		if err != nil {
			plog.Error("fetch message error", pfield.Error(err))
			select {
			case <-time.After(reader.options.RetryInterval):
				// 等待结束则继续
			case <-reader.closeCh:
				break
			}
			continue
		}
		plog.Debug("receive message:",
			pfield.String("Topic", msg.Topic),
			pfield.Int("Partition", msg.Partition),
			pfield.Int64("Offset", msg.Offset),
			pfield.String("Key", string(msg.Key)),
			pfield.String("Value", string(msg.Value)))
		// 处理消息
		if hErr := reader.options.Handler(msg.Value); hErr != nil {
			plog.Error("handle joytalk message error",
				pfield.String("topic", msg.Topic),
				pfield.ByteString("key", msg.Key),
				pfield.Error(hErr),
			)
		}
		if err = reader.inner.CommitMessages(ctx); err != nil {
			plog.Error("failed to commit messages", pfield.Error(err))
		}
	}
}
