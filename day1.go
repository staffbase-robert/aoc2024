package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func mustLen[T any](l []T, want int) {
	if len(l) != want {
		panic(fmt.Sprintf("unexpected length of list, want %d, got %d\nitems:\n%v", want, len(l), l))
	}
}

func mustInt(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return i
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func day1() {
	file, err := os.Open("./input_d1")
	handleError(err)

	bytes, err := io.ReadAll(file)
	handleError(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")

	l1 := make([]int, 0)
	l2 := make([]int, 0)
	for _, line := range lines {
		items := strings.Fields(line)
		mustLen(items, 2)
		l1 = append(l1, mustInt(items[0]))
		l2 = append(l2, mustInt(items[1]))
	}

	sort.Ints(l1)
	sort.Ints(l2)
	totalDist := 0
	for i, left := range l1 {
		right := l2[i]
		totalDist += abs(right - left)
	}
	fmt.Println("result", totalDist)
}

func day1b() {
	file, err := os.Open("./input_d1")
	handleError(err)

	bytes, err := io.ReadAll(file)
	handleError(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")

	l1 := make([]int, 0)
	l2 := make([]int, 0)
	for _, line := range lines {
		items := strings.Fields(line)
		mustLen(items, 2)
		l1 = append(l1, mustInt(items[0]))
		l2 = append(l2, mustInt(items[1]))
	}

	freq := make(map[int]int)
	for _, v := range l2 {
		before, ok := freq[v]
		if !ok {
			freq[v] = 0
		}

		freq[v] = before + 1
	}

	score := 0
	for _, v := range l1 {
		f := freq[v]
		score += v * f
	}
	fmt.Println(score)
}
