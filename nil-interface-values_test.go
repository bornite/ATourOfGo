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
    var i I

    out := captureStdout(func() {
        describe(i)
    })

    if out != "(<nil>, <nil>)\n" {
        t.Errorf("Unexpected string: %s", out)
    }

}
