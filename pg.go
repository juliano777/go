package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
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

func getParameters() {
	// Variáveis de help
	h_host := "Servidor do PostgreSQL (default vazio)"
	h_port := "Porta do PostgreSQL (default 5432)"
	h_user := "Usuário do PostgreSQL (default = usuário do SO)"
	h_database := "Nome do banco (default = usuário do banco)"

	// Flags curtas e longas
	host := pflag.StringP("host", "H", "", h_host)
	port := pflag.IntP("port", "p", 5432, h_port)
	user := pflag.StringP("user", "U", os.Getenv("USER"), h_user)
	database := pflag.StringP("database", "d", "", h_database)

	pflag.Parse()

	// Se não foi passado --database explicitamente → herda de --user (já parseado)
	if !pflag.CommandLine.Changed("database") {
		*database = *user
	}
}

func main() {
	fmt.Println(GetDefaultValues())
}
