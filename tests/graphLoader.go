package tests

import "projekt2/graph"

func LoadTestGraphs() (smallGraph, mediumGraph, largeGraph graph.Graph) {
	smallGraph = graph.NewAdjMatrixGraph(55, 100000000)
	mediumGraph = graph.NewAdjMatrixGraph(170, 100000000)
	largeGraph = graph.NewAdjMatrixGraph(358, 0)
	graph.LoadGraphFromFile("ftv55.atsp", smallGraph)

	return
}
