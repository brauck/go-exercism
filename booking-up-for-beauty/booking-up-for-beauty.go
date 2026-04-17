package main

import (
	b "exercism/booking-up-for-beauty/booking"
	"fmt"
)

var p = fmt.Println

func main() {
	//p(b.Schedule("7/25/2019 13:45:00"))
	//p(b.HasPassed("July 25, 2019 13:45:00"))
	//p(b.IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00"))
	//p(b.Description("7/25/2019 13:45:00"))
	p(b.AnniversaryDate())
}
