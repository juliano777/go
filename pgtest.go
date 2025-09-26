package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/pflag"
)

func main() {
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

	// Monta a connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s",
		*host, *port, *database, *user,
	)

	// Conecta ao banco
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Executa a query SHOW port;
	var result string
	err = conn.QueryRow(context.Background(), "SHOW shared_buffers;").Scan(&result)
	if err != nil {
		log.Fatalf("Erro ao executar query: %v\n", err)
	}

	fmt.Printf("PostgreSQL está ouvindo na porta: %s\n", result)

}
