package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondHandelLength = 90
	minuteHandelLength = 80
	hourHandelLength   = 50
	clockCenterX       = 150
	clockCenterY       = 150
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

func pointToHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * -length}
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}
	return p
}

func secondsInRadians(tm time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(tm.Second()))
}

func minutesInRadians(tm time.Time) float64 {
	return (secondsInRadians(tm) / minutesInClock) + (math.Pi / (minutesInHalfClock / float64(tm.Minute())))
}

func hoursInRadians(tm time.Time) float64 {
	return (minutesInRadians(tm) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(tm.Hour()%12)))
}

func secondHandPoint(tm time.Time) Point {
	return angelToPoint(secondsInRadians(tm))
}

func minuteHandPoint(tm time.Time) Point {
	return angelToPoint(minutesInRadians(tm))
}

func hourHandPoint(tm time.Time) Point {
	return angelToPoint(hoursInRadians(tm))
}

func angelToPoint(angel float64) Point {
	return Point{math.Sin(angel), math.Cos(angel)}
}
