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

func TestDescribe(t *testing.T) {
    var i interface{}

    out := captureStdout(func() {
        describe(i)
    })
    if out != "(<nil>, <nil>)\n" {
        t.Errorf("Unexpected string: %s i: %v", out, i)
    }

    i = 42
    out = captureStdout(func() {
        describe(i)
    })
    if out != "(42, int)\n" {
        t.Errorf("Unexpected string: %s i: %v", out, i)
    }

    i = "hello"
    out = captureStdout(func() {
        describe(i)
    })
    if out != "(hello, string)\n" {
        t.Errorf("Unexpected string: %s i: %v", out, i)
    }
}
