package main

import (
	"iter"
	"testing"

	"github.com/staffbase-robert/aoc2024/utils"
	c "github.com/staffbase-robert/aoc2024/utils/container"
)

func TestAntinodes(t *testing.T) {
	pair := pair{a: c.Point{X: 1, Y: 1}, b: c.Point{X: 2, Y: 2}}

	_next, stop := iter.Pull(pair.shoot(true))
	next := func() c.Point {
		v, _ := _next()
		return v
	}

	utils.MustEq(next(), c.Point{X: 0, Y: 0})
	utils.MustEq(next(), c.Point{X: -1, Y: -1})
	utils.MustEq(next(), c.Point{X: -2, Y: -2})
	utils.MustEq(next(), c.Point{X: -3, Y: -3})
	stop()
}
