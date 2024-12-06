package main

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

func followsRules(rules [][]string, print []string) (bool, int, int) {
	for _, rule := range rules {
		primary := slices.Index(print, rule[0])
		secondary := slices.Index(print, rule[1])
		if primary < 0 || secondary < 0 {
			continue
		}
		if primary < secondary {
			continue
		}
		return false, primary, secondary
	}
	return true, -1, -1
}

func fix(rules [][]string, print []string) []string {
	for {
		fixed, x, y := followsRules(rules, print)
		if fixed {
			return print
		}
		print[x], print[y] = print[y], print[x]
	}
}

func Day5() {
	input := getFileLines("./day5.txt")

	rules := make([][]string, 0)
	isRules := true

	sum := 0
	p2sum := 0
	for _, line := range input {
		if line == "" {
			isRules = false
		} else if isRules {
			split := strings.Split(line, "|")
			rules = append(rules, split)
		} else {
			printOrder := strings.Split(line, ",")
			valid, _, _ := followsRules(rules, printOrder)
			if valid {
				centerIdx := len(printOrder) / 2
				center, err := strconv.Atoi(printOrder[centerIdx])
				if err != nil {
					log.Fatal(err)
				}
				sum = sum + center
			} else {
				fixed := fix(rules, printOrder)
				centerIdx := len(fixed) / 2
				center, err := strconv.Atoi(fixed[centerIdx])
				if err != nil {
					log.Fatal(err)
				}
				p2sum = p2sum + center
			}
		}
	}
	log.Printf("Part 1: %d", sum)
	log.Printf("Part 2: %d", p2sum)
}
