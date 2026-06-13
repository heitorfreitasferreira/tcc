package graph

func (g *Graph) Makespan(order []int) float64 {
	lastNode := 0
	currNode := 0
	makespan := 0.0
	n := g.N
	for _, nextNode := range order {
		makespan += g.Data[lastNode*n*n+currNode*n+nextNode]
		lastNode = currNode
		currNode = nextNode
	}
	makespan += g.Data[lastNode*n*n+currNode*n+0]
	return makespan
}
