package graph

import (
	"log"
	"math"
)

func angle(v1, v2 vector2D) float64 {
	dotProduct := v1[0]*v2[0] + v1[1]*v2[1]
	magnitudeV1 := math.Sqrt(v1[0]*v1[0] + v1[1]*v1[1])
	magnitudeV2 := math.Sqrt(v2[0]*v2[0] + v2[1]*v2[1])

	if magnitudeV1 == 0 || magnitudeV2 == 0 {
		return 0
	}

	cosTheta := dotProduct / (magnitudeV1 * magnitudeV2)

	if cosTheta < -1 {
		cosTheta = -1
	} else if cosTheta > 1 {
		cosTheta = 1
	}

	return math.Acos(cosTheta)
}

// TODO: ta saindo uns grafo esquisito com uns valores mto altos... tipo 5e+08 que não era pra existir...
func turnCost(v1, v2 vector2D) float64 {
	angle := angle(v1, v2)
	if math.IsNaN(angle) {
		log.Printf("NaN detected in angle calculation: v1=%v, v2=%v", v1, v2)
	}
	return float64(maxPenalti) * (angle / math.Pi)
}
