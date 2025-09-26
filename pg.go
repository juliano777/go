package main

import (
	"fmt"
	"os"
)

// Function to get environment variables
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

// Function to get the default values
func GetDefaultValues() map[string]string {

	defaultValues := map[string]string{
		"PGHOST":     "",
		"PGHOSTADDR": "",
		"PGPORT":     "5432",
		"PGDATABASE": "",
		"PGUSER":     "",
		"PGPASSWORD": "",
		"PGPASSFILE": "~/.pgpass",
	}

	pgVarsMap := make(map[string]string)

	for k, v := range defaultValues {

		if GetEnvVars()[k] != "" {
			pgVarsMap[k] = GetEnvVars()[k]
		} else {
			pgVarsMap[k] = v
		}
	}

	return pgVarsMap
}

func main() {
	fmt.Println(GetDefaultValues())
}
