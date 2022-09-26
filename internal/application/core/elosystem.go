package core

import (
	"math"
)

const k = 32
const scale = 400

func probability(elo1, elo2 float64) float64 {
	return 1.0 / (1.0 + math.Pow(10.0, (elo1-elo2)/scale))
}

func CalculateElo(elo1, elo2 float64) float64 {
	a := probability(elo1, elo2)
	a *= float64(k)
	return math.Abs(math.Round(a))
}
