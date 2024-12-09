package main

import (
	"fmt"
	"projekt2/graph"
	"projekt2/menu"
	"projekt2/solver/sa"
	"projekt2/utils"
)

func main() {
	g := new(graph.AdjMatrixGraph)
	err := graph.LoadGraphFromFile("rbg358.atsp", g, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(g.ToString())
	g.SetNoEdgeValue(0)
	programMenu := menu.NewDefaultMenu(g)
	saSolver := sa.NewSimulatedAnnealingATSPSolver(10000, 1e-9, 0.999999995, 10000000000, utils.MinutesToNanoSeconds(10))
	programMenu.SetSaATSPSolver(saSolver)
	programMenu.GetSaATSPSolverPtr().SetGraph(g)
	programMenu.GetSaATSPSolverPtr().SetStartVertex(0)
	path, weight := programMenu.GetSaATSPSolverPtr().Solve()
	fmt.Println(g.PathWithWeightsToString(path))
	fmt.Println(weight)

}
