package main

import (
    "testing"
)

func TestAbs(t *testing.T) {

    v := Vertex{3, 4}
    a :=  v.Abs()
    if a != 5  {
        t.Fatalf("failed test1. a=%v", a)
    }

    a = AbsFunc(v)
    if a != 5  {
        t.Fatalf("failed test2. a=%v", a)
    }

    p := &Vertex{4, 3}
    a = p.Abs()
    if a != 5  {
        t.Fatalf("failed test3. a=%v", a)
    }

    a = AbsFunc(*p)
    if a != 5  {
        t.Fatalf("failed test4. a=%v", a)
    }

}
