package main

import (
	i "exercism/interest-is-interesting/interest"
	"fmt"
)

var p = fmt.Println

func main() {
	p(i.InterestRate(200.75))
	p(i.Interest(200.75))

	p(i.AnnualBalanceUpdate(200.75))

	balance := 200.75
	targetBalance := 214.88
	p(i.YearsBeforeDesiredBalance(balance, targetBalance))

	balance = 1000.0
	targetBalance = 1032.682765146664
	p(i.YearsBeforeDesiredBalance(balance, targetBalance))
}
