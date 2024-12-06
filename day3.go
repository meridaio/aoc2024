package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func calc(line string, r *regexp.Regexp) (sum int) {
	matches := r.FindAllStringSubmatch(line, -1)
	for _, nums := range matches {
		x1, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err)
		}
		x2, err := strconv.Atoi(nums[2])
		if err != nil {
			log.Fatal(err)
		}

		sum = sum + (x1 * x2)
	}
	return
}

func calc2(line string, r *regexp.Regexp) (sum int) {
	matches := r.FindAllStringSubmatchIndex(line, -1)
	for _, nums := range matches {
		pdo := strings.LastIndex(line[:nums[0]], "do()")
		pdont := strings.LastIndex(line[:nums[0]], "don't()")
		if pdo > pdont || pdont < 0 {
			x1, err := strconv.Atoi(line[nums[2]:nums[3]])
			if err != nil {
				log.Fatal(err)
			}
			x2, err := strconv.Atoi(line[nums[4]:nums[5]])
			if err != nil {
				log.Fatal(err)
			}

			sum = sum + (x1 * x2)
		}
	}
	return
}

func Day3() {
	file, err := os.Open("./day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	part1Sum := 0
	part2Sum := 0
	inputStr := ""
	for scanner.Scan() {
		inputStr = inputStr + scanner.Text()
	}
	part1Sum = part1Sum + calc(inputStr, r)

	p2 := inputStr
	for {
		i1 := strings.Index(p2, "don't()")
		i2 := strings.Index(p2, "do()")

		if i1 == -1 {
			break
		} else if i2 == -1 {
			p2 = p2[:i1] + "xDDD"
			break
		} else if i2 < i1 {
			p2 = strings.Replace(p2, "do()", "xDDD", 1)
		} else {
			pre := p2[:i1]
			post := p2[i2+4:]
			p2 = pre + "xDDD" + post
		}
	}
	log.Print(p2)
	part2Sum = part2Sum + calc2(inputStr, r)
	log.Printf("Part 1: %d", part1Sum)
	log.Printf("Part 2: %d", part2Sum)
}
