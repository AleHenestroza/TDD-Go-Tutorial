package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = secondsInHalfClock * 2
	minutesInHalfClock = 30
	minutesInClock     = minutesInHalfClock * 2
	hoursInHalfClock   = 6
	hoursInClock       = hoursInHalfClock * 2
)

// A Point represents a two dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

func SecondsHandPoint(t time.Time) Point {
	return angleToPoint(SecondsInRadians(t))
}

func MinutesHandPoint(t time.Time) Point {
	return angleToPoint(MinutesInRadians(t))
}

func HoursHandPoint(t time.Time) Point {
	return angleToPoint(HoursInRadians(t))
}

func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / float64(t.Second())))
}

func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / minutesInClock) +
		(math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func HoursInRadians(t time.Time) float64 {
	return (MinutesInRadians(t) / hoursInClock) +
		(math.Pi / (hoursInHalfClock / float64(t.Hour()%12)))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
