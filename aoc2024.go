package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type AocDay struct {
	Day   int
	Solve func()
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
	}

	commands := make([]*cli.Command, len(days))
	for i, v := range days {
		commands[i] = &cli.Command{
			Name:    fmt.Sprintf("day%d", v.Day),
			Aliases: []string{fmt.Sprint(v.Day)},
			Action: func(ctx *cli.Context) error {
				v.Solve()
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
