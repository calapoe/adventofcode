package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayTwo1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	numSafe := 0
	for _, line := range lines {
		parts := strings.Fields(line)

		var levels []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			levels = append(levels, num)
		}
		safe := true
		isIncreasing := true
		if levels[1]-levels[0] < 0 {
			isIncreasing = false
		}
		for i := 0; i < len(levels)-1; i++ {
			diff := levels[i+1] - levels[i]
			if diff == 0 {
				safe = false
				break
			}
			if diff < -3 || diff > 3 {
				safe = false
				break
			}
			if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
				safe = false
				break
			}
		}
		if safe {
			numSafe++
		}
	}
	fmt.Println("Output Day 2 Part 1: ", numSafe)
}

func dayTwo2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	numSafe := 0
	for _, line := range lines {
		parts := strings.Fields(line)

		var levels []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			levels = append(levels, num)
		}
		safe := false
		safe = isSafe(levels)

		if !safe {
			for i := 0; i < len(levels); i++ {
				modifiedLevels := []int{}
				modifiedLevels = append(modifiedLevels, levels[:i]...)
				modifiedLevels = append(modifiedLevels, levels[i+1:]...)
				safe = isSafe(modifiedLevels)
				if safe {
					break
				}
			}
		}
		if safe {
			numSafe++
		}
	}
	fmt.Println("Output Day 2 Part 2: ", numSafe)
}

func isSafe(levels []int) bool {
	safe := true
	isIncreasing := true

	if levels[1]-levels[0] < 0 {
		isIncreasing = false
	}

	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		if diff == 0 {
			safe = false
			break
		}
		if diff < -3 || diff > 3 {
			safe = false
			break
		}
		if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			safe = false
			break
		}
	}
	if safe {
		return true
	}
	return false
}
