package helper

import (
	"fmt"
	"reflect"
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

func TestParseStringDate(t *testing.T) {
	locCur := time.Local
	type args struct {
		location *time.Location
		dateTime string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{"year", args{locCur, "2024"},
			DateCreate(locCur, 2024), false},
		{"year-month", args{locCur, "2024-02"},
			DateCreate(locCur, 2024, 2), false},
		{"year-month-day", args{locCur, "2024-02-03"},
			DateCreate(locCur, 2024, 2, 3), false},
		{"year-month-day-hour", args{locCur, "2024-02-03T11"},
			DateCreate(locCur, 2024, 2, 3, 11), false},
		{"year-month-day-hour-minute", args{locCur, "2024-02-03T11:49"},
			DateCreate(locCur, 2024, 2, 3, 11, 49), false},
		{"year-month-day-hour-minute-sec", args{locCur, "2024-02-03T11:49:39"},
			DateCreate(locCur, 2024, 2, 3, 11, 49, 39), false},
		{"err1", args{locCur, "kek-2024-02-03"},
			DateCreate(locCur), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseStringDate(tt.args.dateTime, time.Local)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStringDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStringDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateGetUTC(t *testing.T) {
	loc, _ := time.LoadLocation("")
	// loc := time.Local
	type args struct {
		location *time.Location
		specs    []int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"year", args{loc, []int{2024}},
			time.Date(2024, 1, 1, 0, 0, 0, 0, loc)},
		{"year-month", args{loc, []int{2024, 0}},
			time.Date(2024, 0, 1, 0, 0, 0, 0, loc)},
		{"year-month-day", args{loc, []int{2024, 1, 1}},
			time.Date(2024, 1, 1, 0, 0, 0, 0, loc)},
		{"year-month-day-hour", args{loc, []int{2024, 1, 1, 10}},
			time.Date(2024, 1, 1, 10, 0, 0, 0, loc)},
		{"year-month-day-hour-minute", args{loc, []int{2024, 1, 1, 10, 49}},
			time.Date(2024, 1, 1, 10, 49, 0, 0, loc)},
		{"year-month-day-hour-minute-sec", args{loc, []int{2024, 1, 1, 10, 49, 33}},
			time.Date(2024, 1, 1, 10, 49, 33, 0, loc)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DateCreate(
				tt.args.location, tt.args.specs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateGetUTC() = %v, want %v", got, tt.want)
			}
		})
	}
}
