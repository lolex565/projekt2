package tsAmountTests

import (
	"log"
	"projekt2/graph"
	"projekt2/solver/ts"
	"projekt2/utils"
	"runtime"
	"time"
)

func RunTSAmountTests(startVertexCount int, vertexCountStep int) {
	vertexCount := startVertexCount
	noEdgeValue := -1
	timeoutInNs := utils.MinutesToNanoSeconds(5)
	tsSolver := ts.NewTabuSearchATSPSolver(3000, timeoutInNs, 10, "insert")
	tsSolver.SetStartVertex(0)
	results := make([][]int64, 0)
out:
	for {
		tempResults := make([]int64, 0)
		for i := 0; i < 50; i++ {
			g := graph.NewAdjMatrixGraph(vertexCount, noEdgeValue)
			graph.GenerateRandomGraph(g, vertexCount, -1, 100)
			tsSolver.SetGraph(g)
			startTime := time.Now()
			path, weight := tsSolver.Solve()
			elapsed := time.Since(startTime)
			log.Println("Wierzchołki:", vertexCount, "Czas:", elapsed, "Waga:", weight)
			log.Println(g.PathWithWeightsToString(path))
			if elapsed.Nanoseconds() > timeoutInNs {
				log.Println("Testy przekraczają timeout, zatrzymano na ilości wierzchołków:", vertexCount)
				break out
			}
			tempResults = append(tempResults, elapsed.Nanoseconds())
			runtime.GC()
		}
		results = append(results, tempResults)
		vertexCount += vertexCountStep
	}
	utils.SaveTimesToCSVFile(results, "ts_amount_tests_"+utils.GetDateForFilename()+".csv")
}
