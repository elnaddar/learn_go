package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const secondHandelLength = 90
const clockCenterX = 150
const clockCenterY = 150

func SecondHand(tm time.Time) Point {
	p := secondHandPoint(tm)
	p = Point{p.X * secondHandelLength, p.Y * -secondHandelLength}
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}

	return p
}

func secondsInRadians(tm time.Time) float64 {
	return math.Pi / (30 / float64(tm.Second()))
}

func secondHandPoint(tm time.Time) Point {
	radAngel := secondsInRadians(tm)
	return Point{math.Sin(radAngel), math.Cos(radAngel)}
}
	