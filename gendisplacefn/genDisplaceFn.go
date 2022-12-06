package main

import (
	"fmt"
	"math"
	"time"
)

func GenDisplaceFn(a, v_o, s_o float64) func(float64) float64 {
	fn := func(t float64) float64 {
		s := 0.5*a*math.Pow(t, 2) + v_o*t + s_o
		for i := 0; i < int(t); i++ {
			time.Sleep(1000 * time.Millisecond)
		}
		return s
	}
	return fn
}

func main() {
	genDisplaceFnVar := GenDisplaceFn(10, 2, 1)
	fmt.Print(genDisplaceFnVar(5))
}
