package graph

type metersPerSecond float64

const droneSpeed metersPerSecond = 1
const maxPenalti float64 = 1

type Graph struct {
	N    int
	Data []float64
}

type vector2D [2]float64

func (g *Graph) At(prev, curr, next int) float64 {
	return g.Data[prev*g.N*g.N+curr*g.N+next]
}

func (g *Graph) Set(prev, curr, next int, v float64) {
	g.Data[prev*g.N*g.N+curr*g.N+next] = v
}

func (g *Graph) NumberOfNodes() int {
	return g.N
}

func (g *Graph) IsValidSolution(solution []int) bool {
	visited := make([]bool, g.N)
	visited[0] = true
	for _, node := range solution {
		if node < 0 || node >= g.N {
			return false
		}
		visited[node] = true
	}
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}
