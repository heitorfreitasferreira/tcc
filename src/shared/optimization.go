package shared

type Improvement struct {
	Iteration    int
	Evaluation   int
	BestMakespan float64
	Delta        float64
	BestSequence []int
}

type OptimizationResult struct {
	BestSequence        []int
	BestMakespan        float64
	IterationsCompleted int
	Evaluations         int
	Improvements        []Improvement
}
