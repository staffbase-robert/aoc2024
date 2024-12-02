package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/staffbase-robert/aoc2024/utils"
)

func solve() {
	file, err := os.Open("./input_d0")
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")
	fmt.Println(lines[0])
}

func solveB() {
	file, err := os.Open("./input_d0")
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")
	fmt.Println(lines[0])
}

var isPartTwo = flag.Bool("b", false, "select if part two")

func main() {
	flag.Parse()
	if *isPartTwo {
		solveB()
	} else {
		solve()
	}
}
