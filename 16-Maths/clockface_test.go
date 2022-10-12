package clockface_test

import (
	"math"
	"testing"
	"time"

	. "go-tdd-tutorial/16-Maths"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := SecondsInRadians(c.time)
			if !roughlyEqualsFloat64(got, c.angle) {
				t.Fatalf("wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := MinutesInRadians(c.time)
			if !roughlyEqualsFloat64(got, c.angle) {
				t.Fatalf("wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := HoursInRadians(c.time)
			if !roughlyEqualsFloat64(got, c.angle) {
				t.Fatalf("wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestSecondsHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := SecondsHandPoint(c.time)
			if !roughlyEqualsPoint(got, c.point) {
				t.Fatalf("Wanted %v point, but got %v", c.point, got)
			}
		})
	}
}

func TestMinutesHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := MinutesHandPoint(c.time)
			if !roughlyEqualsPoint(got, c.point) {
				t.Fatalf("Wanted %v point, but got %v", c.point, got)
			}
		})
	}
}

func TestHoursHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := HoursHandPoint(c.time)
			if !roughlyEqualsPoint(got, c.point) {
				t.Fatalf("Wanted %v point, but got %v", c.point, got)
			}
		})
	}
}

func roughlyEqualsFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualsPoint(a, b Point) bool {
	return roughlyEqualsFloat64(a.X, b.X) && roughlyEqualsFloat64(a.Y, b.Y)
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
