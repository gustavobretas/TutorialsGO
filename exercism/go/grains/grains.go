package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number <= 0 {
		return 0, errors.New("negative square")
	} else if number > 64 {
		return 0, errors.New("greater than 64")
	}

	return uint64(math.Pow(float64(2), float64(number-1))), nil
}

func Total() uint64 {

	return uint64(math.Pow(float64(2), float64(64)) - 1)
}
