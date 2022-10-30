package main

import (
	"math/rand"
	"errors"
)

func roll (i int) (int, error) {
	if (i <= 0) {
		return 0, errors.New("maximum value must be higher than 0")
	}
	return rand.Intn(i) + 1, nil
}

// returns a roll between i and min.
func minRoll (i, min int) (int, error) {
	if i <= 0 || min <= 0 {
		return 0, errors.New("values given to minRoll function must be positive.")
	}
	if (min >= i) {
		return 0, errors.New("minimum roll value must be less than maximum roll value.")
	}
	
	return rand.Intn(i-min) + min + 1, nil
}

func addRoll (i, add int) (int, error) {
	if i <= 0 || add <= 0 {
		return 0, errors.New("values given to addroll function must be positive.")
	}
	r, _:= roll(i)
	return r + add, nil
}