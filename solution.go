package square

import "math"

type sidesSpec int

const (
	SidesCircle sidesSpec = iota
	SidesTriangle sidesSpec = iota + 2
	SidesSquare
)

func CalcSquare(sideLen float64, sidesNum sidesSpec) float64 {
	switch sidesNum {
		case SidesSquare:
			return sideLen * sideLen
		case SidesCircle:
			return sideLen * sideLen * math.Pi
		case SidesTriangle:
			return sideLen * sideLen * math.Sqrt(3) / 4
	}
	return 0
}
