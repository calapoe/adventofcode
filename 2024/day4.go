package main

import (
	"fmt"
	"strings"
)

func dayFour1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	out := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 'X' {
				out += searchForXMAS(r, c, grid)
			}
		}
	}
	fmt.Println("Output Day 4 Part 1: ", out)
}

func dayFour2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	out := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 'A' {
				out += searchForMAS(r, c, grid)
			}
		}
	}
	fmt.Println("Output Day 4 Part 2: ", out)
}

func searchForMAS(r int, c int, g [][]rune) int {
	var dirs = [][2][2]int{
		{
			{-1, -1},
			{1, 1},
		},
		{
			{1, -1},
			{-1, 1},
		},
	}
	numXmas := 0
	isXmas := true
	for _, d := range dirs {
		rD1 := r + d[0][0]
		cD1 := c + d[0][1]

		rD2 := r + d[1][0]
		cD2 := c + d[1][1]
		if rD1 < 0 || rD1 >= len(g) || cD1 < 0 || cD1 >= len(g[rD1]) || rD2 < 0 || rD2 >= len(g) ||
			cD2 < 0 || cD2 >= len(g[rD2]) {
			isXmas = false
			break
		}
		if (g[rD1][cD1] == 'M' && g[rD2][cD2] == 'S') || (g[rD1][cD1] == 'S' && g[rD2][cD2] == 'M') {
			continue
		}
		isXmas = false
		break
	}
	if isXmas {
		numXmas++
	}
	return numXmas
}

func searchForXMAS(r int, c int, g [][]rune) int {
	word := "XMAS"
	var dirs = [][2]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}
	numXmas := 0
	for _, d := range dirs {
		isXmas := true
		for i := 1; i < len(word); i++ {
			rN := r + d[0]*i
			cN := c + d[1]*i
			if rN < 0 || rN >= len(g) || cN < 0 || cN >= len(g[rN]) {
				isXmas = false
				break
			}
			if g[rN][cN] != rune(word[i]) {
				isXmas = false
				break
			}
		}
		if isXmas {
			numXmas++
		}
	}
	return numXmas
}
