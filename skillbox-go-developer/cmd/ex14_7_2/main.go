package main

import (
	"math/rand"
	"time"
)

type vec2 struct {
	x, y float64
}

const (
	POINTS_COUNT = 2
)

func genRndPoints(count uint) []vec2 {
	var (
		rt []vec2
	)

	rand.Seed(time.Now().UnixMicro())

	rt = make([]vec2, count)

	for i := 0; i < int(count); i++ {
		rt[i].x = rand.Float64()
		rt[i].y = rand.Float64()
	}

	return rt
}

func mapPoint(a vec2) vec2 {
	return vec2{x: 2*a.x + 10, y: -3*a.y - 5}
}

func main() {
	points := genRndPoints(POINTS_COUNT)
	for i := 0; i < POINTS_COUNT; i++ {
		points[i] = mapPoint(points[i])
	}
}
