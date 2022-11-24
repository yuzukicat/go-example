package genDisplaceFn

import (
	"math"
)

func GenDisplaceFn(a, v_o, s_o float64) func(float64) float64 {
	fn := func(t float64) float64 {
		s := 0.5*a*math.Pow(t, 2) + v_o*t + s_o
		return s
	}
	return fn
}
