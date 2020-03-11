package main

import (
    "testing"
    "bytes"
    "io"
    "os"
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

func TestDo(t *testing.T) {

    out := captureStdout(func() {
            do(21)
    })
    if out != "Twice 21 is 42\n" {
        t.Errorf("test1 : Unexpected string: %s", out)
    }

    out = captureStdout(func() {
            do("hello")
    })
    if out != "\"hello\" is 5 bytes long\n" {
        t.Errorf("test2 : Unexpected string: %s", out)
    }

    out = captureStdout(func() {
            do(true)
    })
    if out != "I don't know about type bool!\n" {
        t.Errorf("test3 : Unexpected string: %s", out)
    }
}
