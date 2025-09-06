package taskTwentyFour

import "math"

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) Distance(other *Point) float64 {
	dx := math.Abs(p.x - other.x)
	dy := math.Abs(p.y - other.y)
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}
