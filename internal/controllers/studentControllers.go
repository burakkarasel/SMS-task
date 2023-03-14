package controllers

import (
	"github.com/burakkarasel/SMS-task/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// createStudent handles post requests to create a new student
func (server *Server) createStudent(ctx *gin.Context) {
	// creating request and response instances
	var req models.CreateStudentApiParams
	var res models.GenericResponse

	// if inputs are not valid we return status bad request with the error
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res = models.GenericResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Messages:   []string{err.Error()},
			Data:       nil,
		}
		ctx.JSON(http.StatusBadRequest, generateResponse(res))
		return
	}

	// then we create new CreateStudentParams to insert student into DB
	arg := models.CreateStudentParams{
		FullName:   req.FullName,
		Year:       req.Year,
		Department: req.Department,
		Email:      req.Email,
	}

	student, err := server.store.CreateStudent(ctx, arg)

	// if any error occurs we return http internal server error with error
	if err != nil {
		res = models.GenericResponse{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Messages:   []string{err.Error()},
			Data:       nil,
		}
		ctx.JSON(http.StatusInternalServerError, generateResponse(res))
		return
	}

	// finally we generate the generic response and return it
	res = models.GenericResponse{
		Success:    true,
		StatusCode: http.StatusCreated,
		Messages:   []string{},
		Data:       student,
	}
	ctx.JSON(http.StatusCreated, generateResponse(res))
	return
}

// listStudents list students with given params
func (server *Server) listStudents(ctx *gin.Context) {
	// creating request and response instances
	var req models.ListStudentsApiParams
	var res models.GenericResponse

	// if inputs are not valid we return status bad request with the error
	if err := ctx.ShouldBindQuery(&req); err != nil {
		res = models.GenericResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Messages:   []string{err.Error()},
			Data:       nil,
		}
		ctx.JSON(http.StatusBadRequest, generateResponse(res))
		return
	}

	// then we create a new ListStudentsParams to execute the DB operation
	arg := models.ListStudentsParams{
		Offset: (req.PageId - 1) * req.PageLimit,
		Limit:  req.PageLimit,
	}

	students, err := server.store.ListStudents(ctx, arg)

	// if any error occurs we return http internal server error with error
	if err != nil {
		res = models.GenericResponse{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Messages:   []string{err.Error()},
			Data:       nil,
		}
		ctx.JSON(http.StatusInternalServerError, generateResponse(res))
		return
	}

	// finally we generate the generic response and return it
	res = models.GenericResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Messages:   []string{},
		Data:       students,
	}
	ctx.JSON(http.StatusOK, generateResponse(res))
	return
}
