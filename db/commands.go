package db

// We will need to store type information somehow
var database map[string]string = make(map[string]string)

func Set(name, value string) string {
	// Check if key already exists
	_, exists := database[name]
	if exists {
		return "ERR: Key already exists in the database"
	}

	database[name] = value
	return "OK"
}

func Get(name string) string {
	return database[name]
}
