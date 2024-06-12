package utils

import "strconv"

func ConvertStringToInt(id string) (int, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return -1, err
	} 

	return intId, err
}