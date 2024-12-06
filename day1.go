package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day1() {
	file, err := os.Open("./day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	l := make([]int, 1000)
	r := make([]int, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		x1, err := strconv.Atoi(lineSplit[0])
		if err != nil {
			panic(err)
		}
		x2, err := strconv.Atoi(lineSplit[3])
		if err != nil {
			panic(err)
		}
		l = append(l, x1)
		r = append(r, x2)
	}

	slices.Sort(l)
	slices.Sort(r)

	fmt.Println(Part1(l, r))
	fmt.Println(Part2(l, r))
}

func Part1(l []int, r []int) int {
	sum := 0

	for i := range l {
		if l[i] < r[i] {
			sum = sum + (r[i] - l[i])
		} else {
			sum = sum + (l[i] - r[i])
		}
	}
	return sum
}

func Part2(l []int, r []int) int {
	counts := make(map[int]int)
	for _, v := range r {
		counts[v] = counts[v] + 1
	}

	sum := 0
	for _, v := range l {
		sum = sum + (v * counts[v])
	}
	return sum
}
