package main

import (
    "testing"
    "bytes"
    "io"
    "os"
    "fmt"
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

    a := Person{"Arthur Dent", 42}
    z := Person{"Zaphod Beeblebrox", 9001}
    out := captureStdout(func() {
            fmt.Println(a, z)
    })
    if out != "Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)\n" {
        t.Errorf("test1 : Unexpected string: %s", out)
    }
}
