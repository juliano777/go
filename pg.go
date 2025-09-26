package main

import (
	"fmt"
	"os"
)

func GetEnvVars() map[string]string {

	// Array of PostgreSQL environment variables
	pgVars := [7]string{
		"PGHOST",
		"PGHOSTADDR",
		"PGPORT",
		"PGDATABASE",
		"PGUSER",
		"PGPASSWORD",
		"PGPASSFILE",
	}

	// Map variable
	pgVarsMap := make(map[string]string)

	// Iterate over pgVars to get the variable names and its values
	for _, i := range pgVars {
		pgVarsMap[i] = os.Getenv(i)
	}

	return pgVarsMap
}

func main() {
	fmt.Println(GetEnvVars()["PGPORT"])
}
