package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := float64(0); i < 20; i++ {
		z := i
		z -= (z*z - x) / (2 * z)

		fmt.Println(z)
	}
	return z

}

func main() {
	Sqrt(2)
}
