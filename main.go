package main

import (
	"fmt"
	"projekt2/graph"
	"projekt2/solver/bnb"
	"projekt2/solver/sa"
	"projekt2/solver/ts"
)

func main() {
	g := graph.NewAdjMatrixGraph(20, -1)
	graph.GenerateRandomGraph(g, 20, -1, 1000)
	fmt.Println(g.ToString())
	greedyHamiltonianPath := g.GetHamiltonianPathGreedy(0)
	fmt.Println(greedyHamiltonianPath)
	fmt.Println(g.PathWithWeightsToString(greedyHamiltonianPath))
	fmt.Println(g.CalculatePathWeight(greedyHamiltonianPath))
	newBnbSolver := new(bnb.BNBATSPSolver)
	newBnbSolver.SetGraph(g)
	newBnbSolver.SetStartVertex(0)
	fmt.Println(newBnbSolver.Solve())
	newSaSolver := sa.NewSimulatedAnnealingATSPSolver(5000, 1e-6, 0.99999995, 100000000000, -1)
	newSaSolver.SetGraph(g)
	newSaSolver.SetStartVertex(0)
	fmt.Println(newSaSolver.Solve())
	newTsSolver := ts.NewTabuSearchATSPSolver(1000000, -1, 10)
	newTsSolver.SetGraph(g)
	newTsSolver.SetStartVertex(0)
	fmt.Println(newTsSolver.Solve())

}
