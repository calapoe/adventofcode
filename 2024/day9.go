package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayNine1(input string) {
	line := strings.TrimSpace(input)

	id := 0
	freeSpace := false
	var groups []rune
	for i := 0; i < len(line); i++ {
		num, _ := strconv.Atoi(string(line[i]))
		if freeSpace {
			for j := 0; j < num; j++ {
				groups = append(groups, '.')
			}
		} else {
			for j := 0; j < num; j++ {
				groups = append(groups, rune('0'+id))
			}
			id++
		}
		freeSpace = !freeSpace
	}

	leftIndex, rightIndex := 0, len(groups)-1
	for groups[leftIndex] != '.' {
		leftIndex++
	}
	for groups[rightIndex] == '.' {
		rightIndex--
	}
	for leftIndex <= rightIndex {
		groups[leftIndex] = groups[rightIndex]
		groups[rightIndex] = '.'
		for groups[leftIndex] != '.' {
			leftIndex++
		}
		for groups[rightIndex] == '.' {
			rightIndex--
		}
	}

	out := 0
	for i := 0; i < len(groups); i++ {
		if groups[i] != '.' {
			out += int(groups[i]-'0') * i
		}
	}
	fmt.Println("Output Day 9 Part 1: ", out)
}

func dayNine2(input string) {
	line := strings.TrimSpace(input)

	id := 0
	freeSpace := false
	var groups []rune
	for i := 0; i < len(line); i++ {
		num, _ := strconv.Atoi(string(line[i]))
		if freeSpace {
			for j := 0; j < num; j++ {
				groups = append(groups, '.')
			}
		} else {
			for j := 0; j < num; j++ {
				groups = append(groups, rune('0'+id))
			}
			id++
		}
		freeSpace = !freeSpace
	}

	//TODO: move each file exactly once in order of decreasing file number

	out := 0
	for i := 0; i < len(groups); i++ {
		if groups[i] != '.' {
			out += int(groups[i]-'0') * i
		}
		fmt.Println(i, int(groups[i]-'0'), out)
	}
	fmt.Println("Output Day 9 Part 2: ", out)
}
