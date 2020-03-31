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

func TestFibonacci(t *testing.T) {

    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()

    out := captureStdout(func() {
        fibonacci(c, quit)
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
