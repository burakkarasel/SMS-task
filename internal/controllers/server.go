package controllers

import (
	"github.com/burakkarasel/SMS-task/internal/db/database"
	"github.com/burakkarasel/SMS-task/internal/models"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests and holds database connection
type Server struct {
	store  database.Store
	router *gin.Engine
}

// NewServer creates a new server with its dependencies
func NewServer(store database.Store) *Server {
	// first we create a new server with given Store
	server := &Server{store: store}

	// then we set the routes an return the server
	server.setRoutes()

	return server
}

// setRoutes sets routes for our service
func (server *Server) setRoutes() {
	router := gin.Default()
	router.POST("/api/v1/students", server.createStudent)

	server.router = router
}

// Start func starts our server
func (server *Server) Start(port string) error {
	return server.router.Run(port)
}

// generateResponse generates a genericResponse instance and returns it
func generateResponse(response models.GenericResponse) gin.H {
	return gin.H{
		"response": response,
	}
}
