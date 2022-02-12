package spinner

type Options struct {
	Prefix *string `json:"prefix,omitempty"`
	Suffix *string `json:"text,omitempty"`
	Color  string  `json:"color,omitempty"`
}

func (opts *Options) normalize() {
	if opts.Color == "" {
		opts.Color = "#3b82f6"
	}
}
