package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func captureStdout(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	f()

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}

func TestString(t *testing.T) {

	loopback := IPAddr{127, 0, 0, 1}

	out := captureStdout(func() {
		fmt.Printf("%v\n", loopback)
	})

	if out != "127.0.0.1\n" {
		t.Errorf("test1 : Unexpected string: %s", out)
	}
}
