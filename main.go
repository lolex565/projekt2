package main

import (
	"projekt2/tests/amountTests"
	"projekt2/tests/tuningTests"
)

func main() {
	//mainMenu := menu.NewMenu()
	//mainMenu.RunInteractiveMenu()

	//g := graph.NewAdjMatrixGraph(171, 100000000)
	//graph.LoadGraphFromFile("ftv170.atsp", g, true)
	//fmt.Println(g.ToString())

	tuningTests.RunTuning()
	amountTests.RunAmountTests()

}
