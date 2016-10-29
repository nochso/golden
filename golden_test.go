package golden

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
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
	os.Remove(path)
}

func TestCase_Test(t *testing.T) {
	c := NewCase(t, "test-fixtures/in.txt")
	c.Test(func(c Case) []byte {
		return bytes.Repeat(c.In.Bytes(), 2)
	}, false)
}

func TestDirSlice(t *testing.T) {
	cases := DirSlice(t, "test-fixtures")
	expLen := 1
	if expLen != len(cases) {
		t.Errorf("expected %d case; got %d", expLen, len(cases))
	}
}

func bEqual(t *testing.T, exp, act []byte) {
	if !bytes.Equal(exp, act) {
		t.Fatal(diff(t, exp, act))
	}
}
