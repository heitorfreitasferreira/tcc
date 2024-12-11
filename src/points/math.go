package points

import "math"

func (p Coordinate2D) EuclideanDistance(other Coordinate2D) float64 {
	return math.Sqrt(math.Pow(p[0]-other[0], 2) + math.Pow(p[1]-other[1], 2))
}

// Retorna angulo em radiano (0..PI)
func (p Coordinate2D) RadiansRadius(other Coordinate2D) float64 {
	dx := other[0] - p[0]
	dy := other[1] - p[1]
	return math.Atan2(dy, dx)
}

func (p Coordinate2D) vector(other Coordinate2D) (float64, float64) {
	return p[0] - other[0], p[1] - other[1]
}

func (p Coordinate2D) Vector(other Coordinate2D) (float64, float64) {
	dx, dy := p.vector(other)
	magnitude := math.Sqrt(dx*dx + dy*dy)
	if magnitude == 0 {
		return 0, 0
	}
	return dx / magnitude, dy / magnitude
}
