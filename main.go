package main

import (
	"fmt"
	"log"
	"projekt2/graph"
	"projekt2/solver/sa"
	"projekt2/solver/ts"
	"projekt2/utils"
	"time"
)

func main() {
	g := new(graph.AdjMatrixGraph)
	err := graph.LoadGraphFromFile("rbg443.atsp", g, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(g.ToString())
	g.SetNoEdgeValue(0)
	newOsaATSPSolver := sa.NewOtherSimulatedAnnealingATSPSolver(10000, 1e-20, 0.99, 1000, -1)
	newOsaATSPSolver.SetGraph(g)
	startTime := time.Now()
	log.Println("Start time:", startTime)
	path, weight := newOsaATSPSolver.Solve()
	log.Println("Zajęło:", time.Since(startTime))
	fmt.Println(g.PathWithWeightsToString(path))
	fmt.Println(weight)

	newTsATSPSolver := ts.NewTabuSearchATSPSolver(1000, utils.MinutesToNanoSeconds(5), 10)
	newTsATSPSolver.SetGraph(g)
	startTime = time.Now()
	log.Println("Start time:", startTime)
	path, weight = newTsATSPSolver.Solve()
	log.Println("Zajęło:", time.Since(startTime))
	fmt.Println(g.PathWithWeightsToString(path))
	fmt.Println(weight)
}
