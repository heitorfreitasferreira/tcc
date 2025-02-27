package pso

import (
	"math/rand"
	"tcc/graph"
	"tcc/shared"
)

type Params struct {
	shared.HyperParams
	C1, C2, W float64
}

type Swarm struct {
	*graph.Graph
	Params
	particles []particle
	gBest     particle
	rng       rand.Rand
}

func (sw *Swarm) Optimize() *SwarmStats {
	sw.initialize()
	sw.evaluate()

	stats := newStats()
	for range sw.Iterations {
		sw.update()
		sw.evaluate()
		stats.addIterData(sw.gBest.bestMakespan, sw.gBest.bestX, sw.gBest.sequence)
	}

	return stats
}

func (sw *Swarm) update() {
	for _, p := range sw.particles {
		r1 := sw.rng.Float64()
		r2 := sw.rng.Float64()
		p.update(sw.W, sw.C1, sw.C2, r1, r2, &sw.gBest)
	}
}

func (sw *Swarm) evaluate() {
	for _, p := range sw.particles {
		p.makespan = sw.Makespan(p.sequence)
		if p.makespan < p.bestMakespan {
			p.bestMakespan = p.makespan
			copy(p.bestX, p.x)
			if p.bestMakespan < sw.gBest.bestMakespan {
				sw.gBest = p
			}
		}
	}
}

func (sw *Swarm) initialize() {
	if sw.Graph == nil {
		panic("Trying to initialize a swarm without a graph")
	}

	for range sw.PopulationSize {
		newParticle := newParticle(sw.NumberOfNodes())
		for i := range newParticle.x {
			newParticle.x[i] = sw.rng.Float64()
		}
		newParticle.setSequence()

		sw.particles = append(sw.particles, newParticle)
	}
}
