package pso

import (
	"math"
	"math/rand"
	"tcc/graph"
)

type particle struct {
	x            []float64
	v            []float64
	bestX        []float64
	clusterId    int
	makespan     float64
	bestMakespan float64
}

func newParticle(dimention, clusterId int) particle {

	return particle{
		x:            make([]float64, dimention),
		v:            make([]float64, dimention),
		bestX:        make([]float64, dimention),
		clusterId:    clusterId,
		bestMakespan: math.MaxFloat64,
	}
}

type Params struct {
	C1, C2, W float64
}

type Swarm struct {
	*graph.Graph
	Params
	particles     []particle
	bestByCluster [][]float64 // var bestInCluster []float64 = bestByCluster[clusterId]
	rng           rand.Rand
}

// func (sw Swarm) Optimize() SwarmStats {
// 	//TODO
// }

func (sw *Swarm) update(p *particle) {
	for i := 0; i < len(p.x); i++ {
		rnd1 := sw.rng.Float64()*2 - 1 // TODO: ver se é [-1,1] ou [0,1] https://ieeexplore.ieee.org/stamp/stamp.jsp?tp=&arnumber=1259748#page=1.54
		rnd2 := sw.rng.Float64()*2 - 1 // TODO: ver se é [-1,1] ou [0,1] https://ieeexplore.ieee.org/stamp/stamp.jsp?tp=&arnumber=1259748#page=1.54

		// Inércia
		p.v[i] *= sw.W
		// Cognitivo
		p.v[i] += sw.C1 * rnd1 * (p.bestX[i] - p.x[i])
		// Social
		p.v[i] += sw.C2 * rnd2 * (sw.bestByCluster[p.clusterId][i] - p.x[i])
		// Att a posição
		p.x[i] += p.v[i]
	}

	seq := discretize(p.x)

	if !sw.Graph.IsValidSolution(seq) {
		panic("Solução inválida")
	}

	p.makespan = sw.Graph.Makespan(seq)
}

func discretize(x []float64) []int {
	//TODO
	return []int{1, 2, 3}
}
