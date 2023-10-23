package db

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
