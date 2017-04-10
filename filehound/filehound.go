package filehound

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

type filterFn func(path string, info os.FileInfo) bool

// Filehound ...
type Filehound struct {
	path     string
	filters  []filterFn
	maxDepth int
}

// Create returns an instance of Filehound
func Create() *Filehound {
	cwd, _ := os.Getwd()
	return &Filehound{path: cwd, maxDepth: 100}
}

// Path sets the search path. Defaults to the cwd
func (f *Filehound) Path(path string) *Filehound {
	f.path = path
	return f
}

// Ext filters by file extension
func (f *Filehound) Ext(exts ...string) *Filehound {
	return f.Filter(func(path string, info os.FileInfo) bool {
		for _, ext := range exts {
			if strings.HasPrefix(ext, ".") {
				ext = ext[1:]
			}
			if filepath.Ext(path)[1:] == ext {
				return true
			}
		}
		return false
	})
}

// Depth sets the max recursion depth
func (f *Filehound) Depth(depth int) *Filehound {
	f.maxDepth = depth
	return f
}

// Size filters files by size
func (f *Filehound) Size(size int64) *Filehound {
	return f.Filter(func(path string, info os.FileInfo) bool {
		return info.Size() == size
	})
}

func (f *Filehound) isMatch(path string, info os.FileInfo) bool {
	if len(f.filters) == 0 {
		return true
	}

	for _, fn := range f.filters {
		if fn(path, info) {
			return true
		}
	}

	return false
}

// Find executes the search
func (f *Filehound) Find() []string {
	depth := 0
	files := make([]string, 0)
	filepath.Walk(f.path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if depth > f.maxDepth {
				return io.EOF
			}
			depth++
		}
		if !info.IsDir() && f.isMatch(path, info) {
			files = append(files, path)
		}
		return nil
	})

	return files
}

// Glob sets ...
func (f *Filehound) Glob(pattern string) *Filehound {
	return f.Filter(func(path string, info os.FileInfo) bool {
		isMatch, _ := filepath.Match(pattern, filepath.Base(path))
		return isMatch
	})
}

// Filter ...
func (f *Filehound) Filter(fn filterFn) *Filehound {
	f.filters = append(f.filters, fn)
	return f
}
