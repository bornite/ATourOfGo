package main

import (
    "testing"
)

func TestFibonacci(t *testing.T) {

    a, f := 0, fibonacci()
    for i := 0; i < 10; i++ {
        a = f()
    }
    if a != 34  {
        t.Fatalf("failed test. a=%d", a)
    }

    for i := 0; i < 20; i++ {
        a = f()
    }
    if a != 514229  {
        t.Fatalf("failed test. a=%d", a)
    }
}
