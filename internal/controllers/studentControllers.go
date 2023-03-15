package controllers

import (
	"database/sql"
	"fmt"
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

	// then we create new CreateStudentParams
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

	// then we create a new ListStudentsParams
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

// getStudent gets one student from the DB
func (server *Server) getStudent(ctx *gin.Context) {
	// creating request and response instances
	var req models.GetOneStudentApiParam
	var res models.GenericResponse

	if err := ctx.ShouldBindUri(&req); err != nil {
		res = models.GenericResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Messages:   []string{err.Error()},
			Data:       nil,
		}
		ctx.JSON(http.StatusBadRequest, generateResponse(res))
		return
	}

	// then we create a new GetOneStudentParam
	arg := models.GetOneStudentParam{
		Id: req.Id,
	}

	student, err := server.store.GetStudent(ctx, arg)

	// if any error occurs we check the error
	if err != nil {
		// if error equals to ErrNoRows we return 400
		if err == sql.ErrNoRows {
			res = models.GenericResponse{
				Success:    false,
				StatusCode: http.StatusNotFound,
				Messages:   []string{"Couldn't find Student with given ID"},
				Data:       nil,
			}
			ctx.JSON(http.StatusNotFound, generateResponse(res))
			return
		}
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
		Data:       student,
	}
	ctx.JSON(http.StatusOK, generateResponse(res))
	return

}

// updateStudent updates the student with given params
func (server *Server) updateStudent(ctx *gin.Context) {
	// creating request and response instances
	var reqBody models.UpdateStudentApiBodyParams
	var reqUri models.UpdateStudentApiUriParam
	var res models.GenericResponse

	// checking for both URI and JSON if inputs are not valid we return status bad request with the error
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		res = models.GenericResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Messages:   []string{err.Error()},
			Data:       nil,
		}
		ctx.JSON(http.StatusBadRequest, generateResponse(res))
		return
	}

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		res = models.GenericResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Messages:   []string{"Invalid Student ID"},
			Data:       nil,
		}
		ctx.JSON(http.StatusBadRequest, generateResponse(res))
		return
	}

	// then we create a new UpdateStudentParams
	arg := models.UpdateStudentParams{
		Id:         reqUri.Id,
		FullName:   reqBody.FullName,
		Year:       reqBody.Year,
		Department: reqBody.Department,
		Email:      reqBody.Email,
	}

	student, err := server.store.UpdateStudent(ctx, arg)

	// if any error occurs we check the error
	if err != nil {
		// if error equals to ErrNoRows we return 400
		if err == sql.ErrNoRows {
			res = models.GenericResponse{
				Success:    false,
				StatusCode: http.StatusNotFound,
				Messages:   []string{"Couldn't find Student with given ID"},
				Data:       nil,
			}
			ctx.JSON(http.StatusNotFound, generateResponse(res))
			return
		}
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
		Data:       student,
	}
	fmt.Println(res)
	ctx.JSON(http.StatusOK, generateResponse(res))
	return
}

// deleteStudent deletes student by given ID
func (server *Server) deleteStudent(ctx *gin.Context) {
	// creating request and response instances
	var req models.DeleteStudentApiParam
	var res models.GenericResponse

	// if the input are not valid we return status bad request with the error
	if err := ctx.ShouldBindUri(&req); err != nil {
		res = models.GenericResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Messages:   []string{err.Error()},
			Data:       nil,
		}
		ctx.JSON(http.StatusBadRequest, generateResponse(res))
		return
	}

	// then we create a new DeleteOneStudentParam
	arg := models.DeleteOneStudentParam{
		Id: req.Id,
	}

	err := server.store.DeleteStudent(ctx, arg)

	// if any error occurs we return http internal server error with error
	if err != nil {
		if err == sql.ErrNoRows {
			res = models.GenericResponse{
				Success:    false,
				StatusCode: http.StatusNotFound,
				Messages:   []string{"Couldn't find Student with given ID"},
				Data:       nil,
			}
			ctx.JSON(http.StatusNotFound, generateResponse(res))
			return
		}
		res = models.GenericResponse{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Messages:   []string{err.Error()},
			Data:       nil,
		}
		ctx.JSON(http.StatusInternalServerError, generateResponse(res))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
	return
}
