package filehound_test

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/nspragg/go-filehound/filehound"
)

var textFiles = qualifyNames("./fixtures/justFiles/dummy.txt")

var justFiles = qualifyNames(
	"./fixtures/justFiles/a.json",
	"./fixtures/justFiles/b.json",
	"./fixtures/justFiles/dummy.txt",
)
var justFilesPath = qualifyName("./fixtures/justFiles")

var nestedFiles = qualifyNames(
	"./fixtures/nested/c.json",
	"./fixtures/nested/d.json",
	"./fixtures/nested/mydir/e.json",
)
var nestedFilesPath = qualifyName("./fixtures/nested")
var deeplyNestedPath = qualifyName("./fixtures/deeplyNested")

var mixedFiles = qualifyNames(
	"./fixtures/mixed/a.json",
	"./fixtures/mixed/aabbcc.json",
	"./fixtures/mixed/ab.json",
	"./fixtures/mixed/z.json",
)
var mixedPath = qualifyName("./fixtures/mixed")

var Ext = filehound.Ext
var Path = filehound.Path
var Size = filehound.Size
var Glob = filehound.Glob
var Depth = filehound.Depth
var IsEmpty = filehound.IsEmpty
var Match = filehound.Match

func qualifyName(name string) string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, name)
}

func qualifyNames(names ...string) []string {
	qualifiedNames := make([]string, 0)
	for _, name := range names {
		qualifiedNames = append(qualifiedNames, qualifyName(name))
	}

	return qualifiedNames
}

func assertFiles(t *testing.T, actual, expected []string) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("unexpected files.  Expected %v. Got : %v.", expected, actual)
	}
}

func TestAllFilesInDirectory(t *testing.T) {
	hound := filehound.New()
	hound.Query(Path(justFilesPath))
	actual := hound.Find()

	assertFiles(t, actual, justFiles)
}

func TestRecursiveSearch(t *testing.T) {
	hound := filehound.New()
	hound.Query(Path(nestedFilesPath))
	actual := hound.Find()

	assertFiles(t, actual, nestedFiles)
}

func TestSearchByExtension(t *testing.T) {
	hound := filehound.New()
	hound.Query(Path(justFilesPath))
	hound.Query(Ext(".txt"))
	actual := hound.Find()

	assertFiles(t, actual, textFiles)
}

func TestSearchByExtensionExcludingPeriod(t *testing.T) {
	hound := filehound.New()
	hound.Query(Path(justFilesPath))
	hound.Query(Ext("txt"))
	actual := hound.Find()

	assertFiles(t, actual, textFiles)
}

func TestQueryVarArgs(t *testing.T) {
	hound := filehound.New()
	hound.Query(
		Path(justFilesPath),
		Ext("txt"))

	actual := hound.Find()

	assertFiles(t, actual, textFiles)
}

func TestSearchByMultipleExtensions(t *testing.T) {
	hound := filehound.New()
	hound.Query(Path(justFilesPath))
	hound.Query(Ext("txt", "json"))
	actual := hound.Find()

	assertFiles(t, actual, justFiles)
}

func TestSearchByFileSize(t *testing.T) {
	hound := filehound.New()
	hound.Query(Path(justFilesPath))
	hound.Query(Size(20))
	actual := hound.Find()

	expected := qualifyNames("./fixtures/justFiles/b.json")
	assertFiles(t, actual, expected)
}

func TestSeatchByFileGlob(t *testing.T) {
	hound := filehound.New()
	hound.Query(Path(mixedPath))
	hound.Query(Glob("*.json"))
	actual := hound.Find()

	assertFiles(t, actual, mixedFiles)
}

func TestDoesNotRecursiveWhenDepthIsZero(t *testing.T) {
	hound := filehound.New()
	hound.Query(Path(deeplyNestedPath))
	hound.Query(Depth(0))
	actual := hound.Find()

	expected := qualifyNames("./fixtures/deeplyNested/c.json", "./fixtures/deeplyNested/d.json")
	assertFiles(t, actual, expected)
}

func TestSearchWhenDepthIsOne(t *testing.T) {
	hound := filehound.New()
	hound.Query(Path(deeplyNestedPath))
	hound.Query(Depth(1))
	actual := hound.Find()

	expected := qualifyNames(
		"./fixtures/deeplyNested/c.json",
		"./fixtures/deeplyNested/d.json",
		"./fixtures/deeplyNested/mydir/e.json")

	assertFiles(t, actual, expected)
}

func TestSearchAtDepthN(t *testing.T) {
	hound := filehound.New()
	hound.Query(Path(deeplyNestedPath))
	hound.Query(Depth(3))
	actual := hound.Find()

	expected := qualifyNames(
		"./fixtures/deeplyNested/c.json",
		"./fixtures/deeplyNested/d.json",
		"./fixtures/deeplyNested/mydir/e.json",
		"./fixtures/deeplyNested/mydir/mydir2/f.json",
		"./fixtures/deeplyNested/mydir/mydir2/mydir3/z.json",
		"./fixtures/deeplyNested/mydir/mydir2/y.json")

	sort.Strings(actual)

	assertFiles(t, actual, expected)
}

func TestSearchByEmptyFile(t *testing.T) {
	hound := filehound.New()
	hound.Query(Path(justFilesPath))
	hound.Query(IsEmpty())
	actual := hound.Find()

	expected := qualifyNames("./fixtures/justFiles/a.json", "./fixtures/justFiles/dummy.txt")

	assertFiles(t, actual, expected)
}

// func TestSearchByMatchingOnRegex(t *testing.T) {
// 	hound := filehound.New()
// 	hound.Query(Path(justFilesPath))
// 	hound.Query(Match("(a|b).json"))
// 	actual := hound.Find()
// 	expected := qualifyNames("./fixtures/justFiles/a.json", "./fixtures/justFiles/b.json")

// 	assertFiles(t, actual, expected)
// }
