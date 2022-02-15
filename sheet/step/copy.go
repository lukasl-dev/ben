package step

type Copy struct {
	// From the path to the file or directory to copy. It can be a URL for remote
	// files.
	From string `json:"from,omitempty" hcl:"from"`

	// To the path to the destination.
	To string `json:"to,omitempty" hcl:"to"`
}
