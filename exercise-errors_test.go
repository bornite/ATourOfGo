package main

import (
    "testing"
    "math"
)

func TestSqrt(t *testing.T) {
    a, err := Sqrt(5)
    if a != math.Sqrt(5) {
        t.Fatalf("failed test. a=%f, err=%v", a, err)
    }
}
