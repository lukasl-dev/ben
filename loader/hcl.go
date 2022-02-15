package loader

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/lukasl-dev/ben/sheet"
	"io"
)

// loadHCL loads a sheet from r via the hashicorp library.
func loadHCL(uri string, r io.Reader, _ Options) (*sheet.Sheet, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var s sheet.Sheet
	return &s, hclsimple.Decode(uri, data, nil, &s)
}
