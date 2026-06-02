package aco

import (
	"math"
)

type ant struct {
	seq     []int
	visited []bool

	lk float64
}

func (aco *ACO) updatePheromones(ants []ant) {
	// Evaporação
	for i := range aco.pheromones {
		for j := range aco.pheromones[i] {
			for k := range aco.pheromones[i][j] { //BUG: varre N³ inteiro (incluindo i==j, j==k, i==k que são inválidos e estão sempre em 0). Deveria iterar só triplas válidas (i!=j && j!=k && i!=k).
				aco.pheromones[i][j][k] *= (1 - aco.Rho)
			}
		}
	}

	// Depósito por formigas (estilo Ant System — todas depositam)
	for _, ant := range ants { //NOTE: (pesquisa) depósito de todas as formigas (AS). Global-best ACS convergiria mais rápido mas arrisca estagnação. MMAS usaria só a melhor iteração + melhor global.
		if ant.lk == 0 {
			continue
		}
		delta := aco.Q / ant.lk
		for idx := 2; idx < len(ant.seq); idx++ { //NOTE: idx começa em 2 porque seq[0]=depot, seq[1]=first. O triple (depot, first, second) é o primeiro depositado. A aresta depot→first nunca é depositada como triple (prev,curr,next).
			prev := ant.seq[idx-2]
			curr := ant.seq[idx-1]
			next := ant.seq[idx]
			aco.pheromones[prev][curr][next] += delta
		}
	}
}

func (aco *ACO) selectNextNode(prev, curr int, visited []bool) int {
	probabilities := make([]float64, len(aco.Graph))
	total := 0.0

	// Construindo as probs
	for next := range aco.Graph[prev][curr] { //NOTE: (pesquisa) sem lista candidata — nearest-neighbor reduziria branching factor em problemas grandes
		if visited[next] || next == prev || next == curr {
			continue
		}

		// η = 1/custo (heurística)
		heuristic := 1.0 / aco.Graph[prev][curr][next] //NOTE: G patológico (bug maxPenalti) corrompe η para ~0, transição nunca escolhida
		prob := math.Pow(aco.pheromones[prev][curr][next], aco.Alpha) *
			math.Pow(heuristic, aco.Beta) //NOTE: (pesquisa) sem τ_min/τ_max — MMAS com limites [τ_min, τ_max] mitigaria diluição 3D e estagnação

		probabilities[next] = prob
		total += prob
	}

	if total == 0 { //BUG: nenhum candidato viável com probabilidade > 0 — retorna -1, walk() faz break, tour incompleto contribui 0 feromônio
		return -1
	}

	// Roleta
	cutoff := aco.rng.Float64() * total
	sum := 0.0
	for next, prob := range probabilities {
		if sum += prob; sum >= cutoff {
			return next
		}
	}
	return -1
}

func (aco *ACO) walk() ant {
	n := len(aco.Graph)
	ant := ant{
		seq:     make([]int, 0, n),
		visited: make([]bool, n),
		lk:      0.0,
	}

	current := 0
	ant.seq = append(ant.seq, current)
	ant.visited[current] = true

	next := aco.rng.Intn(n) //BUG: primeiro passo é sorteado uniformemente (sem heurística, sem feromônio). O loop de depósito (idx=2) nunca reforça esta aresta como triple (prev,curr,next).
	for next == current {
		next = aco.rng.Intn(n)
	}
	ant.seq = append(ant.seq, next)
	ant.visited[next] = true

	for len(ant.seq) < n {
		prev := ant.seq[len(ant.seq)-2]
		curr := ant.seq[len(ant.seq)-1]
		next := aco.selectNextNode(prev, curr, ant.visited)
		if next == -1 { //BUG: tour incompleto (len(seq) < n). Formiga contribui 0 feromônio — perda de exploração.
			break
		}

		ant.seq = append(ant.seq, next)
		ant.visited[next] = true
	}

	if len(ant.seq) == n {
		ant.lk = aco.Makespan(aco.pathWithoutOrigin(ant.seq))
	}
	return ant
}
