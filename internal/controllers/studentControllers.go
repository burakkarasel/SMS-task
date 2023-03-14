package controllers

import (
	"github.com/burakkarasel/SMS-task/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) createStudent(ctx *gin.Context) {
	var req models.CreateStudentApiParams
	var resp models.GenericResponse

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp = models.GenericResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Messages:   []string{err.Error()},
			Data:       nil,
		}
		ctx.JSON(http.StatusBadRequest, generateResponse(resp))
		return
	}

	arg := models.CreateStudentParams{
		FullName:   req.FullName,
		Year:       req.Year,
		Department: req.Department,
		Email:      req.Email,
	}

	student, err := server.store.CreateStudent(ctx, arg)
	if err != nil {
		resp = models.GenericResponse{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Messages:   []string{err.Error()},
			Data:       nil,
		}
		ctx.JSON(http.StatusInternalServerError, generateResponse(resp))
		return
	}

	resp = models.GenericResponse{
		Success:    true,
		StatusCode: http.StatusCreated,
		Messages:   []string{},
		Data:       student,
	}
	ctx.JSON(http.StatusCreated, generateResponse(resp))
	return
}
