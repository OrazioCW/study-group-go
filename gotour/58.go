package main

import (
    "fmt"
    "math"
)

const delta = 1e-10

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, ErrNegativeSqrt(f)
    }
    z := f
    for {
        n := z - (z*z-f)/(2*z)
        if math.Abs(n-z) < delta {
            break
        }
        z = n
    }
    return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}


func main() {
        fmt.Println(Sqrt(2))
        fmt.Println(Sqrt(-2))
}
