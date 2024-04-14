package helper

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	pattern := "^13:00-14:00"
	regex := regexp.MustCompile(pattern)
	testString := "13:00-14:00"
	ok := regex.MatchString(testString)
	fmt.Println("matches", ok, pattern, testString)

	testString = "kek 13:00-14:00"
	ok = regex.MatchString(testString)
	fmt.Println("matches", ok, pattern, testString)
}
