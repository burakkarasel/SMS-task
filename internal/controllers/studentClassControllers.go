package controllers

import (
	"database/sql"
	"github.com/burakkarasel/SMS-task/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

// createStudentClass creates a new mapping between given student and class
func (server *Server) createStudentClass(ctx *gin.Context) {
	// creating request and response instances
	var req models.CreateStudentClassApiParams
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

	// then we create new CreateStudentClassParams
	arg := models.CreateStudentClassParams{
		StudentId: req.StudentId,
		ClassId:   req.ClassId,
	}

	studentClass, err := server.store.CreateStudentClass(ctx, arg)

	// if any error occurs we return http internal server error with error
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				res = models.GenericResponse{
					Success:    false,
					StatusCode: http.StatusConflict,
					Messages:   []string{"Couldn't find data with given ID's"},
					Data:       nil,
				}
				ctx.JSON(http.StatusConflict, generateResponse(res))
				return
			}
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
		StatusCode: http.StatusCreated,
		Messages:   []string{},
		Data:       studentClass,
	}
	ctx.JSON(http.StatusCreated, generateResponse(res))
	return
}

// getStudentClass returns the student class mapping with given student ID and class ID
func (server *Server) getStudentClass(ctx *gin.Context) {
	// creating request and response instances
	var req models.GetOneStudentClassApiParams
	var res models.GenericResponse

	// if inputs are not valid we return status bad request with the error
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

	// then we create new GetOneStudentClassParam
	arg := models.GetOneStudentClassParam{
		StudentClassId: req.StudentClassId,
	}

	studentClass, err := server.store.GetStudentClass(ctx, arg)

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
		Data:       studentClass,
	}
	ctx.JSON(http.StatusOK, generateResponse(res))
	return
}

// listStudentClasses lists the mapping according to given parameters
func (server *Server) listStudentClasses(ctx *gin.Context) {
	// creating request and response instances
	var req models.ListStudentClassesApiParams
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

	if req.StudentId == 0 && req.ClassId == 0 {
		res = models.GenericResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Messages:   []string{"Both Student ID and Class ID cannot be 0"},
			Data:       nil,
		}
		ctx.JSON(http.StatusBadRequest, generateResponse(res))
		return
	}

	if req.StudentId > 0 && req.ClassId > 0 {
		res = models.GenericResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Messages:   []string{"Both Student ID and Class ID cannot be greater than 0"},
			Data:       nil,
		}
		ctx.JSON(http.StatusBadRequest, generateResponse(res))
		return
	}

	if req.StudentId > 0 {
		// then we create new ListClassesOfStudentParams
		arg := models.ListClassesOfStudentParams{
			StudentId: req.StudentId,
			Offset:    req.PageLimit * (req.PageId - 1),
			Limit:     req.PageLimit,
		}

		studentClasses, err := server.store.ListClassesOfStudent(ctx, arg)
		// if any error occurs we check the error
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
		if studentClasses.Student.Id == 0 {
			res = models.GenericResponse{
				Success:    false,
				StatusCode: http.StatusNotFound,
				Messages:   []string{"Student is not found with given ID"},
				Data:       nil,
			}
			ctx.JSON(http.StatusNotFound, generateResponse(res))
			return
		}
		// finally we generate the generic response and return it
		res = models.GenericResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Messages:   []string{},
			Data:       studentClasses,
		}
		ctx.JSON(http.StatusOK, generateResponse(res))
		return
	}

	// then we create new ListStudentsOfClassParams
	arg := models.ListStudentsOfClassParams{
		ClassId: req.ClassId,
		Offset:  req.PageLimit * (req.PageId - 1),
		Limit:   req.PageLimit,
	}

	classStudents, err := server.store.ListStudentsOfClass(ctx, arg)
	// if any error occurs we check the error
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

	if classStudents.Class.Id == 0 {
		res = models.GenericResponse{
			Success:    false,
			StatusCode: http.StatusNotFound,
			Messages:   []string{"Class is not found with given ID"},
			Data:       nil,
		}
		ctx.JSON(http.StatusNotFound, generateResponse(res))
		return
	}

	// finally we generate the generic response and return it
	res = models.GenericResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Messages:   []string{},
		Data:       classStudents,
	}
	ctx.JSON(http.StatusOK, generateResponse(res))
	return
}

func (server *Server) deleteStudentClass(ctx *gin.Context) {
	// creating request and response instances
	var req models.DeleteOneStudentClassApiParams
	var res models.GenericResponse

	// if inputs are not valid we return status bad request with the error
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

	// then we create new DeleteOneStudentClassParam
	arg := models.DeleteOneStudentClassParam{StudentClassId: req.StudentClassId}

	err := server.store.DeleteStudentClass(ctx, arg)

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

	ctx.JSON(http.StatusNoContent, nil)
	return
}
