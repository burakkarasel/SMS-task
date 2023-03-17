package db

import "database/sql"

// Store provides all DB functions we created, this interface also lets us create a mock store
type Store interface {
	Querier
}

// SQLStore provides all DB functions by implementing the Queries struct in it
type SQLStore struct {
	*Queries
}

// NewStore creates a new SQLStore instance for the project
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
	}
}
