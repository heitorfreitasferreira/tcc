package graph

type metersPerSecond float64

const droneSpeed metersPerSecond = 1
const maxPenalti float64 = 1

type Graph [][][]float64 // Tempo de percurso dado os pontos [anterior][atual][proximo]

type vector2D [2]float64
