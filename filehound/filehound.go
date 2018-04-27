package filehound

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type filterFn func(path string, info os.FileInfo) bool

func depth(path string) int {
	parts := strings.Split(path, string(filepath.Separator))
	return len(parts) - 1
}

// Option ...
type Option func(*Filehound)

// Filehound ...
type Filehound struct {
	root     string
	filters  []filterFn
	maxDepth int
}

// Query ....
func (f *Filehound) Query(opts ...Option) {
	for _, opt := range opts {
		opt(f)
	}
}

// New returns an instance of Filehound
func New() *Filehound {
	cwd, _ := os.Getwd()
	return &Filehound{root: cwd, maxDepth: 100}
}

// Depth sets the max recursion depth
func Depth(depth int) Option {
	return func(f *Filehound) {
		f.maxDepth = depth
	}
}

// Size filters files by size
func Size(size int64) Option {
	return func(f *Filehound) {
		f.Filter(func(path string, info os.FileInfo) bool {
			return info.Size() == size
		})
	}
}

// IsEmpty ...
func IsEmpty() Option {
	return Size(0)
}

// Path sets the root of the search path. Defaults to the cwd
func Path(root string) Option {
	return func(f *Filehound) {
		f.root = root
	}
}

// Ext filters by file extension
func Ext(exts ...string) Option {
	return func(f *Filehound) {
		f.Filter(func(path string, info os.FileInfo) bool {
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
}

// Match filters by file regexp
func Match(pattern string) Option {
	return func(f *Filehound) {
		f.Filter(func(path string, info os.FileInfo) bool {
			return regexp.MustCompile(pattern).MatchString(path)
		})
	}
}

// Glob ...
func Glob(pattern string) Option {
	return func(f *Filehound) {
		f.Filter(func(path string, info os.FileInfo) bool {
			isMatch, _ := filepath.Match(pattern, filepath.Base(path))
			return isMatch
		})
	}
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

func (f *Filehound) atMaxDepth(path string) bool {
	depth := depth(filepath.Dir(path)) - depth(f.root)
	return depth > f.maxDepth
}

// Find executes the search
func (f *Filehound) Find() []string {
	var files []string
	filepath.Walk(f.root, func(path string, info os.FileInfo, err error) error {
		if f.atMaxDepth(path) {
			return nil
		}
		if !info.IsDir() && f.isMatch(path, info) {
			files = append(files, path)
		}
		return nil
	})

	return files
}

// Filter ...
func (f *Filehound) Filter(fn filterFn) *Filehound {
	f.filters = append(f.filters, fn)
	return f
}
