package tsTuning

import (
	"log"
	"projekt2/graph"
	"projekt2/solver/ts"
	"projekt2/tests"
	"projekt2/utils"
	"strconv"
	"time"
)

func RunIterationsTuningTS() {
	smallGraph, mediumGraph, largeGraph := tests.LoadTestGraphs()
	iterations := []int{1000, 2000, 3000}
	timeoutInNs := utils.MinutesToNanoSeconds(2)
	runSingleGraphIterTuning(smallGraph, iterations, timeoutInNs, "ts_iterations_small_")
	runSingleGraphIterTuning(mediumGraph, iterations, timeoutInNs, "ts_iterations_medium_")
	runSingleGraphIterTuning(largeGraph, iterations, timeoutInNs, "ts_iterations_large_")
}

func runSingleGraphIterTuning(g graph.Graph, iterations []int, timeoutInNs int64, fileOutName string) {
	results := make([][][]int64, len(iterations))
	for i, _ := range iterations {
		//0 - time, 1 - cost
		results[i] = make([][]int64, 2)
		for j := 0; j < 2; j++ {
			results[i][j] = make([]int64, 10)
		}
	}

	for i, it := range iterations {
		tsSolver := ts.NewTabuSearchATSPSolver(it, timeoutInNs, 10, "insert")
		tsSolver.SetGraph(g)
		tsSolver.SetStartVertex(0)
		for j := 0; j < 10; j++ {
			start := time.Now()
			_, weight := tsSolver.Solve()
			elapsed := time.Since(start)
			log.Println("Iteration: ", it, " Time: ", elapsed, " Weight: ", weight, " Graph size: ", g.GetVertexCount())
			results[i][0][j] = elapsed.Nanoseconds()
			results[i][1][j] = int64(weight)
		}
		utils.SaveTimesToCSVFile(results[i], fileOutName+strconv.Itoa(it)+"iterations_"+utils.GetDateForFilename()+".csv")
	}
}
