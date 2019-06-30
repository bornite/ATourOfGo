package main

import (
    "testing"
)

func TestPow(t *testing.T) {
    a := pow(3, 2, 10)
    if a != 9 {
        t.Fatalf("failed test. a=%v", a)
    }

    b := pow(3, 3, 20)
    if b != 20 {
        t.Fatalf("failed test. b=%v", b)
    }
}
