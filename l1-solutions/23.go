package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func newPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (this *Point) Distance(other *Point) float64 {
	dx := this.x - other.x
	dy := this.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := newPoint(0.0, 0.0)
	p2 := newPoint(4.0, 3.0)
	fmt.Println(p1.Distance(p2))
}
