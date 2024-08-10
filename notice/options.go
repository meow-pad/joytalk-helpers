package notice

import "errors"

func newOptions() *Options {
	return &Options{
		KafkaMaxBytes: 10e6,
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
	if len(options.KafkaTopic) <= 0 {
		return errors.New("kafkaTopic is empty")
	}
	return nil
}

type Options struct {
	KafkaBrokers     []string
	KafkaGroupId     string
	KafkaTopic       string
	KafkaMaxBytes    int
	Handler          func(msgData []byte) error
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

func WithConcernedNotices(notices []string) Option {
	return func(options *Options) {
		options.ConcernedNotices = make(map[string]struct{})
		for _, notice := range notices {
			options.ConcernedNotices[notice] = struct{}{}
		}
	}
}
