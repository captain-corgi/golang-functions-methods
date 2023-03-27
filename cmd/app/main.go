package main

import "fmt"

func main() {
	fn := GenDisplaceFn(10, 2, 1)

	fmt.Println(fn(3))
	fmt.Println(fn(5))
}

// GenDisplaceFn calculate s = ½ a t^2 + vot + so
func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return (a*t*t)/2 + v0*t + s0
	}
}
