package cli

import (
	"fmt"
	"strings"

	"github.com/arkdev9/bad-redis/db"
)

func Cli(input string) string {
	args := strings.Split(input, " ")
	// First arg should always be the command
	// Depending on the command, the number of args will vary
	if len(args) < 2 {
		return "Invalid number of args"
	}
	cmd := args[1]
	switch cmd {
	case "SET":
		// SET key value
		// Check for length of arguments. We need arg=2 and arg=3 to be able to set a key
		if len(args) != 4 {
			return "Invalid number of arguments"
		}
		db.Set(args[2], args[3])
	case "GET":
		if len(args) != 3 {
			return "Invalid number of arguments"
		}
		// GET key
		return db.Get(args[2])
	// case "DEL":
	// 	// DEL key
	// case "EXIT":
	// 	// EXIT
	default:
		fmt.Println("Invalid command")
	}
	return "OK"
}
