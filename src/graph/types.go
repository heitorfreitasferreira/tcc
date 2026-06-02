package graph

import (
	"time"
)

type metersPerSecond float64

const droneSpeed metersPerSecond = 1
const maxPenalti time.Duration = 1 //BUG: time.Duration (int64) como escalar float64 — type confusion. float64(maxPenalti) funciona porque 1ns→1.0, mas quebra se for alterado.

type Graph [][][]float64 // Tempo de percurso dado os pontos [anterior][atual][proximo]

type vector2D [2]float64
