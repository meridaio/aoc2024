package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type AocDay struct {
	Day   int
	Solve func() (int, int)
}

func getFileLines(f string) []string {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func main() {
	days := []AocDay{
		{
			Day:   1,
			Solve: Day1,
		},
		{
			Day:   2,
			Solve: Day2,
		},
		{
			Day:   3,
			Solve: Day3,
		},
		{
			Day:   4,
			Solve: Day4,
		},
		{
			Day:   5,
			Solve: Day5,
		},
		{
			Day:   6,
			Solve: Day6,
		},
		{
			Day:   7,
			Solve: Day7,
		},
		{
			Day:   8,
			Solve: Day8,
		},
	}

	commands := make([]*cli.Command, len(days))
	for i, v := range days {
		commands[i] = &cli.Command{
			Name:    fmt.Sprintf("day%d", v.Day),
			Aliases: []string{fmt.Sprint(v.Day)},
			Action: func(ctx *cli.Context) error {
				fmt.Printf("Day %d:\n", v.Day)
				p1, p2 := v.Solve()

				fmt.Printf("\tPart 1: %d\n", p1)
				fmt.Printf("\tPart 2: %d\n", p2)
				return nil
			},
		}
	}

	app := &cli.App{
		Commands: commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
