package handlers

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"

	"github.com/uacademy/e_commerce/api_gateway/models"
	ecom "github.com/uacademy/e_commerce/api_gateway/proto-gen/e_commerce"
)

// CreateUser godoc
// @Summary     Create user
// @Description create new user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       user body     models.CreateUserModel true "User body"
// @Success     201  {object} models.JSONResult{data=models.User}
// @Failure     400  {object} models.JSONError
// @Failure     500  {object} models.JSONError
// @Router      /v1/user [post]
func (h Handler) CreateUser(c *gin.Context) {
	var body models.CreateUserModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{Error: err.Error()})
		return
	}

	user, err := h.GrpcClients.Auth.CreateUser(c.Request.Context(), &ecom.CreateUserRequest{
		Username: body.Username,
		Password: body.Password,
		UserType: body.User_type,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "OK",
		Data:    user,
	})
}

// GetUser godoc
// @Summary     Get user
// @Description get user by ID
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       id  path     string true "User ID"
// @Success     200 {object} models.JSONResult{data=models.User}
// @Failure     404 {object} models.JSONError
// @Router      /v1/user/{id} [get]
func (h Handler) GetUserById(c *gin.Context) {
	id := c.Param("id")

	user, err := h.GrpcClients.Auth.GetUserByID(c.Request.Context(), &ecom.GetUserByIDRequest{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    user,
	})
}

// ListUsers godoc
// @Summary     List users
// @Description get users
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       offset query    string false "0"
// @Param       limit  query    string false "10"
// @Param       search query    string false "smth"
// @Success     200    {object} models.JSONResult{data=[]models.User}
// @Failure     400    {object} models.JSONError
// @Failure     500    {object} models.JSONError
// @Router      /v1/user [get]
func (h Handler) GetUserList(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")
	searchStr := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	userList, err := h.GrpcClients.Auth.GetUserList(c.Request.Context(), &ecom.GetUserListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
		Search: searchStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    userList,
	})
}

// UpdateUser godoc
// @Summary     Update user
// @Description update user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       user body     models.UpdateUserModel true "User body"
// @Success     200  {object} models.JSONResult{data=models.User}
// @Failure     400  {object} models.JSONError
// @Failure     404  {object} models.JSONError
// @Router      /v1/user [put]
func (h Handler) UpdateUser(c *gin.Context) {
	var body models.UpdateUserModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{Error: err.Error()})
		return
	}

	user, err := h.GrpcClients.Auth.UpdateUser(c.Request.Context(), &ecom.UpdateUserRequest{
		Id:       body.Id,
		Password: body.Password,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    user,
	})
}

// DeleteUser godoc
// @Summary     Delete user
// @Description delete user by ID
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       id  path     string true "User ID"
// @Success     200 {object} models.JSONResult{data=models.User}
// @Failure     400 {object} models.JSONError
// @Router      /v1/user/{id} [delete]
func (h Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.GrpcClients.Auth.DeleteUser(c.Request.Context(), &ecom.DeleteUserRequest{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    user,
	})
}
