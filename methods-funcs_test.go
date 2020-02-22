package main

import (
    "testing"
)

func TestAbs(t *testing.T) {

    v := Vertex{3, 4}
    a := Abs(v)
    if a != 5  {
        t.Fatalf("failed test. a=%v", a)
    }

    v = Vertex{5, 12}
    a = Abs(v)
    if a != 13  {
        t.Fatalf("failed test. a=%v", a)
    }
}
