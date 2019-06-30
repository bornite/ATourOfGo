package main

import (
    "testing"
    "math"
)

func TestSqrt(t *testing.T) {
    a := Sqrt(2)
    b := math.Sqrt(2)
    if math.Abs(a-b) > 1.0e-10 {
        t.Fatalf("failed test. a=%v, b=%v, |a-b|=%v", a, b, math.Abs(a-b))
    }
}
