package main

import (
    "testing"
)

func TestAbs(t *testing.T) {

    v := &Vertex{3, 4}
    a := v.Abs()
    if a != 5  {
        t.Fatalf("failed test1. a=%v", a)
    }

    v.Scale(5)
    a = v.Abs()
    if a != 25  {
        t.Fatalf("failed test2. a=%v", a)
    }

}
