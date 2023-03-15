package controllers

import (
	"os"
	"testing"

	db "github.com/burakkarasel/SMS-task/internal/db/database"
	"github.com/gin-gonic/gin"
)

// newTestServer creates a new test server for our tests
func newTestServer(t *testing.T, store db.Store) *Server {
	server := NewServer(store)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
