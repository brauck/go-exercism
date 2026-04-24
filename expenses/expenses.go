package main

import (
	s "exercism/expenses/expenses"
	"fmt"
)

var p = fmt.Println

func main() {
	records1 := []s.Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
	}

	p(s.Filter(records1, s.Day1Records))

	records2 := []s.Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	period := s.DaysPeriod{From: 1, To: 15}

	p(s.Filter(records2, s.ByDaysPeriod(period)))

	p(s.Filter(records2, s.ByCategory("groceries")))

	records3 := []s.Record{
		{Day: 15, Amount: 16, Category: "entertainment"},
		{Day: 32, Amount: 20, Category: "groceries"},
		{Day: 40, Amount: 30, Category: "entertainment"},
	}

	p1 := s.DaysPeriod{From: 1, To: 30}
	p2 := s.DaysPeriod{From: 31, To: 60}

	p(s.TotalByPeriod(records3, p1))
	p(s.TotalByPeriod(records3, p2))

	p(s.CategoryExpenses(records2, p1, "entertainment"))
	// => 0, error(unknown category entertainment)

	p(s.CategoryExpenses(records2, p1, "rent"))
	// => 1300, nil

	p(s.CategoryExpenses(records2, p2, "rent"))
	// => 0, nil
}
