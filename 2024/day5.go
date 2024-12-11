package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func dayFive1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n\n")

	ordering := strings.Split(lines[0], "\n")
	orderMap := make(map[int][]int)
	for _, parts := range ordering {
		fields := strings.Split(parts, "|")
		left, _ := strconv.Atoi(fields[0])
		right, _ := strconv.Atoi(fields[1])
		orderMap[left] = append(orderMap[left], right)
	}

	out := 0
	for _, parts := range strings.Split(lines[1], "\n") {
		isSafe := true
		pages := strings.Split(parts, ",")
		var newPages []int
		for _, page := range pages {
			num, _ := strconv.Atoi(page)
			newPages = append(newPages, num)
		}

		for i := 0; i < len(newPages)-1; i++ {
			pagesAfter, _ := orderMap[newPages[i]]

			for j := i + 1; j < len(pages); j++ {
				if !slices.Contains(pagesAfter, newPages[j]) {
					isSafe = false
					break
				}
			}
		}
		if isSafe {
			out += newPages[len(newPages)/2]
		}
	}
	fmt.Println("Output Day 5 Part 1: ", out)
}

func dayFive2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n\n")

	ordering := strings.Split(lines[0], "\n")
	orderMap := make(map[int][]int)
	for _, parts := range ordering {
		fields := strings.Split(parts, "|")
		left, _ := strconv.Atoi(fields[0])
		right, _ := strconv.Atoi(fields[1])
		orderMap[left] = append(orderMap[left], right)
	}

	out := 0
	for _, parts := range strings.Split(lines[1], "\n") {
		isSafe := true
		pages := strings.Split(parts, ",")
		var newPages []int
		for _, page := range pages {
			num, _ := strconv.Atoi(page)
			newPages = append(newPages, num)
		}

		for i := 0; i < len(newPages)-1; i++ {
			pagesAfter, _ := orderMap[newPages[i]]

			for j := i + 1; j < len(pages); j++ {
				if !slices.Contains(pagesAfter, newPages[j]) {
					isSafe = false
					break
				}
			}
		}
		if !isSafe {
			ordered := []int{}
			remaining := map[int]bool{}
			inDegree := map[int]int{}
			for _, num := range newPages {
				remaining[num] = true

				for _, dep := range orderMap[num] {
					inDegree[dep]++
				}
			}
			queue := []int{}
			for num := range remaining {
				if inDegree[num] == 0 {
					queue = append(queue, num)
				}
			}
			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]
				if slices.Contains(newPages, current) {
					ordered = append(ordered, current)
				}
				delete(remaining, current)

				for _, neighbor := range orderMap[current] {
					inDegree[neighbor]--
					if inDegree[neighbor] == 0 {
						queue = append(queue, neighbor)
					}
				}
			}

			out += ordered[len(ordered)/2]
		}
	}
	fmt.Println("Output Day 5 Part 2: ", out)
}
