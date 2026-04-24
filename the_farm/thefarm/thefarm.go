package thefarm

import (
	"errors"
	"fmt"
)

type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	result := 0.0
	if cows <= 0 {
		err := errors.New("something went wrong")
		return result, err
	}
	amount, errAmount := fc.FodderAmount(cows)
	factor, errFactor := fc.FatteningFactor()
	if errAmount != nil {
		return result, errAmount
	} else if errFactor != nil {
		return result, errFactor
	}

	/* if amount == 0 {
		return result, errAmount
	} else if factor == 0 {
		return result, errFactor
	} */

	result = amount / float64(cows) * factor
	return result, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		err := errors.New("invalid number of cows")
		return 0.0, err
	}
	return DivideFood(fc, cows)
}

// TODO: define the 'ValidateNumberOfCows' function
type MyCustomError struct {
	cows    int
	message string
}

func (e *MyCustomError) Error() string {
	if e.cows < 0 {
		e.message = "there are no negative cows"
	} else if e.cows == 0 {
		e.message = "no cows don't need food"
	}
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) error {
	err := MyCustomError{cows: cows}
	if cows <= 0 {
		return &err
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
