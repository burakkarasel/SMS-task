package controllers

import (
	db "github.com/burakkarasel/SMS-task/internal/db/database"
	"github.com/burakkarasel/SMS-task/internal/models"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests and holds database connection
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new server with its dependencies
func NewServer(store db.Store) *Server {
	// first we create a new server with given Store
	server := &Server{store: store}

	// then we set the routes and return the server
	server.setRoutes()

	return server
}

// setRoutes sets routes for our service
func (server *Server) setRoutes() {
	router := gin.Default()

	// students endpoint
	router.POST("/api/v1/students", server.createStudent)
	router.GET("/api/v1/students", server.listStudents)
	router.GET("/api/v1/students/:id", server.getStudent)
	router.PATCH("/api/v1/students/:id", server.updateStudent)
	router.DELETE("/api/v1/students/:id", server.deleteStudent)

	// classes endpoint
	router.POST("/api/v1/classes", server.createClass)
	router.GET("/api/v1/classes", server.listClasses)
	router.GET("/api/v1/classes/:id", server.getClass)
	router.PATCH("/api/v1/classes/:id", server.updateClass)
	router.DELETE("/api/v1/classes/:id", server.deleteClass)

	// studentClasses endpoint
	router.POST("/api/v1/studentClasses", server.createStudentClass)
	router.GET("/api/v1/studentClasses", server.listStudentClasses)
	router.GET("/api/v1/studentClasses/:studentClassId", server.getStudentClass)
	router.DELETE("/api/v1/studentClasses/:studentClassId", server.deleteStudentClass)

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
