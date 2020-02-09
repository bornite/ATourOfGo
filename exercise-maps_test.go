package main

import (
    "testing"
)

type TestCase struct {
    in   string
    want map[string]int
}

var tests = []TestCase {
    {"I am learning Go!", map[string]int{
        "I": 1, "am": 1, "learning": 1, "Go!": 1,
    }},
    {"The quick brown fox jumped over the lazy dog.", map[string]int{
        "The": 1, "quick": 1, "brown": 1, "fox": 1, "jumped": 1,
        "over": 1, "the": 1, "lazy": 1, "dog.": 1,
    }},
    {"I ate a donut. Then I ate another donut.", map[string]int{
        "I": 2, "ate": 2, "a": 1, "donut.": 2, "Then": 1, "another": 1,
    }},
    {"A man a plan a canal panama.", map[string]int{
        "A": 1, "man": 1, "a": 2, "plan": 1, "canal": 1, "panama.": 1,
    }},
}

func TestWordCount(t *testing.T) {
    for _, c := range tests {
        got := WordCount(c.in)
        if len(c.want) != len(got) {
            t.Fatalf("failed len test. got=%v, want=%v", len(got), len(c.want))
        } else {
            for k := range c.want {
                if c.want[k] != got[k] {
                    t.Fatalf("failed len test. got=%v, want=%v", got[k], c.want[k])
                }
            }
        }
    }
}
