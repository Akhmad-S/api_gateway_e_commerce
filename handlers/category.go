package handlers

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"

	"github.com/uacademy/e_commerce/api_gateway/models"
	ecom "github.com/uacademy/e_commerce/api_gateway/proto-gen/e_commerce"
)

// CreateCategory godoc
// @Summary     Create category
// @Description create new category
// @Tags        categories
// @Accept      json
// @Produce     json
// @Param       category      body     models.CreateCategoryModel true  "Category body"
// @Param       Authorization header   string                     false "Authorization"
// @Success     201           {object} models.JSONResult{data=models.Category}
// @Failure     400           {object} models.JSONError
// @Failure     500           {object} models.JSONError
// @Router      /v1/category [post]
func (h Handler) CreateCategory(c *gin.Context) {
	var body models.CreateCategoryModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{Error: err.Error()})
		return
	}

	category, err := h.GrpcClients.Category.CreateCategory(c.Request.Context(), &ecom.CreateCategoryRequest{
		CategoryTitle: body.CategoryTitle,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "OK",
		Data:    category,
	})
}

// GetCategory godoc
// @Summary     Get category
// @Description get category by ID
// @Tags        categories
// @Accept      json
// @Produce     json
// @Param       id            path     string true  "Category ID"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResult{data=models.Category}
// @Failure     404           {object} models.JSONError
// @Router      /v1/category/{id} [get]
func (h Handler) GetCategoryById(c *gin.Context) {
	id := c.Param("id")

	category, err := h.GrpcClients.Category.GetCategoryById(c.Request.Context(), &ecom.GetCategoryByIdRequest{
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
		Data:    category,
	})
}

// ListCategorys godoc
// @Summary     List categories
// @Description get categories
// @Tags        categories
// @Accept      json
// @Produce     json
// @Param       offset        query    string false "0"
// @Param       limit         query    string false "10"
// @Param       search        query    string false "smth"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResult{data=[]models.Category}
// @Failure     400           {object} models.JSONError
// @Failure     500           {object} models.JSONError
// @Router      /v1/category [get]
func (h Handler) GetCategoryList(c *gin.Context) {
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

	categoryList, err := h.GrpcClients.Category.GetCategoryList(c.Request.Context(), &ecom.GetCategoryListRequest{
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
		Data:    categoryList,
	})
}

// UpdateCategory godoc
// @Summary     Update category
// @Description update category
// @Tags        categories
// @Accept      json
// @Produce     json
// @Param       category      body     models.UpdateCategoryModel true  "Category body"
// @Param       Authorization header   string                     false "Authorization"
// @Success     200           {object} models.JSONResult{data=models.Category}
// @Failure     400           {object} models.JSONError
// @Failure     404           {object} models.JSONError
// @Router      /v1/category [put]
func (h Handler) UpdateCategory(c *gin.Context) {
	var body models.UpdateCategoryModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{Error: err.Error()})
		return
	}

	category, err := h.GrpcClients.Category.UpdateCategory(c.Request.Context(), &ecom.UpdateCategoryRequest{
		Id:            body.Id,
		CategoryTitle: body.CategoryTitle,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    category,
	})
}

// DeleteCategory godoc
// @Summary     Delete category
// @Description delete category by ID
// @Tags        categories
// @Accept      json
// @Produce     json
// @Param       id            path     string true  "Category ID"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResult{data=models.Category}
// @Failure     400           {object} models.JSONError
// @Router      /v1/category/{id} [delete]
func (h Handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	category, err := h.GrpcClients.Category.DeleteCategory(c.Request.Context(), &ecom.DeleteCategoryRequest{
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
		Data:    category,
	})
}
