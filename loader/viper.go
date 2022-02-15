package loader

import (
	"github.com/lukasl-dev/ben/sheet"
	"github.com/spf13/viper"
	"io"
)

// loadViper loads a sheet from r via the viper library.
func loadViper(uri string, r io.Reader, opts Options) (*sheet.Sheet, error) {
	v := viper.New()
	v.SetConfigType(extensionOf(uri, opts))
	if err := v.ReadConfig(r); err != nil {
		return nil, err
	}

	var s sheet.Sheet
	return &s, v.Unmarshal(&s)
}
