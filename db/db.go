package db

import (
	"fmt"
	"strconv"
)

// We will need to store type information somehow
var database map[string]string = make(map[string]string)

func Set(name, value string) string {
	database[name] = value
	return "OK"
}

func Get(name string) string {
	value, exists := database[name]
	if !exists {
		return "Error: Key does not exist"
	}
	return value
}

func Del(name string) string {
	_, exists := database[name]
	if !exists {
		return "(integer) 0"
	}
	delete(database, name)
	return "(integer) 1"
}

func Incr(name string, by ...int) string {
	value, exists := database[name]
	if !exists {
		return "Error: Key does not exist"
	}

	// Make sure that existing value can be parsed to an integer
	intVal, err := strconv.Atoi(value)
	if err != nil {
		return "Error: Could not parse existing value to integer before increment"
	}

	if len(by) > 0 {
		intVal += by[0]
	} else {
		intVal += 1
	}

	database[name] = fmt.Sprint(intVal)
	return "OK"
}
