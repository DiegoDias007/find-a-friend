package utils

import (
	"math/rand"
)

func GenerateRandomNumber(generate_range int) int {
	randomNumber := rand.Intn(generate_range)
	return randomNumber
}