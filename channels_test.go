package main

import (
    "testing"
)

func TestSum(t *testing.T) {

    s := []int{1, 2, 3, 4, 5}

    c := make(chan int)
    go sum(s, c)

    out := <- c
    if out != 15 {
      t.Errorf("test1 : Unexpected value: %v", out)
    }
}
