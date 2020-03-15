package main

import (
    "fmt"
    "math"
)

type ErrNegativeSquareRoot float64

func (e ErrNegativeSquareRoot) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSquareRoot(x)
    }

    z := 1.0
    for diff := 1.0; math.Abs(diff) > 1e-10; {
        diff = ((z * z) - x) / (2 * z)
        z = z - diff
    }
    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
