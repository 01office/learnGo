package letsgo

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) erea() float64 {
	return c.radius * c.radius * math.Pi
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

