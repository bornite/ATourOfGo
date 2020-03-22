package main

import (
    "testing"
    "bytes"
    "io"
    "os"
    "image"
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

func TestMain(t *testing.T) {

    m := image.NewRGBA(image.Rect(0, 0, 100, 100))

    out := captureStdout(func() {
            fmt.Println(m.Bounds())
    })
    if out != "(0,0)-(100,100)\n" {
        t.Errorf("test1 : Unexpected string: %s", out)
    }

}
