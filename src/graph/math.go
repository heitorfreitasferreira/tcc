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

//BUG: maxPenalti como time.Duration (int64) — float64(maxPenalti) só não dá overflow porque 1ns→1.0.
//     O TODO original reportava valores ~5e+08 no grafo, provavelmente NaN propagation ou bug de versão anterior.
//     Se maxPenalti fosse alterado para 1e9, o cast para nanossegundos daria 0, subestimando o turn cost.
func turnCost(v1, v2 vector2D) float64 {
	angle := angle(v1, v2)
	if math.IsNaN(angle) {
		log.Printf("NaN detected in angle calculation: v1=%v, v2=%v", v1, v2)
	}
	return float64(maxPenalti) * (angle / math.Pi)
}
