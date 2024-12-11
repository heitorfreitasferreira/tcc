package graph

import (
	"time"
)

type metersPerSecond float64

const droneSpeed metersPerSecond = 1
const maxPenalti time.Duration = 1

type Graph [][][]float64 // Tempo de percurso dado os pontos [anterior][atual][proximo]

type vector2D [2]float64
