package filehound_test

import (
	"fmt"

	"github.com/nspragg/go-filehound/filehound"
)

func ExampleFind() {
	fh := filehound.New()
	files := fh.Find()

	fmt.Println(files)
}

func ExampleSize() {
	fh := filehound.New()
	files := fh.Find()

	fmt.Println(files)
}
