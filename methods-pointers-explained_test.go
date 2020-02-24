package main

import (
    "testing"
)

func TestAbs(t *testing.T) {

    v := Vertex{3, 4}
    a := Abs(v)
    if a != 5  {
        t.Fatalf("failed test1. a=%v", a)
    }
    Scale(&v, 10)
    a = Abs(v)
    if a != 50  {
        t.Fatalf("failed test2. a=%v", a)
    }

    v = Vertex{12, 5}
    a = Abs(v)
    if a != 13  {
        t.Fatalf("failed test3. a=%v", a)
    }
    Scale(&v, 0.1)
    a = Abs(v)
    if a != 1.3  {
        t.Fatalf("failed test4. a=%v", a)
    }

}
