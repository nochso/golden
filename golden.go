package golden

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pmezard/go-difflib/difflib"
)

var (
	// Extension that is added to the name of the input file to identify the
	// matching golden file.
	Extension = ".golden"
	// BasePath is put in front of paths passed to any of the Dir* functions.
	BasePath = "."
	// ChannelSize used by Dir() is arbitrary ┐(￣ヘ￣)┌
	ChannelSize = 32
)

// Dir returns a Case channel from a given directory.
//
// See ChannelSize for the channel size to be used.
// Any errors while walking the file system will fail and are not ignored.
func Dir(t *testing.T, path string) <-chan Case {
	path = filepath.Join(BasePath, path)
	ch := make(chan Case)
	walker := func(path string, info os.FileInfo, err error) error {
		must(t, err)
		if info.Mode().IsRegular() && !strings.HasSuffix(path, Extension) {
			ch <- NewCase(t, path)
		}
		return nil
	}
	go func() {
		must(t, filepath.Walk(path, walker))
		close(ch)
	}()
	return ch
}

// DirSlice returns a Case slice from a given directory.
//
// Any errors while walking the file system will fail and are not ignored.
func DirSlice(t *testing.T, path string) []Case {
	sl := []Case{}
	for c := range Dir(t, path) {
		sl = append(sl, c)
	}
	return sl
}

// File provides read/write access to test files.
type File struct {
	Case *Case // The case this file belongs to.
	Path string
}

func newFile(c *Case, path string) File {
	return File{c, path}
}

// Update the file by writing b to it.
func (f File) Update(b []byte) {
	if f.Case.T != nil {
		f.Case.T.Logf("updating golden file: %s", f.Path)
	}
	before := []byte{}
	if f.Exists() {
		before = f.Bytes()
	}
	f.Case.T.Log(diff(f.Case.T, before, b))
	must(f.Case.T, ioutil.WriteFile(f.Path, b, 0644))
}

// Reader returns a ReadCloser.
//
// This is basically os.File: remember to call Close(), especially if you have
// many files or read them multiple times.
func (f File) Reader() io.ReadCloser {
	fr, err := os.Open(f.Path)
	must(f.Case.T, err)
	return fr
}

// Bytes returns the content as a byte slice.
//
// It will fail when the file could not be read.
func (f File) Bytes() []byte {
	b, err := ioutil.ReadFile(f.Path)
	must(f.Case.T, err)
	return b
}

// String returns content as a string.
//
// It will fail when the file could not be read.
func (f File) String() string {
	return string(f.Bytes())
}

func (f File) Exists() bool {
	_, err := os.Stat(f.Path)
	return err == nil
}

// Case provides input and expected output for a single test case.
type Case struct {
	In  File
	Out File
	T   *testing.T
}

// NewCase returns a Case based on the given input file.
func NewCase(t *testing.T, path string) Case {
	c := Case{T: t}
	c.In = newFile(&c, path)
	c.Out = newFile(&c, path+Extension)
	return c
}

// Test this case with a function fn that takes a Case for input and returns
// []byte output.
//
// If update is false, the return value of fn is compared to the golden master.
// If update is true, the golden file will be overwritten with the output of
// the given function.
func (c Case) Test(fn func(Case) []byte, update bool) {
	c.T.Run(c.In.Path, func(t *testing.T) {
		// Let the current case work with the named sub test
		c.T = t
		actual := fn(c)
		if update { // update without testing
			c.Out.Update(actual)
			return
		}
		expected := c.Out.Bytes()
		if !bytes.Equal(actual, expected) {
			c.T.Error(diff(c.T, expected, actual))
		}
	})
}

func diff(t *testing.T, exp, act []byte) string {
	context := 1
	if testing.Verbose() {
		context = 3
	}
	a := difflib.SplitLines(string(exp))
	b := difflib.SplitLines(string(act))
	ud := difflib.UnifiedDiff{
		A:        a,
		B:        b,
		Context:  context,
		FromFile: "Expected",
		ToFile:   "Actual",
	}
	diff, err := difflib.GetUnifiedDiffString(ud)
	must(t, err)
	return fmt.Sprintf(
		"Bytes/Lines: %+d/%+d\n%s",
		len(act)-len(exp),
		len(b)-len(a),
		diff,
	)
}

// must fail a non-nil T or panic when err is set
func must(t *testing.T, err error) {
	if err == nil {
		return
	}
	if t == nil {
		panic(err)
	}
	t.Error(err)
}
