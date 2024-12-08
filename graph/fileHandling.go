package graph

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// LoadGraphFromFile wczytuje graf z pliku tekstowego o podanej ścieżce.
// Plik musi mieć na pierwszej linii liczbę wierzchołków, a na kolejnych liniach macierz sąsiedztwa.
// `useTabsAsSeparator` - jeśli true, używa tabów jako separatorów wartości, domyślnie spacje.
func LoadGraphFromFile(filePath string, graph *AdjMatrixGraph, useTabsAsSeparator ...bool) error {
	// Domyślnie ustawienia: spacje jako separator
	separator := " "

	// Jeśli podano parametr `useTabsAsSeparator` i jest ustawiony na true, zmieniamy separator
	if len(useTabsAsSeparator) > 0 && useTabsAsSeparator[0] {
		separator = "\t"
	}

	// Otwórz plik
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Buforowane czytanie pliku
	scanner := bufio.NewScanner(file)

	// Odczytaj liczbę wierzchołków (pierwsza linia)
	if !scanner.Scan() {
		return errors.New("plik jest pusty lub uszkodzony")
	}
	vertexCount, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return errors.New("błąd podczas odczytu liczby wierzchołków")
	}

	// Inicjalizacja grafu
	graph.vertexCount = vertexCount
	graph.adjMatrix = make([][]int, vertexCount)
	for i := 0; i < vertexCount; i++ {
		graph.adjMatrix[i] = make([]int, vertexCount)
	}

	// Odczytuj kolejne wiersze macierzy sąsiedztwa
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Dziel linię według separatora (spacja lub tab)
		values := strings.Split(line, separator)

		if len(values) != vertexCount {
			return errors.New("niewłaściwa liczba elementów w wierszu macierzy sąsiedztwa")
		}

		// Przetwórz wartości w wierszu
		for col, value := range values {
			weight, err := strconv.Atoi(value)
			if err != nil {
				return errors.New("błąd podczas konwersji wartości macierzy do liczby całkowitej")
			}
			graph.adjMatrix[row][col] = weight
		}
		row++
	}

	if row != vertexCount {
		return errors.New("niewłaściwa liczba wierszy w macierzy sąsiedztwa")
	}

	return nil
}

// SaveGraphToFile zapisuje graf do pliku tekstowego o podanej ścieżce.
// Pierwsza linia zawiera liczbę wierzchołków, a kolejne linie zawierają macierz sąsiedztwa.
// `useTabsAsSeparator` - jeśli true, używa tabów jako separatorów wartości, domyślnie spacje.
func SaveGraphToFile(g Graph, filePath string, useTabsAsSeparator ...bool) error {
	// Domyślnie ustawienia: spacje jako separator
	separator := " "

	// Jeśli podano parametr `useTabsAsSeparator` i jest ustawiony na true, zmieniamy separator
	if len(useTabsAsSeparator) > 0 && useTabsAsSeparator[0] {
		separator = "\t"
	}

	// Utwórz lub otwórz plik do zapisu
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Strings.Builder do buforowania danych przed zapisem do pliku
	var out strings.Builder

	// Zapisz liczbę wierzchołków
	out.WriteString(strconv.Itoa(g.GetVertexCount()) + "\n")

	// Zapisz macierz sąsiedztwa
	for i := 0; i < g.GetVertexCount(); i++ {
		for j := 0; j < g.GetVertexCount(); j++ {
			// Pobierz wagę krawędzi między wierzchołkami i, j
			edge := g.GetEdge(i, j)
			out.WriteString(strconv.Itoa(edge.Weight))
			// Dodaj separator między wartościami, oprócz ostatniej w wierszu
			if j < g.GetVertexCount()-1 {
				out.WriteString(separator)
			}
		}
		out.WriteString("\n")
	}

	// Zapisz zawartość do pliku
	_, err = file.WriteString(out.String())
	if err != nil {
		return err
	}

	return nil
}
