package handlers

import (
	"net/http"

	"github.com/uacademy/e_commerce/api_gateway/models"
	ecom "github.com/uacademy/e_commerce/api_gateway/proto-gen/e_commerce"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware ...
func (h Handler) AuthMiddleware(userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		hasAccessResponse, err := h.GrpcClients.Auth.HasAccess(c.Request.Context(), &ecom.TokenRequest{
			Token: token,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.JSONError{
				Error: err.Error(),
			})
			c.Abort()
			return
		}

		if !hasAccessResponse.HasAccess {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		if userType != "*" {
			if hasAccessResponse.User.UserType != userType {
				c.JSON(http.StatusUnauthorized, "Permission Denied")
				c.Abort()
			}
		}

		c.Next()
		//
	}
}

// Login godoc
// @Summary     Login
// @Description Login
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       login body     models.LoginModel true "Login body"
// @Success     201   {object} models.JSONResult{data=models.TokenResponse}
// @Failure     400   {object} models.JSONError
// @Router      /v1/login [post]
func (h Handler) Login(c *gin.Context) {
	var body models.LoginModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{Error: err.Error()})
		return
	}

	// TODO - validation should be here
	tokenResponse, err := h.GrpcClients.Auth.Login(c.Request.Context(), &ecom.LoginRequest{
		Username: body.Username,
		Password: body.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "OK",
		Data:    tokenResponse,
	})
}
