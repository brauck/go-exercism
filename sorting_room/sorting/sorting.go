package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// ///////////////////
type NumberBoxContaining struct {
	Value int
}

func (nb NumberBoxContaining) Number() int {
	return nb.Value
}

// ////////////////////
// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	val := float64(nb.Number())
	return fmt.Sprintf("This is a box containing the number %.1f", val)
}

type FancyNumber struct {
	//n string
	N string
}

func (i FancyNumber) Value() string {
	//return i.n
	return i.N
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	_, ok := fnb.(FancyNumber)
	if ok {
		//fmt.Println(str, fnb)
		var i, err = strconv.Atoi(fnb.Value())
		if err == nil {
			return i
		}
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	result := 0.0
	_, ok := fnb.(FancyNumber)
	if ok {
		//fmt.Println(str, fnb)
		var i, err = strconv.ParseFloat(fnb.Value(), 32)
		if err == nil {
			result = i
		}
	}
	return fmt.Sprintf("This is a fancy box containing the number %.1f", result)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	default:
		return "Return to sender"
	}
}
