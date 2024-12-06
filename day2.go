package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Level = int

func isSafe(levels []Level) bool {
	increasing := levels[0] < levels[1]

	for i := range levels {
		if i == len(levels)-1 {
			return true
		}

		if increasing {
			diff := levels[i+1] - levels[i]
			if diff > 3 || diff < 1 {
				return false
			}
		} else {
			diff := levels[i] - levels[i+1]
			if diff > 3 || diff < 1 {
				return false
			}
		}
	}
	return true
}

func problemDampener(levels []Level) bool {
	log.Printf("%v", levels)
	for i := range levels {
		remd := removeIndex(levels, i)
		log.Printf("%d %v", i, remd)
		if isSafe(remd) {
			return true
		}
	}
	return false
}

func removeIndex(arr []int, pos int) []int {
	ret := make([]int, 0, len(arr)-1)
	ret = append(ret, arr[:pos]...)
	return append(ret, arr[pos+1:]...)
}

func Day2() {
	file, err := os.Open("./day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	part1Count := 0
	part2Count := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		levels := make([]Level, len(lineSplit))
		for i, v := range lineSplit {
			l, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			levels[i] = l
		}

		if isSafe(levels) {
			part1Count = part1Count + 1
			part2Count = part2Count + 1
		} else if problemDampener(levels) {
			log.Printf("%v", levels)
			part2Count = part2Count + 1
		}
	}
	log.Printf("Part 1: %d", part1Count)
	log.Printf("Part 2: %d", part2Count)
}
