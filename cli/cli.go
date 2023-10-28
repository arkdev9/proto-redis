package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/arkdev9/bad-redis/db"
)

// Buffers to store multi blocks
var buffer []string = make([]string, 0)
var multiMode bool = false

func Cli(input string) string {
	args := strings.Split(input, " ")
	// First arg should always be the command
	// Depending on the command, the number of args will vary
	if len(args) < 1 {
		return "Invalid number of args"
	}
	cmd := args[0]
	cmd = strings.ToUpper(cmd)

	if multiMode && cmd != "EXEC" {
		buffer = append(buffer, input)
		return "OK"
	}

	switch cmd {
	case "SET":
		// SET key value
		// Check for length of arguments. We need arg=2 and arg=3 to be able to set a key
		if len(args) != 3 {
			return "Invalid number of arguments"
		}
		db.Set(args[1], args[2])
	case "GET":
		if len(args) != 2 {
			return "Invalid number of arguments"
		}

		return db.Get(args[1])
	case "DEL":
		if len(args) != 2 {
			return "Invalid number of arguments"
		}
		return db.Del(args[1])
	case "INCR":
		if len(args) != 2 {
			return "Invalid number of arguments"
		}
		return db.Incr(args[1])
	case "INCRBY":
		if len(args) != 3 {
			return "Invalid number of arguments"
		}
		incrByVal, ok := strconv.Atoi(args[2])
		if ok != nil {
			return fmt.Sprintf("Invalid increment value\n%s", ok)
		}
		return db.Incr(args[1], incrByVal)
	case "MULTI":
		multiMode = true
		return "OK"
	case "EXEC":
		// Disable multi mode
		multiMode = false
		// Flush buffer and execute them by recursively running Cli with inputs
		fmt.Println("Executing...")
		for _, bufferInput := range buffer {
			fmt.Println(Cli(bufferInput))
		}
		return "OK"
	case "DISCARD":
		multiMode = false
		buffer = make([]string, 0)
		return "OK"
	case "EXIT":
		panic("Received exit")
	default:
		return "Invalid command"
	}
	return "OK"
}
