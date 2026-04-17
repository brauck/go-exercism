package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// "July 25, 2019 13:45:00"
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	//"Thursday, July 25, 2019 13:45:00"
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	hour := t.Hour()
	if hour >= 12 && hour < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	//"Thursday, July 25, 2019 13:45:00"
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	weekDay := t.Weekday()
	month := t.Month()
	year := t.Year()
	day := t.Day()
	hour := t.Hour()
	minuets := t.Minute()

	desctiprion := fmt.Sprintf("You have an appointment on %v, %v %d, %d, at %02d:%02d.",
		weekDay, month, day, year, hour, minuets)

	return desctiprion
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	anniversary := time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)

	return anniversary
}
