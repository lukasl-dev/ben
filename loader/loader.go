package loader

import (
	"github.com/lukasl-dev/ben/sheet"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// loader represents a loader function.
type loader = func(uri string, r io.Reader, opts Options) (*sheet.Sheet, error)

// loaders is a map from a file extension ('.' included) to a loader function.
var loaders = map[string]loader{
	"hcl": loadHCL,
}

type Options struct {
	// Format is the file format to use. If it is set, the given Format will be
	// used instead of the format specified by the file's extension.
	Format string `json:"format,omitempty"`
}

// NoOptions are the default Options.
var NoOptions = Options{}

// Load loads a sheet from uri. The way how the sheet is loaded is determined by
// the file extension. If uri should have a http(s) scheme, the sheet is fetched
// from the remote location via a GET request. Otherwise, the file is loaded
// from the local file system.
func Load(uri string, opts Options) (*sheet.Sheet, error) {
	r, err := reader(uri)
	if err != nil {
		return nil, err
	}
	return load(uri, r, opts)
}

// load loads a sheet from r using the appropriate loader from loaders.
func load(uri string, r io.Reader, opts Options) (*sheet.Sheet, error) {
	ext := extensionOf(uri, opts)
	l, ok := loaders[ext]
	if !ok {
		l = loadViper
	}
	return l(uri, r, opts)
}

// reader opens a reader for uri.
func reader(uri string) (io.Reader, error) {
	u, err := url.ParseRequestURI(uri)
	if err == nil && (u.Scheme == "http" || u.Scheme == "https") {
		return fetch(uri)
	}
	return open(uri)
}

// fetch dispatches a RESt request to uri and returns a reader that reads fro
func fetch(uri string) (io.Reader, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

// open opens the file located at path and returns a reader for it.
func open(path string) (io.Reader, error) {
	return os.Open(path)
}

func extensionOf(uri string, opts Options) string {
	if opts.Format != "" {
		return opts.Format
	}
	return strings.TrimLeft(filepath.Ext(uri), ".")
}
