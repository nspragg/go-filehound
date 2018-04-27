package filehound_test

import (
	"fmt"

	"github.com/nspragg/go-filehound/filehound"
)

func ExampleExt() {
	fh := filehound.New()
	fh.Query(filehound.Ext(".txt"))
	files := fh.Find()

	fmt.Println(files)
}

func ExampleSize() {
	fh := filehound.New()
	fh.Query(filehound.Size(1024))
	files := fh.Find()

	fmt.Println(files)
}
