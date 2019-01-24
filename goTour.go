package main

import (
    "fmt"
)

"This "
func Sqrt(x float64) float64 {
    z:= 1.0
    for count := 0; count <= 10; count++ {
        z -= (z*z -x) / (2*x)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(20))
}
