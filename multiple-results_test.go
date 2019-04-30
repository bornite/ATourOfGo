package main

import (
    "testing"
)

func TestSwap(t *testing.T) {
    a, b := swap("XYZ", "12345678")
    if a != "12345678" || b != "XYZ" {
        t.Fatalf("failed test. a=%s, b=%s", a, b)
    }

}
