package main

import (
	"log"
	"slices"
)

type PointHistory struct {
	Up    bool
	Down  bool
	Right bool
	Left  bool
}

func step(matrix [][]rune, visited [][]bool, pointHistory [][]PointHistory) (newMatrix [][]rune, loopDetected bool) {
	for x, row := range matrix {
		up := slices.Index(row, '^')
		down := slices.Index(row, 'v')
		right := slices.Index(row, '>')
		left := slices.Index(row, '<')

		if up > 0 {
			newX := x - 1

			//log.Printf("moving up from (%d, %d) to (%d, %d)", x, up, newX, up)
			visited[x][up] = true
			if newX < 0 {
				//log.Printf("out of bounds")
				return nil, false
			}
			if pointHistory[x][up].Up {
				return nil, true
			} else {
				pointHistory[x][up].Up = true
			}
			if matrix[newX][up] == '#' {
				//log.Printf("Encountered obstacle, turning right")
				matrix[x][up] = '>'
				return matrix, false
			}
			matrix[newX][up] = matrix[x][up]
			matrix[x][up] = '.'
			return matrix, false
		} else if down > 0 {
			newX := x + 1

			//log.Printf("moving down from (%d, %d) to (%d, %d)", x, down, newX, down)
			visited[x][down] = true
			if newX >= len(matrix) {
				//log.Printf("out of bounds")
				return nil, false
			}
			if pointHistory[x][down].Down {
				return nil, true
			} else {
				pointHistory[x][down].Down = true
			}
			if matrix[newX][down] == '#' {
				//log.Printf("Encountered obstacle, turning left")
				matrix[x][down] = '<'
				return matrix, false
			}
			matrix[newX][down] = matrix[x][down]
			matrix[x][down] = '.'
			return matrix, false
		} else if right > 0 {
			newY := right + 1

			//log.Printf("moving right from (%d, %d) to (%d, %d)", x, right, x, newY)
			visited[x][right] = true
			if newY >= len(matrix[0]) {
				//log.Printf("out of bounds")
				return nil, false
			}
			if pointHistory[x][right].Right {
				return nil, true
			} else {
				pointHistory[x][right].Right = true
			}
			if matrix[x][newY] == '#' {
				//log.Printf("Encountered obstacle, turning down")
				matrix[x][right] = 'v'
				return matrix, false
			}
			matrix[x][newY] = matrix[x][right]
			matrix[x][right] = '.'
			return matrix, false
		} else if left > 0 {
			newY := left - 1

			//log.Printf("moving left from (%d, %d) to (%d, %d)", x, left, x, newY)
			visited[x][left] = true
			if newY < 0 {
				//log.Printf("out of bounds")
				return nil, false
			}
			if pointHistory[x][left].Left {
				return nil, true
			} else {
				pointHistory[x][left].Left = true
			}
			if matrix[x][newY] == '#' {
				//log.Printf("Encountered obstacle, turning up")
				matrix[x][left] = '^'
				return matrix, false
			}
			matrix[x][newY] = matrix[x][left]
			matrix[x][left] = '.'
			return matrix, false
		} else {
			continue
		}
	}
	return nil, false
}

func walkMatrix(matrix [][]rune) [][]bool {
	visited := make([][]bool, len(matrix))
	pointHistory := make([][]PointHistory, len(matrix))
	for x, row := range matrix {
		visited[x] = make([]bool, len(row))
		pointHistory[x] = make([]PointHistory, len(row))
	}
	t := 0
	for matrix != nil {
		t = t + 1
		var loopDetected bool
		matrix, loopDetected = step(matrix, visited, pointHistory)
		if loopDetected {
			return nil
		}
	}
	return visited
}

func copyMatrix(matrix [][]rune) [][]rune {
	matrixCopy := make([][]rune, len(matrix))
	for x, v := range matrix {
		matrixCopy[x] = make([]rune, len(matrix))
		copy(matrixCopy[x], v)
	}
	return matrixCopy
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
	originalMatrix := copyMatrix(matrix)

	walked := walkMatrix(matrix)

	sum := 0
	for _, v := range walked {
		for _, b := range v {
			if b {
				sum = sum + 1
			}
		}
	}

	p2Sum := 0
	for i, v := range walked {
		for j, b := range v {
			if b {
				m := copyMatrix(originalMatrix)
				log.Printf("(%d, %d)", i, j)
				if m[i][j] == '.' {
					m[i][j] = '#'

					w := walkMatrix(m)
					if w == nil {
						//log.Print("loop detected")
						p2Sum = p2Sum + 1
					}
				}
			}
		}
	}

	log.Printf("Part 1: %d", sum)
	log.Printf("Part 2: %d", p2Sum)
}
