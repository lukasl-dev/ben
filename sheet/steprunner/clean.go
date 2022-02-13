package steprunner

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/lukasl-dev/ben/internal"
	"github.com/lukasl-dev/ben/sheet/step"
)

// Clean runs a clear step.
func Clean(st step.Step) error {
	excluded := make(map[string]bool)
	for _, exclude := range st.Clean.Exclude {
		exclude, _ = filepath.Abs(exclude)
		excluded[exclude] = true
	}

	err := walkDirExcludeRoot(st.Clean.Path, removeNotExcluded(excluded))
	if err != nil {
		return &internal.Error{
			Prefix: "step",
			Origin: st.Name,
			Err:    err,
			Suggestions: []string{
				"Check if the directory exists",
			},
		}
	}

	return nil
}

// walkDirExcludeRoot is like filepath.Walk, but excludes the root directory.
func walkDirExcludeRoot(root string, walkFn fs.WalkDirFunc) error {
	return filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if path == root {
			return err
		}
		return walkFn(path, entry, err)
	})
}

// removeNotExcluded removes all filed and directories that
// are not excluded.
func removeNotExcluded(excluded map[string]bool) fs.WalkDirFunc {
	return func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		path, _ = filepath.Abs(path)
		if excluded[path] {
			if entry.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if isSubpathOfExcluded(path, excluded) {
			return nil
		}
		return os.RemoveAll(path)
	}
}

// isSubpathOfExcluded reports whether the given path is a
// subpath of an entry in the given excluded paths.
func isSubpathOfExcluded(path string, excluded map[string]bool) bool {
	for excludedPath := range excluded {
		if strings.HasPrefix(excludedPath, path) {
			return true
		}
	}
	return false
}
