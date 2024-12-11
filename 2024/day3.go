package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func dayThree1(input string) {
	matches := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`).FindAllStringSubmatch(input, -1)

	out := 0
	for _, match := range matches {
		numOne, _ := strconv.Atoi(match[1])
		numTwo, _ := strconv.Atoi(match[2])

		out += numOne * numTwo
	}
	fmt.Println("Output Day 3 Part 1: ", out)
}

func dayThree2(input string) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	out := 0
	do := true
	for _, match := range matches {
		fmt.Println(match[0], match[1], match[2])
		if match[0] == "do()" {
			do = true
		} else if match[0] == "don't()" {
			do = false
		} else {
			if do {
				numOne, _ := strconv.Atoi(match[1])
				numTwo, _ := strconv.Atoi(match[2])

				out += numOne * numTwo
			}
		}
	}
	fmt.Println("Output Day 3 Part 1: ", out)
}
