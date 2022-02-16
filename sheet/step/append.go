package step

type Append struct {
	File    string `json:"file,omitempty" hcl:"file"`
	Content string `json:"content,omitempty" hcl:"content,optional"`
	Read    string `json:"read,omitempty" hcl:"read,optional"`
}
