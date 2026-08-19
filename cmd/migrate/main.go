package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/student-ankitpandit/go-server/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: migrate <up | down>")
	}

	cfg := config.MustLoad()

	m, err := migrate.New(
		"file://migrations",
		cfg.DatabaseUrl)

	if err != nil {
		log.Fatalf("migrate.new %s", err)
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatalf("m.Up failed to up the migrations: %v", err)
		}
	case "down":
		if err := m.Down(); err != nil {
			log.Fatalf("m.Down failed to drop a table or coloumn: %v", err)
		}
	default:
		log.Fatalf("unknown command: %s", os.Args[1])
	}

	fmt.Println("migration ran successfully")
}
