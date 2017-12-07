package main

import (
	"fmt"

	"github.com/martinmurphy/advent2017/day03/manhattan"
)

func main() {
	puzzleLocation := 368078
	fmt.Printf("Distance from location \"%d\" is %d\n", puzzleLocation, manhattan.Distance(puzzleLocation))
	targetValue, _, _, _ := manhattan.GoUntilGreaterValue(puzzleLocation)
	fmt.Printf("targetValue is %d\n", targetValue)
}
