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

func TestMain(t *testing.T) {

    out := captureStdout(func() {
            main()
    })

    expected := `
0
1
1
2
3
5
8
13
21
34
quit
`
    if out == expected {
      t.Errorf("test1 : Unexpected string: %s", out)
    }
}
