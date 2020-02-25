package main

import (
    "testing"
)

func TestAbs(t *testing.T) {

    v := Vertex{3, 4}
    v.Scale(2)
    ScaleFunc(&v, 10)

    expect := Vertex{60, 80}
    if v != expect  {
        t.Fatalf("failed test1. v=%v", v)
    }

    p := &Vertex{4, 3}
    p.Scale(3)
    ScaleFunc(p, 8)

    expect = Vertex{96, 72}
    if *p != expect  {
        t.Fatalf("failed test2. *p=%v", *p)
    }

}
