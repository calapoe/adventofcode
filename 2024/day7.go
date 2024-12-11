package main

import (
	"fmt"
	"strconv"
	"strings"
)

func daySeven1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	out := 0
	for _, part := range lines {
		parts := strings.Split(part, ":")
		res, _ := strconv.Atoi(parts[0])
		values := strings.Fields(parts[1])
		numbers := make([]int, len(values))
		for i, val := range values {
			numbers[i], _ = strconv.Atoi(val)
		}

		if evaluate1(res, numbers, 0, 0) {
			out += res
		}
	}
	fmt.Println("Output Day 7 Part 1: ", out)
}

func daySeven2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	out := 0
	for _, part := range lines {
		parts := strings.Split(part, ":")
		res, _ := strconv.Atoi(parts[0])
		values := strings.Fields(parts[1])
		numbers := make([]int, len(values))
		for i, val := range values {
			numbers[i], _ = strconv.Atoi(val)
		}

		if evaluate2(res, numbers, 0, 0) {
			out += res
		}
	}
	fmt.Println("Output Day 7 Part 2: ", out)
}

func evaluate2(target int, numbers []int, index int, currRes int) bool {
	if index == len(numbers) {
		return currRes == target
	}
	currNum := numbers[index]
	newRes := currNum + currRes
	if newRes <= target {
		if evaluate2(target, numbers, index+1, newRes) {
			return true
		}
	}

	newRes = currRes * currNum
	if newRes <= target {
		if evaluate2(target, numbers, index+1, newRes) {
			return true
		}
	}

	newRes, _ = strconv.Atoi(fmt.Sprintf("%d%d", currRes, currNum))
	if newRes <= target {
		return evaluate2(target, numbers, index+1, newRes)
	}
	return false
}

func evaluate1(target int, numbers []int, index int, currRes int) bool {
	if index == len(numbers) {
		return currRes == target
	}
	currNum := numbers[index]
	newRes := currNum + currRes
	if newRes <= target {
		if evaluate1(target, numbers, index+1, newRes) {
			return true
		}
	}

	newRes = currRes * currNum
	if newRes <= target {
		return evaluate1(target, numbers, index+1, newRes)
	}
	return false
}
