package utils

import (
	"fmt"
	"strconv"
)

func ConvertStringToInt(convert string) (int, error) {
	intConvert, err := strconv.Atoi(convert)
	if err != nil {
		return -1, fmt.Errorf("could not parse string.")
	} 

	return intConvert, nil
}

func ConvertStringToFloat(convert string) (float64, error) {
	floatConvert, err := strconv.ParseFloat(convert, 64)
	if err != nil {
		return -1, fmt.Errorf("could not parse string.")
	}

	return floatConvert, nil
}