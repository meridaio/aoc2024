package main

import (
	"log"
	"slices"
)

func step(matrix [][]rune, visited [][]bool) [][]rune {
	for x, row := range matrix {
		up := slices.Index(row, '^')
		down := slices.Index(row, 'v')
		right := slices.Index(row, '>')
		left := slices.Index(row, '<')

		if up > 0 {
			newX := x - 1

			log.Printf("moving up from (%d, %d) to (%d, %d)", x, up, newX, up)
			visited[x][up] = true
			if newX >= len(matrix) {
				log.Printf("out of bounds")
				return nil
			}
			if matrix[newX][up] == '#' {
				log.Printf("Encountered obstacle, turning right")
				matrix[x][up] = '>'
				return matrix
			}
			matrix[newX][up] = matrix[x][up]
			matrix[x][up] = '.'
			return matrix
		} else if down > 0 {
			newX := x + 1

			log.Printf("moving down from (%d, %d) to (%d, %d)", x, down, newX, down)
			visited[x][down] = true
			if newX < 0 {
				log.Printf("out of bounds")
				return nil
			}
			if matrix[newX][down] == '#' {
				log.Printf("Encountered obstacle, turning left")
				matrix[x][down] = '<'
				return matrix
			}
			matrix[newX][down] = matrix[x][down]
			matrix[x][down] = '.'
			return matrix
		} else if right > 0 {
			newY := right + 1

			log.Printf("moving right from (%d, %d) to (%d, %d)", x, right, x, newY)
			visited[x][right] = true
			if newY >= len(matrix[0]) {
				log.Printf("out of bounds")
				return nil
			}
			if matrix[x][newY] == '#' {
				log.Printf("Encountered obstacle, turning down")
				matrix[x][right] = 'v'
				return matrix
			}
			matrix[x][newY] = matrix[x][right]
			matrix[x][right] = '.'
			return matrix
		} else if left > 0 {
			newY := left - 1

			log.Printf("moving left from (%d, %d) to (%d, %d)", x, left, x, newY)
			visited[x][left] = true
			if newY < 0 {
				log.Printf("out of bounds")
				return nil
			}
			if matrix[x][newY] == '#' {
				log.Printf("Encountered obstacle, turning up")
				matrix[x][left] = '^'
				return matrix
			}
			matrix[x][newY] = matrix[x][left]
			matrix[x][left] = '.'
			return matrix
		} else {
			continue
		}
	}
	return nil
}

func walkMatrix(matrix [][]rune) [][]bool {
	visited := make([][]bool, len(matrix))
	for x, row := range matrix {
		visited[x] = make([]bool, len(row))
	}
	t := 0
	for matrix != nil {
		t = t + 1
		matrix = step(matrix, visited)
	}
	return visited
}

func Day6() {
	input := getFileLines("./day6.txt")

	matrix := make([][]rune, len(input))
	for x, line := range input {
		row := make([]rune, len(line))
		for y, c := range line {
			row[y] = c
		}
		matrix[x] = row
	}
	walked := walkMatrix(matrix)

	sum := 0
	for _, v := range walked {
		for _, b := range v {
			if b {
				sum = sum + 1
			}
		}
	}
	log.Printf("Part 1: %d", sum)
}
