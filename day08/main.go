package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/staffbase-robert/aoc2024/utils"
	c "github.com/staffbase-robert/aoc2024/utils/container"
	"github.com/staffbase-robert/aoc2024/utils/set"
)

var inputFile = flag.String("input", "example", "select input file")

var con c.Container
var antennas = make(map[string][]c.Point)

func solve() {
	handleInput()
	antinodes := set.New[c.Point]()
	for _, points := range antennas {
		pairs := makePairs(points)
		for _, pair := range pairs {
			lIter, rIter := pair.antinodes()
			a := lIter()
			b := rIter()
			if _, err := con.At(a); err == nil {
				antinodes.Add(a)
			}
			if _, err := con.At(b); err == nil {
				antinodes.Add(b)
			}
		}
	}
	fmt.Println("part1", len(antinodes))

	antinodes = set.New[c.Point]()
	for _, points := range antennas {
		pairs := makePairs(points)
		utils.MustTrue(len(pairs) > 0)
		for _, pa := range pairs {
			antinodes.Add(pa.a, pa.b)
			lIter, rIter := pa.antinodes()
			for {
				p := lIter()
				if _, err := con.At(p); err != nil {
					break
				}
				antinodes.Add(c.Point{X: p.X, Y: p.Y})
			}

			for {
				p := rIter()
				if _, err := con.At(p); err != nil {
					break
				}
				antinodes.Add(c.Point{X: p.X, Y: p.Y})
			}
		}
	}
	fmt.Println("part2", len(antinodes))
}

func main() {
	flag.Parse()
	solve()
}

func handleInput() {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)
	lines := strings.Split(string(bytes), "\n")

	con = c.Container{
		Lines: lines,
	}

	for y, line := range lines {
		for x, r := range line {
			s := string(r)
			if s == "." {
				continue
			}
			if _, exists := antennas[s]; !exists {
				antennas[s] = make([]c.Point, 0)
			}

			sl := antennas[s]
			sl = append(sl, c.Point{X: x, Y: y})
			antennas[s] = sl
		}
	}
}

type pair struct {
	a c.Point
	b c.Point
}

func makePairs(points []c.Point) []pair {
	ret := make([]pair, 0)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			ret = append(ret, pair{points[i], points[j]})
		}
	}
	return ret
}

type iterator func() c.Point

func (p pair) antinodes() (lhs iterator, rhs iterator) {
	vecA2B := p.b.Sub(p.a)
	curB := p.b
	lhs = func() c.Point {
		curB = curB.Add(vecA2B)
		return curB
	}

	vecB2A := p.a.Sub(p.b)
	curA := p.a
	rhs = func() c.Point {
		curA = curA.Add(vecB2A)
		return curA
	}

	return
}
