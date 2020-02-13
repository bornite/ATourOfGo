package main

import (
    "testing"
    "math"
)

func TestCompute(t *testing.T) {

    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }

    a := compute(hypot)
    if a != 5 {
        t.Fatalf("failed test. a=%f", a)
    }
}
