package main

import (
	"fmt"
	"strings"
)

func daySix1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	curRow, curCol, curDir := 0, 0, 3
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '^' {
				curRow = r
				curCol = c
			}
		}
	}
	visited := make(map[[2]int]bool)
	for {
		visited[[2]int{curRow, curCol}] = true
		direction := directions[curDir]
		nextRow := curRow + direction[0]
		nextCol := curCol + direction[1]

		if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
			break
		}

		if grid[nextRow][nextCol] == '#' {
			curDir = (curDir + 1) % 4
			direction = directions[curDir]
			nextRow = curRow + direction[0]
			nextCol = curCol + direction[1]

			if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[curRow]) {
				break
			}
		}
		curRow = nextRow
		curCol = nextCol
	}

	fmt.Println("Output Day 6 Part 1: ", len(visited))

}

func daySix2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	curRow, curCol, curDir := 0, 0, 3
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '^' {
				curRow = r
				curCol = c
			}
		}
	}
	out := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] != '.' {
				continue
			}
			grid[r][c] = '#'
			if searchForLoop(grid, curRow, curCol, curDir, directions) {
				out += 1
			}
			grid[r][c] = '.'
		}
	}

	fmt.Println("Output Day 6 Part 2: ", out)
}

func searchForLoop(grid [][]rune, curRow, curCol, curDir int, directions [][2]int) bool {
	visited := make(map[[3]int]bool)
	inLoop := false
	for {
		guardPos := [3]int{curRow, curCol, curDir}
		if visited[guardPos] {
			inLoop = true
			break
		}
		visited[guardPos] = true

		nextRow := curRow + directions[curDir][0]
		nextCol := curCol + directions[curDir][1]
		if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
			break
		}

		if grid[nextRow][nextCol] == '#' {
			curDir = (curDir + 1) % 4
		} else {
			curRow = nextRow
			curCol = nextCol
		}
	}
	return inLoop
}
