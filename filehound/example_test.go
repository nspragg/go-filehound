package filehound_test

import (
	"fmt"
	"path/filepath"

	"github.com/nspragg/go-filehound/filehound"
)

func ExamplePath() {
	client := filehound.New()
	client.Query(filehound.Path("./fixtures/examples"))
	files := client.Find()

	fmt.Printf("%d files found\n", len(files))

	// Output: 2 files found
}

func ExampleExt() {
	client := filehound.New()
	client.Query(filehound.Path("./fixtures/examples"))
	client.Query(filehound.Ext(".json"))
	files := client.Find()

	fmt.Println(filepath.Base(files[0]))

	// Output: example.json
}

func ExampleGlob() {
	client := filehound.New()
	client.Query(filehound.Path("./fixtures/examples"))
	client.Query(filehound.Glob("*.txt"))
	files := client.Find()

	fmt.Printf("%d file found\n", len(files))

	// Output: 1 file found
}

func ExampleSize() {
	bytes := int64(33)

	client := filehound.New()
	client.Query(filehound.Path("./fixtures/examples"))
	client.Query(filehound.Size(bytes))
	files := client.Find()

	fmt.Printf("%d file found that is %d bytes in length\n", len(files), bytes)

	// Output: 1 file found that is 33 bytes in length
}
