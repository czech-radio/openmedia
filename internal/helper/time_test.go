package helper

import (
	"fmt"
	"testing"
	"time"
)

func Test_IsOlderThanOneISOweek(t *testing.T) {
	type tpair struct{ Input, ExpectedOutput any }
	timeNow := time.Now()
	weekDay := int(time.Now().Weekday())
	addWeek := 7 - weekDay
	testPairs := []tpair{

		// Input date is same ISOweek
		{timeNow.AddDate(0, 0, 0), false},

		// Input date older ISOweek
		{timeNow.AddDate(0, 0, -7), true},
		{timeNow.AddDate(0, 0, -19), true},

		// Input date is newer
		{timeNow.AddDate(0, 0, addWeek), false},
		{timeNow.AddDate(0, 0, 7), false},
		{timeNow.AddDate(0, 0, 10), false},
	}
	for i := range testPairs {
		ok := IsOlderThanOneISOweek(testPairs[i].Input.(time.Time), timeNow)
		if ok != testPairs[i].ExpectedOutput {
			t.Errorf("pair %d failed for inputs: %v, %v", i, testPairs[i].Input, timeNow)
		}
	}
}

func TestDateRangesIntersection(t *testing.T) {
	timeZone, _ := time.LoadLocation("")
	testCases := []struct {
		name      string
		r1        [2]time.Time
		r2        [2]time.Time
		intersect bool
	}{
		{
			name: "Whole intersection",
			r1: [2]time.Time{
				time.Date(2024, 3, 10, 8, 0, 0, 0, timeZone),
				time.Date(2024, 3, 10, 12, 0, 0, 0, timeZone)},
			r2: [2]time.Time{
				time.Date(2024, 3, 10, 9, 0, 0, 0, timeZone),
				time.Date(2024, 3, 10, 10, 0, 0, 0, timeZone)},
			intersect: true,
		},
		{
			name: "Partial Intersection right",
			r1: [2]time.Time{
				time.Date(2024, 3, 10, 8, 0, 0, 0, timeZone),
				time.Date(2024, 3, 10, 12, 0, 0, 0, timeZone)},
			r2: [2]time.Time{
				time.Date(2024, 3, 10, 10, 0, 0, 0, timeZone),
				time.Date(2024, 3, 10, 14, 0, 0, 0, timeZone)},
			intersect: true,
		},
		{
			name: "Partial Intersection left",
			r1: [2]time.Time{
				time.Date(2024, 3, 10, 10, 0, 0, 0, timeZone),
				time.Date(2024, 3, 10, 14, 0, 0, 0, timeZone)},
			r2: [2]time.Time{
				time.Date(2024, 3, 10, 8, 0, 0, 0, timeZone),
				time.Date(2024, 3, 10, 12, 0, 0, 0, timeZone)},
			intersect: true,
		},
		{
			name: "No Intersection before",
			r1: [2]time.Time{
				time.Date(2024, 4, 10, 0, 0, 0, 0, timeZone),
				time.Date(2024, 4, 11, 0, 0, 0, 0, timeZone)},
			r2: [2]time.Time{
				time.Date(2024, 3, 10, 0, 0, 0, 0, timeZone),
				time.Date(2024, 3, 10, 0, 0, 0, 0, timeZone)},
			intersect: false,
		},
		{
			name: "No Intersection After",
			r1: [2]time.Time{
				time.Date(2024, 2, 10, 0, 0, 0, 0, timeZone),
				time.Date(2024, 2, 11, 0, 0, 0, 0, timeZone)},
			r2: [2]time.Time{
				time.Date(2024, 3, 10, 0, 0, 0, 0, timeZone),
				time.Date(2024, 3, 11, 0, 0, 0, 0, timeZone)},
			intersect: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dateRange, ok := DateRangesIntersection(tc.r1, tc.r2)
			if ok != tc.intersect {
				t.Errorf("expected intersect to be %t; got %v", tc.intersect, dateRange)
			}
		})
	}
}

func TestCzechDateToUTC(t *testing.T) {
	from, err := CzechDateToUTC(2024, 3, 4, 0)
	if err != nil {
		t.Error(err)
	}
	to, err := CzechDateToUTC(2024, 3, 5, 0)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(from)
	fmt.Println(to)
}
