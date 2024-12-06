package main

import "log"

func d4p1TestDir(matrix [][]rune, x, y, xdir, ydir int) bool {
	xmas := []rune("XMAS")
	for xmasPos := 0; xmasPos < len(xmas); xmasPos++ {
		xt := x + (xmasPos * xdir)
		yt := y + (xmasPos * ydir)
		if xt < 0 || yt < 0 || xt >= len(matrix) || yt >= len(matrix[x]) {
			return false
		}
		if matrix[xt][yt] != xmas[xmasPos] {
			return false
		}
	}
	return true
}

func d4p1Test(matrix [][]rune, x, y int) (sum int) {
	if matrix[x][y] == 'X' {
		for tx := -1; tx <= 1; tx++ {
			for ty := -1; ty <= 1; ty++ {
				if d4p1TestDir(matrix, x, y, tx, ty) {
					sum = sum + 1
				}
			}
		}
	}
	return
}

func d4p2Test(matrix [][]rune, x, y int) bool {
	if matrix[x][y] != 'A' {
		return false
	}

	// this is extremely dumb and there's definitely a better way
	xt1 := x - 1
	yt1 := y + 1
	xt2 := x + 1
	yt2 := y - 1

	xt3 := x - 1
	yt3 := y - 1
	xt4 := x + 1
	yt4 := y + 1

	if xt1 < 0 || xt2 < 0 || yt1 < 0 || yt2 < 0 || xt1 >= len(matrix) || xt2 >= len(matrix) || yt1 >= len(matrix[0]) || yt2 >= len(matrix[0]) {
		return false
	}
	// xdd
	if ((matrix[xt1][yt1] == 'M' && matrix[xt2][yt2] == 'S') || (matrix[xt1][yt1] == 'S' && matrix[xt2][yt2] == 'M')) &&
		((matrix[xt3][yt3] == 'M' && matrix[xt4][yt4] == 'S') || (matrix[xt3][yt3] == 'S' && matrix[xt4][yt4] == 'M')) {
		return true
	}
	return false
}

func Day4() {
	input := getFileLines("./day4.txt")

	matrix := make([][]rune, len(input))
	for i, v := range input {
		cs := make([]rune, len(v))
		for j, c := range v {
			cs[j] = c
		}
		matrix[i] = cs
	}

	sum := 0
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[0]); y++ {
			sum = sum + d4p1Test(matrix, x, y)
		}
	}

	p2sum := 0
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[0]); y++ {
			if d4p2Test(matrix, x, y) {
				p2sum = p2sum + 1
				//printx(matrix, x, y)
			}
		}
	}
	log.Printf("Part 1: %d", sum)
	log.Printf("Part 2: %d", p2sum)
}
