package checksum

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// EvalChecksum calculates the checkum ps the passed spreadsheet
func EvalChecksum(spreadsheet string) int {
	check := 0
	lines := strings.Split(spreadsheet, "\n")
	for _, line := range lines {
		min, max := getMinMax(line)
		check += int(max - min)
	}

	return check
}

func getMinMax(line string) (min, max float64) {
	max = 0.0 - math.MaxFloat64
	min = math.MaxFloat64

	nums := strings.Split(line, "	")
	for _, num := range nums {
		fnum, _ := strconv.ParseFloat(num, 64)
		min = math.Min(min, fnum)
		max = math.Max(max, fnum)
	}

	return
}

// EvalChecksum2 calculates the checkum ps the passed spreadsheet
func EvalChecksum2(spreadsheet string) int {
	check := 0
	lines := strings.Split(spreadsheet, "\n")
	for _, line := range lines {
		min, max := getDivisibles(line)
		check += int(max / min)
	}

	return check
}

func getDivisibles(line string) (min, max int) {
	fmt.Printf("LINE: %s\n", line)

	nums := strings.Split(line, "	")

	for i := 0; i < (len(nums) - 1); i++ {
		num1st, _ := strconv.Atoi(nums[i])
		fmt.Printf("checking: %d against", num1st)
		for j := i + 1; j < len(nums); j++ {
			num2nd, _ := strconv.Atoi(nums[j])
			fmt.Printf(" %d,", num2nd)
			if num1st%num2nd == 0 {
				max = num1st
				min = num2nd
				fmt.Printf(" FOUND\n")
				return
			} else if num2nd%num1st == 0 {
				max = num2nd
				min = num1st
				fmt.Printf(" FOUND\n")
				return
			}
		}
	}

	fmt.Printf(" NOTHING\n")
	return
}
