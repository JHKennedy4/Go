package main

import "time"

type Dates struct {
	Start string
	End   string
}

// takes a start date and end date of format YYYY-MM-DD
// and picks an array of Dates that maximize the utilization of the vacation
// days in combination with weekends for the purpose of planning trips
func PickDates(startDate string, endDate string, vacation int) []Dates {
	format := "2006-1-2 MST"
	start, err := time.Parse(format, startDate+" EST")
	if err != nil {
		panic(err)
	}
	end, err := time.Parse(format, endDate+" EST")
	if err != nil {
		panic(err)
	}

	vac := time.Duration(vacation*24) * time.Hour
	for start.Before(end) {
		depth := start.Add(vac).Weekday()
		// if depth == saturday or sunday
		// calculate remaining

		// start interval is always start
		// get nearest weekend
		// get slice of time prior to weekend
		// if duration is less than time prior to weekend
		// subtract duration prior to weekend
		// get remaining vacation
		// get nearest weekend
		// get slice of time prior to weekend
		// if duration is less than time prior to weekend
		// subtract duration prior to weekend
		// else
		// end := weekend + remaining vacation
		// need a recursive func with a switch
		if depth == 0 || depth == 6 {

		}

		// add a day
		start = start.Add(time.Duration(24) * time.Hour)
	}
	return []Dates{}
}
