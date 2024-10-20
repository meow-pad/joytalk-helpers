package client

type Options struct {
	RequestUri string
	PayReqUri  string
}

func (options *Options) complete() error {
	if len(options.RequestUri) > 0 && options.RequestUri[len(options.RequestUri)-1] != '/' {
		options.RequestUri = options.RequestUri + "/"
	}
	if len(options.PayReqUri) > 0 && options.PayReqUri[len(options.PayReqUri)-1] != '/' {
		options.PayReqUri = options.PayReqUri + "/"
	}
	if len(options.RequestUri) > 0 && len(options.PayReqUri) <= 0 {
		options.PayReqUri = options.RequestUri
	}
	return nil
}

type Option func(*Options)

func WithRequestUri(reqUri string) Option {
	return func(options *Options) {
		options.RequestUri = reqUri
	}
}

func WithPayuRequestUri(reqUri string) Option {
	return func(options *Options) {
		options.PayReqUri = reqUri
	}
}
