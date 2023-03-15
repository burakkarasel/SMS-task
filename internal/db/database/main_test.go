package db

import (
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

// TestMain sets up the DB connection for testing
func TestMain(m *testing.M) {
	url := flag.String("url", "DB source URL", "Add your DB URL")
	flag.Parse()

	testDB, err := sql.Open("postgres", *url)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
