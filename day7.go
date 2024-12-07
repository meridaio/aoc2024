package main

import (
	"log"
	"strconv"
	"strings"
)

type CalibrationEquation struct {
	Sum            int
	EquationValues []int
}

type Operator = func(int, int) int

func add(x, y int) int {
	return x + y
}

func mult(x, y int) int {
	return x * y
}

func cat(x, y int) int {
	z, err := strconv.Atoi(strconv.Itoa(x) + strconv.Itoa(y))
	if err != nil {
		log.Fatal(err)
	}
	return z
}

var part1Funcs = []Operator{
	add,
	mult,
}

var part2Funcs = []Operator{
	add,
	mult,
	cat,
}

func evalPossibilities(x, y int, operators []Operator) []int {
	r := make([]int, len(operators))
	for i, op := range operators {
		r[i] = op(x, y)
	}
	return r
}

func test(ce CalibrationEquation, operators []Operator) bool {
	possibleSums := make([]int, 1)
	possibleSums[0] = ce.EquationValues[0]

	for i := range ce.EquationValues {
		if i == 0 {
			continue
		}
		opCount := len(operators)
		newPos := make([]int, len(possibleSums)*opCount)
		for j, p := range possibleSums {
			ps := evalPossibilities(p, ce.EquationValues[i], operators)
			for l, up := range ps {
				newPos[(j*opCount)+l] = up
			}
		}
		possibleSums = newPos
	}

	for _, v := range possibleSums {
		if v == ce.Sum {
			return true
		}
	}
	return false
}

func parseLine(line string) CalibrationEquation {
	s := strings.Split(line, ": ")
	sum, err := strconv.Atoi(s[0])
	if err != nil {
		log.Fatal(err)
	}

	multiples := strings.Split(s[1], " ")
	xs := make([]int, len(multiples))
	for i, v := range multiples {
		x, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		xs[i] = x
	}
	return CalibrationEquation{
		Sum:            sum,
		EquationValues: xs,
	}
}

func Day7() (int, int) {
	input := getFileLines("./day7.txt")
	equations := make([]CalibrationEquation, len(input))
	for i, line := range input {
		equations[i] = parseLine(line)
	}

	sum := 0
	for _, eq := range equations {
		if test(eq, part1Funcs) {
			sum = sum + eq.Sum
		}
	}

	sum2 := 0
	for _, eq := range equations {
		if test(eq, part2Funcs) {
			sum2 = sum2 + eq.Sum
		}
	}
	return sum, sum2
}
