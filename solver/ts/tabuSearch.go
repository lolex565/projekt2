package ts

import (
	"fmt"
	"log"
	"math"
	"projekt2/graph"
	"time"
)

const (
	NeighborhoodSwap   = "swap"
	NeighborhoodInsert = "insert"
)

// TsATSPSolver struktura solvera Tabu Search
type TsATSPSolver struct {
	graph              graph.Graph
	startVertex        int
	iterations         int   // Maksymalna liczba iteracji
	timeout            int64 // Limit czasu w nanosekundach (-1 oznacza brak limitu)
	startTime          time.Time
	tabuTenure         int    // Ile iteracji ruch pozostaje tabu
	neighborhoodMethod string // Metoda sąsiedztwa: "swap" lub "insert"
}

// SetGraph ustawia graf dla solvera
func (t *TsATSPSolver) SetGraph(g graph.Graph) {
	t.graph = g
}

// GetGraph zwraca przypisany graf
func (t *TsATSPSolver) GetGraph() graph.Graph {
	return t.graph
}

// SetStartVertex ustawia wierzchołek startowy
func (t *TsATSPSolver) SetStartVertex(startVertex int) {
	t.startVertex = startVertex
}

// SetTimeout ustawia czas wykonania w nanosekundach, np. 1s = 1e9 ns
func (t *TsATSPSolver) SetTimeout(timeout int64) {
	t.timeout = timeout
}

// GetTimeout zwraca aktualnie ustawiony czas wykonania
func (t *TsATSPSolver) GetTimeout() int64 {
	return t.timeout
}

// SetNeighborhoodMethod ustawia metodę sąsiedztwa
func (t *TsATSPSolver) SetNeighborhoodMethod(method string) error {
	if method != NeighborhoodSwap && method != NeighborhoodInsert {
		return fmt.Errorf("nieprawidłowa metoda sąsiedztwa: %s", method)
	}
	t.neighborhoodMethod = method
	return nil
}

// GetNeighborhoodMethod zwraca aktualnie ustawioną metodę sąsiedztwa
func (t *TsATSPSolver) GetNeighborhoodMethod() string {
	return t.neighborhoodMethod
}

// NewTabuSearchATSPSolver tworzy nowy solver Tabu Search z podanymi parametrami
func NewTabuSearchATSPSolver(iterations int, timeout int64, tabuTenure int) TsATSPSolver {
	return TsATSPSolver{
		iterations:         iterations,
		timeout:            timeout,
		tabuTenure:         tabuTenure,
		neighborhoodMethod: NeighborhoodSwap, // Domyślna metoda
	}
}

// calculateCost oblicza koszt ścieżki, zakładając, że path już kończy się na startVertex
func (t *TsATSPSolver) calculateCost(path []int) int {
	return t.graph.CalculatePathWeight(path)
}

// getAllNeighbors generuje wszystkie sąsiednie rozwiązania na podstawie wybranej metody sąsiedztwa
func (t *TsATSPSolver) getAllNeighbors(currentPath []int) []struct {
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
		// Brak sensownych sąsiadów, bo tylko start, jeden wierzchołek i start
		return neighbors
	}

	switch t.neighborhoodMethod {
	case NeighborhoodSwap:
		for i := 1; i < vertexCount-1; i++ {
			for j := i + 1; j < vertexCount-1; j++ {
				newPath := make([]int, vertexCount)
				copy(newPath, currentPath)
				// Zamiana dwóch wierzchołków
				newPath[i], newPath[j] = newPath[j], newPath[i]

				c := t.calculateCost(newPath)
				neighbors = append(neighbors, struct {
					path []int
					i, j int
					cost int
				}{path: newPath, i: i, j: j, cost: c})
			}
		}

	case NeighborhoodInsert:
		for i := 1; i < vertexCount-1; i++ {
			for j := 1; j < vertexCount-1; j++ {
				if i == j {
					continue
				}
				newPath := make([]int, vertexCount)
				copy(newPath, currentPath)
				// Przeniesienie wierzchołka z pozycji i do pozycji j
				element := newPath[i]
				copy(newPath[i:], newPath[i+1:])
				newPath[j] = element

				c := t.calculateCost(newPath)
				neighbors = append(neighbors, struct {
					path []int
					i, j int
					cost int
				}{path: newPath, i: i, j: j, cost: c})
			}
		}

	default:
		log.Printf("Nieznana metoda sąsiedztwa: %s. Domyślnie użyto 'swap'.\n", t.neighborhoodMethod)
		// Fallback do swap
		for i := 1; i < vertexCount-1; i++ {
			for j := i + 1; j < vertexCount-1; j++ {
				newPath := make([]int, vertexCount)
				copy(newPath, currentPath)
				newPath[i], newPath[j] = newPath[j], newPath[i]

				c := t.calculateCost(newPath)
				neighbors = append(neighbors, struct {
					path []int
					i, j int
					cost int
				}{path: newPath, i: i, j: j, cost: c})
			}
		}
	}

	return neighbors
}

// Solve implementuje algorytm Tabu Search dla problemu ATSP
func (t *TsATSPSolver) Solve() ([]int, int) {
	vertexCount := t.graph.GetVertexCount()
	if vertexCount == 0 {
		return nil, -1
	}

	// Start czasu
	t.startTime = time.Now()

	// Pobieramy początkowe rozwiązanie zachłanne
	currentSolution := t.graph.GetHamiltonianPathGreedy(t.startVertex)
	currentCost := t.calculateCost(currentSolution)

	// Ustawiamy najlepsze rozwiązanie
	bestSolution := make([]int, len(currentSolution))
	copy(bestSolution, currentSolution)
	bestCost := currentCost

	// Tabu lista: przechowujemy czas tabu dla par indeksów (i,j)
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
				log.Println("Zatrzymano przy iteracji:", iteration, "z powodu przekroczenia limitu czasu.")
				break
			}
		}

		// Generujemy wszystkich sąsiadów
		neighbors := t.getAllNeighbors(currentSolution)
		if len(neighbors) == 0 {
			// Brak sąsiadów, kończymy
			break
		}

		// Wybór najlepszego sąsiada nie będącego tabu lub łamiącego tabu w przypadku poprawy globalnego optimum
		var chosenNeighbor struct {
			path []int
			i, j int
			cost int
		}
		chosenNeighbor.cost = math.MaxInt32

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

		// Jeśli nie znaleźliśmy lepszego sąsiada, wybieramy najlepszego tabu
		if chosenNeighbor.cost == math.MaxInt32 {
			// Wybieramy po prostu najlepszego sąsiada bez względu na tabu
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
		// Ruch (prevI, prevJ) jest tabu przez tabuTenure iteracji
		tabuList[prevI][prevJ] = t.tabuTenure
		tabuList[prevJ][prevI] = t.tabuTenure

		// Zmniejszamy karencję tabu dla wszystkich ruchów
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
