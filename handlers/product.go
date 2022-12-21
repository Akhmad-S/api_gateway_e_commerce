package handlers

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"

	"github.com/uacademy/e_commerce/api_gateway/models"
	ecom "github.com/uacademy/e_commerce/api_gateway/proto-gen/e_commerce"
)

// CreateProduct godoc
// @Summary     Create product
// @Description create new product
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       product       body     models.CreateProductModel true  "Product body"
// @Param       Authorization header   string                    false "Authorization"
// @Success     201           {object} models.JSONResult{data=models.Product}
// @Failure     400           {object} models.JSONError
// @Failure     500           {object} models.JSONError
// @Router      /v1/product [post]
func (h Handler) CreateProduct(c *gin.Context) {
	var body models.CreateProductModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{Error: err.Error()})
		return
	}

	product, err := h.GrpcClients.Product.CreateProduct(c.Request.Context(), &ecom.CreateProductRequest{
		CategoryId: body.CategoryId,
		Title:      body.Title,
		Descrip:    body.Descrip,
		Price:      body.Price,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "OK",
		Data:    product,
	})
}

// GetProduct godoc
// @Summary     Get product
// @Description get product by ID
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       id            path     string true  "Product ID"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResult{data=models.Product}
// @Failure     404           {object} models.JSONError
// @Router      /v1/product/{id} [get]
func (h Handler) GetProductById(c *gin.Context) {
	id := c.Param("id")

	product, err := h.GrpcClients.Product.GetProductById(c.Request.Context(), &ecom.GetProductByIdRequest{
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
		Data:    product,
	})
}

// ListProducts godoc
// @Summary     List products
// @Description get products
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       offset        query    string false "0"
// @Param       limit         query    string false "10"
// @Param       search        query    string false "smth"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResult{data=[]models.Product}
// @Failure     400           {object} models.JSONError
// @Failure     500           {object} models.JSONError
// @Router      /v1/product [get]
func (h Handler) GetProductList(c *gin.Context) {
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

	productList, err := h.GrpcClients.Product.GetProductList(c.Request.Context(), &ecom.GetProductListRequest{
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
		Data:    productList,
	})
}

// UpdateProduct godoc
// @Summary     Update product
// @Description update product
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       product       body     models.UpdateProductModel true  "Product body"
// @Param       Authorization header   string                    false "Authorization"
// @Success     200           {object} models.JSONResult{data=models.Product}
// @Failure     400           {object} models.JSONError
// @Failure     404           {object} models.JSONError
// @Router      /v1/product [put]
func (h Handler) UpdateProduct(c *gin.Context) {
	var body models.UpdateProductModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{Error: err.Error()})
		return
	}

	product, err := h.GrpcClients.Product.UpdateProduct(c.Request.Context(), &ecom.UpdateProductRequest{
		Id:    body.Id,
		Title: body.Title,
		Price: body.Price,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    product,
	})
}

// DeleteProduct godoc
// @Summary     Delete product
// @Description delete product by ID
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       id            path     string true  "Product ID"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResult{data=models.Product}
// @Failure     400           {object} models.JSONError
// @Router      /v1/product/{id} [delete]
func (h Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	product, err := h.GrpcClients.Product.DeleteProduct(c.Request.Context(), &ecom.DeleteProductRequest{
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
		Data:    product,
	})
}
