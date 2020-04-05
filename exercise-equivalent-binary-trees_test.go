package main

import (
    "testing"
    "bytes"
    "io"
    "os"
    "golang.org/x/tour/tree"
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

func TestSame(t *testing.T) {

    ret := Same(tree.New(1), tree.New(1))
    if ret == false {
      t.Errorf("test1 : Unexpected return value: %v", ret)
    }

    ret = Same(tree.New(1), tree.New(2))
    if ret == true {
      t.Errorf("test2 : Unexpected return value: %v", ret)
    }
}
