package amountTests

type TestRunner interface {
	SetStartVertexAmount(startVertexAmount int)
	SetVertexAmountInterval(vertexAmountInterval int)
	SetIntervalCount(intervalCount int)
	SetRandomGraphGeneratorValues(noEdgeValue, maxWeight int)
	SaveResultsToCsv()
}
