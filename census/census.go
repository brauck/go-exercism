package main

import (
	s "exercism/census/census"
	"fmt"
)

var p = fmt.Println

func main() {
	/* name := "Matthew Sanabria"
	age := 29
	address := map[string]string{"street": "Main St."}
	p(s.NewResident(name, age, address))

	name = "Matthew Sanabria"
	age = 0
	address = make(map[string]string)

	resident := s.NewResident(name, age, address)

	p(resident.HasRequiredInfo()) */

	name1 := "Matthew Sanabria"
	age1 := 29
	address1 := map[string]string{"street": "Main St."}

	resident1 := s.NewResident(name1, age1, address1)

	name2 := "Rob Pike"
	age2 := 0
	address2 := make(map[string]string)

	resident2 := s.NewResident(name2, age2, address2)

	residents := []*s.Resident{resident1, resident2}

	p(s.Count(residents))
}
