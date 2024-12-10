package amountTests

import (
	"projekt2/tests/amountTests/saAmountTests"
	"projekt2/tests/amountTests/tsAmountTests"
)

func RunAmountTests() {
	tsAmountTests.RunTSAmountTests(100, 50)
	saAmountTests.RunSAAmountTests(100, 50)
}
