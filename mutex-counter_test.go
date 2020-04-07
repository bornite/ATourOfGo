package main

import (
    "testing"
    "time"
)

func TestMain(t *testing.T) {

    c := SafeCounter{v: make(map[string]int)}
    for i := 0; i < 1000; i++ {
        go c.Inc("somekey")
    }
    for i := 0; i < 500; i++ {
        go c.Inc("thiskey")
    }

    time.Sleep(time.Second)
    out := c.Value("somekey")
    if out != 1000 {
      t.Errorf("test1 : Unexpected return value: %v", out)
    }

    out = c.Value("thiskey")
    if out != 500 {
      t.Errorf("test2 : Unexpected return value: %v", out)
    }
}
