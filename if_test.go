package main

import (
    "testing"
)

func TestSqrt(t *testing.T) {
    a := sqrt(2)
    if a != "1.4142135623730951" {
        t.Fatalf("failed test. a=%v", a)
    }
}
