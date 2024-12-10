package saAmountTests

import (
	"log"
	"projekt2/graph"
	"projekt2/solver/sa"
	"projekt2/utils"
	"runtime"
	"time"
)

func RunSAAmountTests(sizes []int) {
	noEdgeValue := -1
	timeoutInNs := utils.SecondsToNanoSeconds(30)
	saSolver := sa.NewSimulatedAnnealingATSPSolver(10000, 1e-3, 0.95, 1000, timeoutInNs)
	saSolver.SetStartVertex(0)
	results := make([][]int64, 0)
out:
	for _, vertexCount := range sizes {
		tempResults := make([]int64, 0)
		for i := 0; i < 10; i++ {
			g := graph.NewAdjMatrixGraph(vertexCount, noEdgeValue)
			graph.GenerateRandomGraph(g, vertexCount, -1, 100)
			saSolver.SetGraph(g)
			startTime := time.Now()
			_, weight := saSolver.Solve()
			elapsed := time.Since(startTime)
			log.Println("Wierzchołki:", vertexCount, "Czas:", elapsed, "Waga:", weight)
			if elapsed.Nanoseconds() > timeoutInNs {
				log.Println("Testy przekraczają timeout, zatrzymano na ilości wierzchołków:", vertexCount)
				break out
			}
			tempResults = append(tempResults, elapsed.Nanoseconds())
			runtime.GC()
		}
		results = append(results, tempResults)
	}
	utils.SaveTimesToCSVFile(results, "sa_amount_tests_"+utils.GetDateForFilename()+".csv")
}
