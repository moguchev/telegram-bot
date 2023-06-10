package bot

type (
	options struct {
		debug bool
	}

	Option func(o *options)
)

// WithDebug enables debug output
func WithDebug() Option {
	return func(o *options) {
		o.debug = true
	}
}
