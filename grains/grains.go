package grains

import "math"
import "errors"

func Total() uint64 {
	const numberOfSquares = 64
	var sum uint64
	var grains uint64
	for square := 1; square <= numberOfSquares; square++ {
		grains, _ = Square(square)
		sum += grains
	}
	return sum
}

func Square(square int) (uint64, error) {
	if (square < 1) || (square > 64) {
		return 0, errors.New("Invalid square.")
	}
	return uint64(math.Pow(2, float64(square-1))), nil
}
