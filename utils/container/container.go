package stringcontainer2d

import (
	"errors"
	"fmt"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Container struct {
	Lines []string
}

func NewPadded(lines []string) Container {
	tb := strings.Repeat("#", len(lines[0])+2)

	ret := []string{tb}
	for _, line := range lines {
		ret = append(ret, fmt.Sprintf("#%s#", line))
	}

	ret = append(ret, tb)
	return Container{ret}
}

func (c Container) Print() {
	for _, line := range c.Lines {
		fmt.Println(line)
	}
}

var ErrOutOfBounds = errors.New("out of bounds")

func (c Container) At(p Point) (string, error) {
	x := p.X
	y := p.Y
	if x < 0 {
		return "", ErrOutOfBounds
	}
	if y < 0 {
		return "", ErrOutOfBounds
	}
	if x >= len(c.Lines[0]) {
		return "", ErrOutOfBounds
	}
	if y >= len(c.Lines) {
		return "", ErrOutOfBounds
	}
	return c.Lines[y][x : x+1], nil
}
