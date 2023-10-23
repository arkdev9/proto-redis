package main

import (
	"fmt"

	"github.com/arkdev9/bad-redis/cli"
)

func main() {
	for {
		fmt.Print("$ ")
		var input string
		fmt.Scan(&input)
		fmt.Println(cli.Cli(input))
	}
}
