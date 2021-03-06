package golden

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/k0kubun/pp"
)

var (
	update        = flag.Bool("update", false, "update golden files")
	expectedCases = 3
)

func TestFile(t *testing.T) {
	c := NewCase(t, "test-fixtures/in.txt")
	exp := []byte("abc")

	bEqual(t, exp, []byte(c.In.String()))
	bEqual(t, exp, c.In.Bytes())

	r := c.In.Reader()
	act, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	bEqual(t, exp, act)
}

func TestFile_Update(t *testing.T) {
	path := "test-fixtures/update.txt"
	c := NewCase(t, path)
	if c.In.Exists() {
		t.Errorf("Expecting file not to exist: %s", path)
	}
	b := []byte("foo")
	c.In.Update(b)
	bEqual(t, b, c.In.Bytes())

	c.In.Update([]byte{})
	bEqual(t, []byte{}, c.In.Bytes())
	os.Remove(path)
}

func TestCase_Diff(t *testing.T) {
	c := NewCase(t, "test-fixtures/in.txt")
	act := bytes.Repeat(c.In.Bytes(), 2)
	if *update {
		c.Out.Update(act)
	}
	c.Diff(string(act))
}

func TestDirSlice(t *testing.T) {
	cases := DirSlice(t, "test-fixtures")
	if expectedCases != len(cases) {
		t.Errorf("expected %d cases; got %d", expectedCases, len(cases))
	}
}

func TestFile_Split(t *testing.T) {
	testSplit(t, NewCase(t, "test-fixtures/split.txt"))
	testSplit(t, NewCase(t, "test-fixtures/split-crlf.txt"))
}

func testSplit(t *testing.T, c Case) {
	s := c.In.Split("===")
	pp.ColoringEnabled = false
	if *update {
		c.Out.Update([]byte(pp.Sprint(s)))
	}
	c.Diff(pp.Sprint(s))
}

func TestTestDir(t *testing.T) {
	count := 0
	TestDir(t, "test-fixtures", func(tc Case) {
		count++
	})
	if count != expectedCases {
		t.Errorf("expected %d cases; got %d", expectedCases, count)
	}
}

func bEqual(t *testing.T, exp, act []byte) {
	if !bytes.Equal(exp, act) {
		t.Fatal(diff(t, exp, act))
	}
}
