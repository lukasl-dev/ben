package steprunner

import (
	"fmt"
	"github.com/lukasl-dev/ben/sheet/step"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Clean runs a clear step.
func Clean(base step.Base, clear step.Clean) error {
	excluded := make(map[string]bool)
	for _, exclude := range clear.Clean.Exclude {
		exclude, _ = filepath.Abs(exclude)
		excluded[exclude] = true
	}

	if err := walkDirExcludeRoot(clear.Clean.Path, removeNotExcluded(excluded)); err != nil {
		return fmt.Errorf("step: %s: %w", base.Name, err)
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
