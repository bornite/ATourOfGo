package namedresults

import (
    "testing"
)

func TestSplit(t *testing.T) {
    a, b := split(17)
    if a != 7 || b != 10 {
        t.Fatalf("failed test. a=%d, b=%d", a, b)
    }
}
