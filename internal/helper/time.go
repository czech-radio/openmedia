package helper

import (
	"time"

	"github.com/ncruces/go-strftime"
)

// Sleeper sleeps for specified durration
func Sleeper(duration int, time_unit string) {
	switch time_unit {
	case "ms":
		time.Sleep(time.Duration(duration) * time.Millisecond)
	case "s":
		time.Sleep(time.Duration(duration) * time.Second)
	case "m":
		time.Sleep(time.Duration(duration) * time.Minute)
	default:
		panic("Wrong time time_unit")
	}
}

// IsOlderThanOneISOweek
func IsOlderThanOneISOweek(dateToCheck, dateNow time.Time) bool {
	year_check, week_check := dateToCheck.ISOWeek()
	year_now, week_now := dateNow.ISOWeek()
	if year_check < year_now {
		return true
	}
	return week_check < week_now
}

// TimeCurrent
func TimeCurrent() string {
	tm := time.Now()
	return strftime.Format("%FT%T", tm)
}

// DateRangesIntersection
func DateRangesIntersection(rA, rB [2]time.Time) ([2]time.Time, bool) {
	resrange := [2]time.Time{}

	// Special cases
	// rA is default zero time!
	if rA[0].IsZero() && rA[1].IsZero() {
		return rB, true
	}

	if rA[0].After(rB[1]) {
		return resrange, false
	}
	if rA[1].Before(rB[0]) {
		return resrange, false
	}

	// Get intersec start time
	var start time.Time
	if rA[0].Before(rB[0]) {
		start = rB[0]
	} else {
		start = rA[0]
	}

	// Get intersec end time
	var end time.Time
	if rA[1].Before(rB[1]) {
		end = rA[1]
	} else {
		end = rB[1]
	}
	resrange[0] = start
	resrange[1] = end
	return resrange, true
}

// DateInRange
func DateInRange(interval [2]time.Time, dateToCheck time.Time) bool {
	if interval[0].Before(dateToCheck) && interval[1].After(dateToCheck) {
		return true
	}
	if dateToCheck.Equal(interval[0]) {
		return true
	}
	if dateToCheck.Equal(interval[1]) {
		return true
	}
	return false
}

// CzechDateToUTC
func CzechDateToUTC(year, month, day, hour int) (
	time.Time, error) {
	var res time.Time
	location, err := time.LoadLocation("Europe/Prague")
	if err != nil {
		return res, err
	}
	mont := time.Month(month)
	czechDate := time.Date(
		year, mont, day, hour,
		0, 0, 0, location,
	)
	return czechDate.UTC(), nil
}
