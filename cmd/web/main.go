package main

import (
	"database/sql"
	"flag"
	"github.com/burakkarasel/SMS-task/internal/controllers"
	"github.com/burakkarasel/SMS-task/internal/db/database"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// DB URL
	url := flag.String("url", "DB source URL", "Add your DB URL")
	port := flag.String("port", "Port for the server", "Add you Port value")
	flag.Parse()

	conn, err := sql.Open("postgres", *url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to DB")
	runDBMigration("file://internal/db/migration", *url)

	store := database.NewStore(conn)

	server := controllers.NewServer(store)
	log.Println("Starting at port: ", *port)

	if err := server.Start(*port); err != nil {
		log.Fatal("Couldn't start server: ", err)
	}
}

// runDBMigration runs the migrations at the start of the program
// if an error occurs exits the program
// if error is ErrNoChange moves on
func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("db migrated successfully")
}
