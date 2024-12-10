package amountTests

import (
	"projekt2/tests/amountTests/saAmountTests"
	"projekt2/tests/amountTests/tsAmountTests"
)

func RunAmountTests() {
	saAmountTests.RunSAAmountTests(100, 50)
	tsAmountTests.RunTSAmountTests(100, 50)
}
