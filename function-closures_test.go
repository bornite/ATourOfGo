package main

import (
    "testing"
)

func TestAdder(t *testing.T) {

    f := adder()
    a := 0
    for i := 0; i < 10; i++ {
        a = f(i)
    }
    if a != 45  {
        t.Fatalf("failed test. a=%d", a)
    }
}
