package controllers

import (
	"database/sql"
	"github.com/burakkarasel/SMS-task/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// createClass creates a new class in the DB
func (server *Server) createClass(ctx *gin.Context) {
	// creating request and response instances
	var req models.CreateClassApiParams
	var res models.GenericResponse

	// if inputs are not valid we return status bad request with the error
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res = models.CreateGenericResponse(false, http.StatusBadRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, generateResponse(res))
		return
	}

	// then we create new CreateClassParams
	arg := models.CreateClassParams{
		Name:      req.Name,
		Professor: req.Professor,
	}

	class, err := server.store.CreateClass(ctx, arg)

	// if any error occurs we return http internal server error with error
	if err != nil {
		res = models.CreateGenericResponse(false, http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, generateResponse(res))
		return
	}

	// finally we generate the generic response and return it
	res = models.CreateGenericResponse(true, http.StatusCreated, "", class)
	ctx.JSON(http.StatusCreated, generateResponse(res))
	return
}

// listClasses lists the classes with pagination
func (server *Server) listClasses(ctx *gin.Context) {
	// creating request and response instances
	var req models.ListClassesApiParams
	var res models.GenericResponse

	// if inputs are not valid we return status bad request with the error
	if err := ctx.ShouldBindQuery(&req); err != nil {
		res = models.CreateGenericResponse(false, http.StatusBadRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, generateResponse(res))
		return
	}

	// then we create new ListClassesParams
	arg := models.ListClassesParams{
		Limit:  req.PageLimit,
		Offset: (req.PageId - 1) * req.PageLimit,
	}

	classes, err := server.store.ListClasses(ctx, arg)

	// if any error occurs we return http internal server error with error
	if err != nil {
		res = models.CreateGenericResponse(false, http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, generateResponse(res))
		return
	}

	// finally we generate the generic response and return it
	res = models.CreateGenericResponse(true, http.StatusOK, "", classes)
	ctx.JSON(http.StatusOK, generateResponse(res))
	return
}

// getClass gets the class with the given ID from the DB
func (server *Server) getClass(ctx *gin.Context) {
	// creating request and response instances
	var req models.GetOneClassApiParam
	var res models.GenericResponse

	if err := ctx.ShouldBindUri(&req); err != nil {
		res = models.CreateGenericResponse(false, http.StatusBadRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, generateResponse(res))
		return
	}

	// then we create new GetOneClassParam
	arg := models.GetOneClassParam{
		Id: req.Id,
	}

	class, err := server.store.GetClass(ctx, arg)

	// if any error occurs we check for the error
	if err != nil {
		// if error is no rows error we return bad request
		if err == sql.ErrNoRows {
			res = models.CreateGenericResponse(false, http.StatusNotFound, "Couldn't find class with given ID", nil)
			ctx.JSON(http.StatusNotFound, generateResponse(res))
			return
		}
		res = models.CreateGenericResponse(false, http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, generateResponse(res))
		return
	}

	// finally we generate the generic response and return it
	res = models.CreateGenericResponse(true, http.StatusOK, "", class)
	ctx.JSON(http.StatusOK, generateResponse(res))
	return
}

// updateClass updates the class with given values
func (server *Server) updateClass(ctx *gin.Context) {
	// creating request and response instances
	var reqBody models.UpdateClassApiBodyParams
	var reqUri models.UpdateClassApiUriParam
	var res models.GenericResponse

	// if inputs are not valid we return status bad request with the error, we check both URI and json body
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		res = models.CreateGenericResponse(false, http.StatusBadRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, generateResponse(res))
		return
	}

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		res = models.CreateGenericResponse(false, http.StatusBadRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, generateResponse(res))
		return
	}

	// then we create new UpdateClassParams
	arg := models.UpdateClassParams{
		Id:        reqUri.Id,
		Name:      reqBody.Name,
		Professor: reqBody.Professor,
	}

	class, err := server.store.UpdateClass(ctx, arg)

	// if any error occurs we check for the error
	if err != nil {
		// if error is no rows error we return bad request
		if err == sql.ErrNoRows {
			res = models.CreateGenericResponse(false, http.StatusNotFound, "Couldn't find class with given ID", nil)
			ctx.JSON(http.StatusNotFound, generateResponse(res))
			return
		}
		res = models.CreateGenericResponse(false, http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, generateResponse(res))
		return
	}

	// finally we generate the generic response and return it
	res = models.CreateGenericResponse(true, http.StatusOK, "", class)
	ctx.JSON(http.StatusOK, generateResponse(res))
	return
}

// deleteClass deletes the class with given ID
func (server *Server) deleteClass(ctx *gin.Context) {
	// creating request and response instances
	var req models.DeleteClassApiParam
	var res models.GenericResponse

	// if inputs are not valid we return status bad request with the error, we check both URI and json body
	if err := ctx.ShouldBindUri(&req); err != nil {
		res = models.CreateGenericResponse(false, http.StatusBadRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, generateResponse(res))
		return
	}

	// then we create new DeleteOneClassParam
	arg := models.DeleteOneClassParam{
		Id: req.Id,
	}

	err := server.store.DeleteClass(ctx, arg)

	// if any error occurs we check for the error
	if err != nil {
		// if error is no rows error we return bad request
		if err == sql.ErrNoRows {
			res = models.CreateGenericResponse(false, http.StatusNotFound, "Couldn't find class with given ID", nil)
			ctx.JSON(http.StatusNotFound, generateResponse(res))
			return
		}
		res = models.CreateGenericResponse(false, http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, generateResponse(res))
		return
	}

	// finally we return no content with nil body
	ctx.JSON(http.StatusNoContent, nil)
	return
}
