package ts

import (
	"log"
	"math"
	"projekt2/graph"
	"projekt2/solver"
	"time"
)

type TabuSearchATSPSolver struct {
	graph       graph.Graph
	startVertex int
	iterations  int   // maksymalna liczba iteracji
	timeout     int64 // limit czasu w nanosekundach (-1 oznacza brak limitu)
	startTime   time.Time
	tabuTenure  int // ile iteracji ruch pozostaje tabu
}

// SetGraph ustawia graf dla solvera
func (t *TabuSearchATSPSolver) SetGraph(g graph.Graph) {
	t.graph = g
}

// GetGraph zwraca przypisany graf
func (t *TabuSearchATSPSolver) GetGraph() graph.Graph {
	return t.graph
}

// SetStartVertex ustawia wierzchołek startowy
func (t *TabuSearchATSPSolver) SetStartVertex(startVertex int) {
	t.startVertex = startVertex
}

// SetTimeout ustawia czas wykonania w nanosekundach, np. 1s = 1e9 ns
func (t *TabuSearchATSPSolver) SetTimeout(timeout int64) {
	t.timeout = timeout
}

// GetTimeout zwraca aktualnie ustawiony czas wykonania
func (t *TabuSearchATSPSolver) GetTimeout() int64 {
	return t.timeout
}

// NewTabuSearchATSPSolver tworzy nowy solver z podanymi parametrami
// iterations - maksymalna liczba iteracji
// timeout - limit czasu w nanosekundach (-1 oznacza brak limitu)
// tabuTenure - ile iteracji dany ruch pozostaje tabu
func NewTabuSearchATSPSolver(iterations int, timeout int64, tabuTenure int) solver.ATSPSolver {
	return &TabuSearchATSPSolver{
		iterations: iterations,
		timeout:    timeout,
		tabuTenure: tabuTenure,
	}
}

// calculateCost oblicza koszt ścieżki, zakładając, że path już kończy się na startVertex
func (t *TabuSearchATSPSolver) calculateCost(path []int) int {
	return t.graph.CalculatePathWeight(path)
}

// getAllNeighbors generuje wszystkie sąsiednie rozwiązania poprzez zamianę dwóch wierzchołków (oprócz startVertex na początku i końcu)
func (t *TabuSearchATSPSolver) getAllNeighbors(currentPath []int) []struct {
	path []int
	i, j int
	cost int
} {
	vertexCount := len(currentPath)
	neighbors := []struct {
		path []int
		i, j int
		cost int
	}{}

	if vertexCount <= 3 {
		// brak sensownych sąsiadów, bo tylko start, jeden wierzchołek i start
		return neighbors
	}

	for i := 1; i < vertexCount-1; i++ {
		for j := i + 1; j < vertexCount-1; j++ {
			newPath := make([]int, vertexCount)
			copy(newPath, currentPath)
			// zamiana dwóch wierzchołków
			newPath[i], newPath[j] = newPath[j], newPath[i]

			c := t.calculateCost(newPath)
			neighbors = append(neighbors, struct {
				path []int
				i, j int
				cost int
			}{path: newPath, i: i, j: j, cost: c})
		}
	}

	return neighbors
}

// Solve - implementacja metody rozwiązywania problemu ATSP przy użyciu Tabu Search
func (t *TabuSearchATSPSolver) Solve() ([]int, int) {
	vertexCount := t.graph.GetVertexCount()
	if vertexCount == 0 {
		return nil, -1
	}

	// start czasu
	t.startTime = time.Now()

	// Pobieramy początkowe rozwiązanie zachłanne
	currentSolution := t.graph.GetHamiltonianPathGreedy(t.startVertex)
	currentCost := t.calculateCost(currentSolution)

	// Ustawiamy najlepsze rozwiązanie
	bestSolution := make([]int, len(currentSolution))
	copy(bestSolution, currentSolution)
	bestCost := currentCost

	// Tabu lista: będziemy przechowywać czas tabu dla par indeksów (i,j)
	// Możemy użyć 2D slice, gdzie tabuList[i][j] > 0 oznacza, że ruch zamiany i,j jest tabu
	tabuList := make([][]int, vertexCount)
	for i := 0; i < vertexCount; i++ {
		tabuList[i] = make([]int, vertexCount)
	}

	// Główna pętla Tabu Search
	for iteration := 0; iteration < t.iterations; iteration++ {
		// Sprawdzenie limitu czasu
		if t.timeout != -1 {
			elapsed := time.Since(t.startTime).Nanoseconds()
			if elapsed >= t.timeout {
				log.Println("Przekroczono limit czasu. Kończenie algorytmu Tabu Search.")
				break
			}
		}

		// Generujemy wszystkich sąsiadów
		neighbors := t.getAllNeighbors(currentSolution)
		if len(neighbors) == 0 {
			// brak sąsiadów, kończymy
			break
		}

		// Wybór najlepszego sąsiada nie będącego tabu lub łamiącego tabu w przypadku poprawy globalnego optimum
		var chosenNeighbor struct {
			path []int
			i, j int
			cost int
		}
		chosenNeighbor.cost = math.MaxInt

		for _, neigh := range neighbors {
			isTabu := (tabuList[neigh.i][neigh.j] > 0 || tabuList[neigh.j][neigh.i] > 0)
			// Wybieramy najlepszego sąsiada:
			// 1. Nie tabu i o minimalnym koszcie
			// lub
			// 2. Tabu, ale poprawia globalne optimum (aspiration criterion)
			if neigh.cost < chosenNeighbor.cost {
				if !isTabu || (isTabu && neigh.cost < bestCost) {
					chosenNeighbor = neigh
				}
			}
		}

		// Jeśli nie znaleźliśmy lepszego sąsiada, to znaczy wszyscy są tabu i żaden nie poprawia bestCost
		// Wówczas wybieramy mimo wszystko najlepszego (nawet tabu), aby nie utknąć
		if chosenNeighbor.cost == math.MaxInt {
			// wybieramy po prostu najlepszego sąsiada bez względu na tabu
			bestNonImproving := neighbors[0]
			for _, neigh := range neighbors {
				if neigh.cost < bestNonImproving.cost {
					bestNonImproving = neigh
				}
			}
			chosenNeighbor = bestNonImproving
		}

		// Przejście do wybranego sąsiada
		prevI, prevJ := chosenNeighbor.i, chosenNeighbor.j
		currentSolution = chosenNeighbor.path
		currentCost = chosenNeighbor.cost

		// Aktualizacja najlepszego rozwiązania
		if currentCost < bestCost {
			bestCost = currentCost
			copy(bestSolution, currentSolution)
		}

		// Aktualizacja listy tabu:
		//  - ruch (prevI, prevJ) jest tabu przez tabuTenure iteracji
		tabuList[prevI][prevJ] = t.tabuTenure
		tabuList[prevJ][prevI] = t.tabuTenure

		// Zmniejszamy karencje tabu dla wszystkich ruchów
		for i := 0; i < vertexCount; i++ {
			for j := 0; j < vertexCount; j++ {
				if tabuList[i][j] > 0 {
					tabuList[i][j]--
				}
			}
		}
	}

	log.Println("Zakończono Tabu Search. Najlepszy znaleziony koszt:", bestCost)
	return bestSolution, bestCost
}
