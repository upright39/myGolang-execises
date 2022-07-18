package execises9

import (
	"fmt"
	"time"
)

func Schedule(AppointmentDate string) {

	value, err := time.Parse("01/02/2006 15:04:05", AppointmentDate)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("new Date: %v\n", value)
	}

}

func HasPassed(AppiontmentDate string) {

	values, err := time.Parse("January _2, 2006 15:04:05", AppiontmentDate)
	if err != nil {
		fmt.Println(err)
	}
	now := time.Now()
	// T := now.Add(3 * time.Hour)

	// fmt.Println("Before", now.Before(values))
	fmt.Println("After:", now.After(values))

	// fmt.Println(T)

}

func IsAfternoonAppointment(myAppointmen string) {
	values, err := time.Parse("Monday, January _2, 2006 15:04:05", myAppointmen)
	timeCheck := values.Hour()

	if err != nil {
		fmt.Println(err)
	} else if timeCheck >= 12 && timeCheck < 18 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func Description(myAppointment string) {
	val, err := time.Parse("01/02/2006 15:04:05", myAppointment)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf(" You have an appointment on %v,%v %v at %v.\n", val.Format("Monday"), val.Format("January _2,"), val.Format("2006"), val.Format("15:04"))
	}
}

func AniversaryDate() time.Time {
	aniverayDate := time.Date(2020, time.September, 15, 0, 0, 0, 0, time.UTC)

	return aniverayDate
}
