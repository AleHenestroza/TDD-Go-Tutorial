package clockface

import (
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

const secondsHandLength = 90
const clockCentreX = 150
const clockCentreY = 150

// SecondsHand is the unit vector of the seconds hand of an analogue clock at time 't', represented as a Point
func SecondsHand(t time.Time) Point {
	p := secondsHandPoint(t)                                    // convert
	p = Point{p.X * secondsHandLength, p.Y * secondsHandLength} // scale
	p = Point{p.X, -p.Y}                                        // flip
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}           // translate
	return p
}

func secondsHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / float64(t.Second())))
}
