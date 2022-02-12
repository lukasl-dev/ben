package spinner

type Options struct {
	Indent  uint   `json:"indent,omitempty"`
	Pending string `json:"pending,omitempty"`
	Success string `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`
}
