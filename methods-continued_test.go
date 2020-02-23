package main

import (
    "testing"
)

func TestAbs(t *testing.T) {

    f := MyFloat(-2.0)
    if f != -2.0  {
        t.Fatalf("failed test1. f=%v", f)
    }

    f = MyFloat(2)
    if f != 2  {
        t.Fatalf("failed test2. f=%v", f)
    }
}
