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

func main() {
	c1 := Circle{2}
	fmt.Println(c1.erea())
}

// http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/

