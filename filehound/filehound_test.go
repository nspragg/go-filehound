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
	actual := filehound.Create().
		Path(justFilesPath).
		Find()

	assertFiles(t, actual, justFiles)
}

func TestRecursiveSearch(t *testing.T) {
	actual := filehound.Create().
		Path(nestedFilesPath).
		Find()

	assertFiles(t, actual, nestedFiles)
}

func TestSearchByExtension(t *testing.T) {
	actual := filehound.Create().
		Path(justFilesPath).
		Ext(".txt").
		Find()

	assertFiles(t, actual, textFiles)
}

func TestSearchByExtensionExcludingPeriod(t *testing.T) {
	actual := filehound.Create().
		Path(justFilesPath).
		Ext("txt").
		Find()

	assertFiles(t, actual, textFiles)
}

func TestSearchByMultipleExtensions(t *testing.T) {
	actual := filehound.Create().
		Path(justFilesPath).
		Ext("txt", "json").
		Find()

	assertFiles(t, actual, justFiles)
}

func TestSearchByFileSize(t *testing.T) {
	actual := filehound.Create().
		Path(justFilesPath).
		Size(20).
		Find()

	expected := qualifyNames("./fixtures/justFiles/b.json")
	assertFiles(t, actual, expected)
}

func TestSeatchByFileGlob(t *testing.T) {
	actual := filehound.Create().
		Path(mixedPath).
		Glob("*.json").
		Find()

	assertFiles(t, actual, mixedFiles)
}

func TestDoesNotRecursiveWhenDepthIsZero(t *testing.T) {
	actual := filehound.Create().
		Path(deeplyNestedPath).
		Depth(0).
		Find()

	expected := qualifyNames("./fixtures/deeplyNested/c.json", "./fixtures/deeplyNested/d.json")
	assertFiles(t, actual, expected)
}

func TestSearchWhenDepthIsOne(t *testing.T) {
	actual := filehound.Create().
		Path(deeplyNestedPath).
		Depth(1).
		Find()

	expected := qualifyNames(
		"./fixtures/deeplyNested/c.json",
		"./fixtures/deeplyNested/d.json",
		"./fixtures/deeplyNested/mydir/e.json")

	assertFiles(t, actual, expected)
}

func TestSearchAtDepthN(t *testing.T) {
	actual := filehound.Create().
		Path(deeplyNestedPath).
		Depth(3).
		Find()

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
	actual := filehound.Create().
		Path(justFilesPath).
		IsEmpty().
		Find()

	expected := qualifyNames("./fixtures/justFiles/a.json", "./fixtures/justFiles/dummy.txt")

	assertFiles(t, actual, expected)
}

// // Matching by regular expressions
// func TestSearchByMatchingOnRegex(t *testing.T) {
// 	actual := filehound.Create().
// 		Path(justFilesPath).
// 		Match("a|b\\.\\w{3}").
// 		Find()

// 	expected := qualifyNames("./fixtures/justFiles/a.json", "./fixtures/justFiles/b.txt")

// 	assertFiles(t, actual, expected)
// }
