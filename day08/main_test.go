package main

import (
	"testing"

	"github.com/staffbase-robert/aoc2024/utils"
	c "github.com/staffbase-robert/aoc2024/utils/container"
)

func TestAntinodes(t *testing.T) {
	pair := pair{a: c.Point{X: 1, Y: 1}, b: c.Point{X: 2, Y: 2}}

	_, iter := pair.antinodes()

	utils.MustEq(iter(), c.Point{X: 0, Y: 0})
	utils.MustEq(iter(), c.Point{X: -1, Y: -1})
	utils.MustEq(iter(), c.Point{X: -2, Y: -2})
	utils.MustEq(iter(), c.Point{X: -3, Y: -3})
}
