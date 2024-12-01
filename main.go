package main

import (
	"flag"
	"fmt"
)

var day = flag.Int("d", 1, "select day")
var isPartTwo = flag.Bool("b", false, "select if part two")

var solutions = map[string]func(){
	"day1":  day1,
	"day1b": day1b,
}

func main() {
	flag.Parse()
	b := ""
	if *isPartTwo {
		b = "b"
	}
	key := fmt.Sprintf("day%d%s", *day, b)
	if _, ok := solutions[key]; !ok {
		fmt.Printf("solution for %s doesn't exist\n", key)
		return
	}
	fmt.Println("running", key)
	solutions[key]()
}
