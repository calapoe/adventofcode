package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayTen1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]int, len(lines))
	for i, line := range lines {
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			grid[i] = append(grid[i], num)
		}
	}

	trailheads := [][2]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				trailheads = append(trailheads, [2]int{i, j})
			}
		}
	}

	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	out := 0
	for _, trailhead := range trailheads {
		// move in all directions to get to 9
		visited := make(map[[2]int]bool)
		out += computeScore(grid, trailhead[0], trailhead[1], directions, 0, visited)
	}
	fmt.Println("Output Day 10 Part 1: ", out)
}

func dayTen2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]int, len(lines))
	for i, line := range lines {
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			grid[i] = append(grid[i], num)
		}
	}

	trailheads := [][2]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				trailheads = append(trailheads, [2]int{i, j})
			}
		}
	}

	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	out := 0
	for _, trailhead := range trailheads {
		out += computeRating(grid, trailhead[0], trailhead[1], directions, 0)
	}
	fmt.Println("Output Day 10 Part 2: ", out)
}

func computeRating(grid [][]int, row int, column int, directions [][2]int, index int) int {
	if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[0]) {
		return 0
	}

	if grid[row][column] != index {
		return 0
	}

	if index == 9 {
		return 1
	}
	out := 0
	for _, direction := range directions {
		nextRow := row + direction[0]
		nextCol := column + direction[1]
		out += computeRating(grid, nextRow, nextCol, directions, index+1)
	}
	return out
}

func computeScore(grid [][]int, row int, column int, directions [][2]int, index int, visited map[[2]int]bool) int {
	if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[0]) {
		return 0
	}

	if grid[row][column] != index {
		return 0
	}

	currPos := [2]int{row, column}
	if visited[currPos] {
		return 0
	}
	visited[currPos] = true

	if index == 9 {
		return 1
	}
	out := 0
	for _, direction := range directions {
		nextRow := row + direction[0]
		nextCol := column + direction[1]
		out += computeScore(grid, nextRow, nextCol, directions, index+1, visited)
	}
	return out
}
