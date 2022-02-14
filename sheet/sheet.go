package sheet

import (
	"bytes"
	"github.com/lukasl-dev/ben/sheet/job"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
	"os"
)

// Sheet is a collection of jobs. It represents the content of a configuration
// file.
type Sheet struct {
	// Name is the display name of the sheet.
	Name string `json:"name,omitempty"`

	// Description is a short description about the usage of the sheet.
	Description string `json:"description,omitempty"`

	// Jobs is a slice of jobs that can be executed on the sheet. The key is the
	// unique name of a job.
	Jobs []job.Job `json:"jobs,omitempty"`

	// WorkDir is the global working directory of a sheet.
	WorkDir string `json:"workDir,omitempty"`
}

// Validate validates s.
func (s Sheet) Validate() error {
	for _, j := range s.Jobs {
		for _, step := range j.Steps {
			if err := step.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// Load loads a sheet from a file or from a URL.
func Load(pathOrURL string) (*Sheet, error) {
	r, err := reader(pathOrURL)
	if err != nil {
		return nil, err
	}

	v := viper.New()
	v.SetConfigFile(pathOrURL)
	if err := v.ReadConfig(r); err != nil {
		return nil, err
	}

	var s *Sheet
	return s, v.Unmarshal(&s, func(config *mapstructure.DecoderConfig) {
		config.TagName = "json"
	})
}

// reader returns a reader for the given path or URL.
func reader(pathOrURL string) (io.Reader, error) {
	u, err := url.ParseRequestURI(pathOrURL)
	if err == nil && (u.Scheme == "http" || u.Scheme == "https") {
		return fetch(pathOrURL)
	}
	return read(pathOrURL)
}

// fetch unmarshals a sheet from the response body of the HTTP request that is
// sent to the given url.
func fetch(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

// read unmarshals a sheet from a file that is located at the given local path.
func read(path string) (io.Reader, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(content), nil
}
