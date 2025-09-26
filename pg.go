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

// Function to get parameters
func GetParams() map[string]string {

	// Help variables
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

	// Monta a connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s",
		*host, *port, *database, *user,
	)

	return pgVarsMap
}

func main() {
	fmt.Println(GetEnvVars()["PGPORT"])
}
