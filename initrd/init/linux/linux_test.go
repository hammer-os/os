package linux

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestTouch(t *testing.T) {
	if err := Touch("/tmp/a/b/c/file", 0644); err != nil {
		t.Fatalf("touch /tmp/a/b/c/file")
	}
	defer os.RemoveAll("/tmp/a")

	fi, err := os.Stat("/tmp/a/b/c/file")
	if err != nil {
		t.Fatalf("stat /tmp/a/b/c/file")
	}
	if fi.Mode() != 0644 {
		t.Fatalf("expected mode 0664, have %d", fi.Mode())
	}
	if fi.Size() != 0 {
		t.Fatalf("expected size 0, have %d", fi.Size())
	}
	if fi.IsDir() != false {
		t.Fatalf("expected file, have directory")
	}

	if err := Touch("/tmp/a/b/c/file", 0644); err != nil {
		t.Fatalf("touch /tmp/a/b/c/file")
	}
}

func TestWrite(t *testing.T) {
	b := []byte("hello world")
	if err := Write("/tmp/a/b/c/file", b, 0644); err != nil {
		t.Fatalf("touch /tmp/a/b/c/file")
	}
	defer os.RemoveAll("/tmp/a")

	fi, err := os.Stat("/tmp/a/b/c/file")
	if err != nil {
		t.Fatalf("stat /tmp/a/b/c/file")
	}
	if fi.Mode() != 0644 {
		t.Fatalf("expected mode 0664, have %d", fi.Mode())
	}
	if fi.Size() != int64(len(b)) {
		t.Fatalf("expected size %d, have %d", len(b), fi.Size())
	}
	if fi.IsDir() != false {
		t.Fatalf("expected file, have directory")
	}

	data, err := ioutil.ReadFile("/tmp/a/b/c/file")
	if err != nil {
		t.Fatalf("reading /tmp/a/b/c/file: %v", err)
	}
	if !bytes.Equal(b, data) {
		t.Fatalf("exepcted content %q, have %q", b, data)
	}
}
