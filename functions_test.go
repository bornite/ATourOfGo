package main

import (
    "testing"
)

func TestAdd (t *testing.T) {
    result := add(42, 13)
    if result != 55 {
        t.Fatalf("failed test. result=%d", result)
    }
}
