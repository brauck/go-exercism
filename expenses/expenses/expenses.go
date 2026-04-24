package expenses

import (
	"errors"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	result := []Record{}
	for _, v := range in {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	from := p.From
	to := p.To

	return func(r Record) bool {
		if r.Day >= from && r.Day <= to {
			return true
		}
		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {

	return func(r Record) bool {
		if r.Category == c {
			return true
		}
		return false
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	days := Filter(in, ByDaysPeriod((p)))
	result := 0.0

	for _, d := range days {
		result += d.Amount
	}
	return result
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	expenses := 0.0
	isPresent := false
	for _, r := range in {
		if r.Category == c {
			isPresent = true
			break
		}
	}

	if !isPresent {
		return 0, errors.New("unknown category entertainment")
	}
	found := Filter(in, ByCategory(c))
	expenses = TotalByPeriod(found, p)

	return expenses, nil
}

func Day1Records(r Record) bool {
	return r.Day == 1
}
