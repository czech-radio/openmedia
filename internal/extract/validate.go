package extract

import "strconv"

// ValidateStopaz
func ValidateStopaz(stopaz string) (string, error) {
	milliSeconds, err := strconv.ParseInt(stopaz, 10, 64)
	specVal := RowFieldSpecialValueCodeMap[RowFieldValueNotValid]
	if err != nil {
		return stopaz, err
	}
	if milliSeconds < 0 {
		// negative stopaz is invalid
		return specVal, err
	}
	return stopaz, nil
}
