package main

import (
	"fmt"
	"strings"
)

func dayEight1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	antennas := make(map[rune][][2]int)
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] != '.' {
				antennas[grid[r][c]] = append(antennas[grid[r][c]], [2]int{r, c})
			}
		}
	}
	antiNodes := make(map[[2]int]bool)
	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				firstPos := positions[i]
				secondPos := positions[j]

				row := secondPos[0] - firstPos[0]
				col := secondPos[1] - firstPos[1]

				antiNode1 := [2]int{firstPos[0] - row, firstPos[1] - col}
				antiNode2 := [2]int{secondPos[0] + row, secondPos[1] + col}

				if antiNode1[0] >= 0 && antiNode1[0] < len(grid) && antiNode1[1] >= 0 && antiNode1[1] < len(grid[0]) {
					antiNodes[antiNode1] = true
				}

				if antiNode2[0] >= 0 && antiNode2[0] < len(grid) && antiNode2[1] >= 0 && antiNode2[1] < len(grid[0]) {
					antiNodes[antiNode2] = true
				}
			}
		}
	}

	fmt.Println("Output Day 8 Part 1: ", len(antiNodes))
}

func dayEight2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	antennas := make(map[rune][][2]int)
	antiNodes := make(map[[2]int]bool)
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] != '.' {
				antennas[grid[r][c]] = append(antennas[grid[r][c]], [2]int{r, c})
				antiNodes[[2]int{r, c}] = true
			}
		}
	}
	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				firstPos := positions[i]
				secondPos := positions[j]

				row := secondPos[0] - firstPos[0]
				col := secondPos[1] - firstPos[1]

				antiNodeRow := firstPos[0] - row
				antiNodeCol := firstPos[1] - col

				for antiNodeRow >= 0 && antiNodeRow < len(grid) && antiNodeCol >= 0 && antiNodeCol < len(grid[0]) {
					antiNodes[[2]int{antiNodeRow, antiNodeCol}] = true
					antiNodeRow -= row
					antiNodeCol -= col
				}

				antiNodeRow = secondPos[0] + row
				antiNodeCol = secondPos[1] + col
				for antiNodeRow >= 0 && antiNodeRow < len(grid) && antiNodeCol >= 0 && antiNodeCol < len(grid[0]) {
					antiNodes[[2]int{antiNodeRow, antiNodeCol}] = true
					antiNodeRow += row
					antiNodeCol += col
				}
			}
		}
	}

	fmt.Println("Output Day 8 Part 2: ", len(antiNodes))
}
