package utils

import (
	"errors"
	"log"
	"math"
	"strconv"
)

func String2Int64(num string) int64 {
	number, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		log.Fatal("Couldn't parse number from string")
	}
	return number
}

func String2Float64(num string) (float64, error) {
	floatValue, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return -1, errors.New("failed to parse number from string")
	}
	return floatValue, nil
}

func Floor(num float64) float64 {
	return math.Floor(num*math.Pow(10, 8)) / math.Pow(10, 8)
}
