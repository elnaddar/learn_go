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
	clockCenterX       = 150
	clockCenterY       = 150
)

func pointToHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * -length}
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}
	return p
}

func secondsInRadians(tm time.Time) float64 {
	return math.Pi / (30 / float64(tm.Second()))
}

func minutesInRadians(tm time.Time) float64 {
	return (secondsInRadians(tm) / 60) + (math.Pi / (30 / float64(tm.Minute())))
}

func secondHandPoint(tm time.Time) Point {
	return angelToPoint(secondsInRadians(tm))
}

func minuteHandPoint(tm time.Time) Point {
	return angelToPoint(minutesInRadians(tm))
}

func angelToPoint(angel float64) Point {
	return Point{math.Sin(angel), math.Cos(angel)}
}
