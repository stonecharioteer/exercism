package booking

import (
    "fmt"
    "time"
)
// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout:= "1/2/2006 15:04:05"
    t,_ := time.Parse(layout, date)
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	timestamp, e := time.Parse("January _2, 2006 15:04:05",date)
    
    if e != nil {
        msg := fmt.Sprintln("date: ", date, " timestamp: ", timestamp, " DATEERROR")
        panic(msg)
    }
    
    return timestamp.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    t, e := time.Parse("Monday, January _2, 2006 15:04:05", date)
    if e!=nil {
        msg := fmt.Sprintln("date: ", date, " timestamp: ", t, " DATEERROR")
        panic(msg)
    }
    return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    layout:="1/2/2006 15:04:05"
	t,e := time.Parse(layout, date)
    if e!=nil {
        msg := fmt.Sprintln("date: ", date, " timestamp: ", t, " DATEERROR")
        panic(msg)
    }
    return t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
