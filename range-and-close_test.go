package main

import (
    "testing"
)


func Test(t *testing.T) {

    c := make(chan int, 50)
    n := cap(c)
    go fibonacci(n, c)
    out := 0
    for i := range c {
        out = i
    }

    if out != 7778742049  {
      t.Errorf("test1 : Unexpected string: %d", out)
    }
}
