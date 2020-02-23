package main

import (
    "testing"
)

func TestAbs(t *testing.T) {

    v := Vertex{3, 4}
    v.Scale(10)
    a := v.Abs()
    if a != 50  {
        t.Fatalf("failed test1. a=%v", a)
    }
}
