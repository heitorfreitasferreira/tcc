package graph

// Makespan... calcula o tempo da solução (caminho), deve conter todos os números de 1 até o último ponto
// o drone sempre sai da origem (0), e volta ao destino (0)
// Logo um order= {1, 4, 3, 2, 4, 5} indica que o drone vai começar no ponto 0, vai para o 1, dps
// para o 4, até chegar no ponto 5, após isso é considerado que o mesmo volta para o 0
func (g Graph) Makespan(order []int) float64 {
	lastNode := 0
	currNode := 0

	makespan := 0.0
	for _, nextNode := range order {
		makespan += g[lastNode][currNode][nextNode]
		lastNode = currNode
		currNode = nextNode
	}

	makespan += g[lastNode][currNode][0] // Voltando para a origem
	return makespan
}

func (g Graph) IsValidSolution(solution []int) bool {
	numberOfNodes := len(g)
	visited := make([]bool, numberOfNodes)

	visited[0] = true // é suposto que sempre sai do primeiro nó e volta para o mesmo
	for _, node := range solution {
		visited[node] = true
	}

	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}

func (g Graph) NumberOfNodes() int {
	return len(g)
}
