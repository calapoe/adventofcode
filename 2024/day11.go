package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayEleven1(input string) {
	lines := strings.Fields(input)
	var numLines []int
	for _, numStr := range lines {
		num, _ := strconv.Atoi(numStr)
		numLines = append(numLines, num)
	}
	fmt.Println(numLines)
	for i := 0; i < 25; i++ {
		var newLines []int
		for _, num := range numLines {
			numStr := strconv.Itoa(num)
			if num == 0 {
				newLines = append(newLines, 1)
			} else if len(numStr)%2 == 0 {
				num1, _ := strconv.Atoi(numStr[:len(numStr)/2])
				num2, _ := strconv.Atoi(numStr[len(numStr)/2:])
				newLines = append(newLines, num1)
				newLines = append(newLines, num2)
			} else {
				newNum := num * 2024
				newLines = append(newLines, newNum)
			}
		}
		numLines = newLines
	}
	fmt.Println("Output Day 11 Part 1: ", len(numLines))
}

func dayEleven2(input string) {
	lines := strings.Fields(input)
	numMap := make(map[int]int)
	for _, numStr := range lines {
		num, _ := strconv.Atoi(numStr)
		numMap[num]++
	}
	fmt.Println(numMap)
	for i := 0; i < 75; i++ {
		newMap := make(map[int]int)
		for num, count := range numMap {
			numStr := strconv.Itoa(num)
			if num == 0 {
				newMap[1] += count
			} else if len(numStr)%2 == 0 {
				num1, _ := strconv.Atoi(numStr[:len(numStr)/2])
				num2, _ := strconv.Atoi(numStr[len(numStr)/2:])
				newMap[num1] += count
				newMap[num2] += count
			} else {
				newMap[num*2024] += count
			}
		}
		numMap = newMap
	}

	out := 0
	for _, count := range numMap {
		out += count
	}
	fmt.Println("Output Day 11 Part 2: ", out)
}
