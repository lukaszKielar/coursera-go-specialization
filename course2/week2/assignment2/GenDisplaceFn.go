package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a float64, vo float64, so float64) func(float64) float64 {
	f := func(t float64) float64 {
		return (0.5 * a * math.Pow(t, 2)) + (vo * t) + so
	}
	return f
}

func main() {
	var acceleration, initialVelocity, initialDisplacement, time float64

	fmt.Print("Enter acceleration: ")
	fmt.Scanln(&acceleration)

	fmt.Print("Enter initial velocity: ")
	fmt.Scanln(&initialVelocity)

	fmt.Print("Enter initial displacement: ")
	fmt.Scanln(&initialDisplacement)

	fmt.Println(acceleration, initialVelocity, initialDisplacement)
	f := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Print("Enter time: ")
	fmt.Scanln(&time)

	res := f(time)
	fmt.Printf("Displacement after time t=%.f: %.f", time, res)
	fmt.Println()
}
