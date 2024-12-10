package tsTuning

import "projekt2/graph"

func RunIterationsTuningTS() {
	smallGraph := graph.NewAdjMatrixGraph(55, 100000000)
	mediumGraph := graph.NewAdjMatrixGraph(170, 100000000)
	largeGraph := graph.NewAdjMatrixGraph(358, 0)
}
