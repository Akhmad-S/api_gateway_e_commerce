package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/uacademy/e_commerce/api_gateway/models"
	ecom "github.com/uacademy/e_commerce/api_gateway/proto-gen/e_commerce"
)

// CreateOrder godoc
// @Summary     Create order
// @Description create new order
// @Tags        orders
// @Accept      json
// @Produce     json
// @Param       order         body     models.CreateOrderModel true  "Order body"
// @Param       Authorization header   string                  false "Authorization"
// @Success     201           {object} models.JSONResult{data=models.Order}
// @Failure     400           {object} models.JSONError
// @Failure     500           {object} models.JSONError
// @Router      /v1/order [post]
func (h Handler) CreateOrder(c *gin.Context) {
	var body models.CreateOrderModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{Error: err.Error()})
		return
	}

	_, err := h.GrpcClients.Product.GetProductById(c.Request.Context(), &ecom.GetProductByIdRequest{
		Id: body.Product_id,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	order, err := h.GrpcClients.Order.CreateOrder(c.Request.Context(), &ecom.CreateOrderRequest{
		ProductId:   body.Product_id,
		Quantity:    body.Quantity,
		UserName:    body.User_name,
		UserAddress: body.User_address,
		UserPhone:   body.User_phone,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "OK",
		Data:    order,
	})
}

// ListOrders godoc
// @Summary     List orders
// @Description get orders
// @Tags        orders
// @Accept      json
// @Produce     json
// @Param       offset        query    string false "0"
// @Param       limit         query    string false "10"
// @Param       search        query    string false "smth"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResult{data=[]models.Order}
// @Failure     400           {object} models.JSONError
// @Failure     500           {object} models.JSONError
// @Router      /v1/order [get]
func (h Handler) GetOrderList(c *gin.Context) {
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

	orderList, err := h.GrpcClients.Order.GetOrderList(c.Request.Context(), &ecom.GetOrderListRequest{
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
		Data:    orderList,
	})
}

// GetOrder godoc
// @Summary     Get order
// @Description get order by ID
// @Tags        orders
// @Accept      json
// @Produce     json
// @Param       id            path     string true  "Order ID"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResult{data=models.PackedOrderModel}
// @Failure     404           {object} models.JSONError
// @Router      /v1/order/{id} [get]
func (h Handler) GetOrderById(c *gin.Context) {
	id := c.Param("id")

	order, err := h.GrpcClients.Order.GetOrderById(c.Request.Context(), &ecom.GetOrderByIdRequest{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	product, err := h.GrpcClients.Product.GetProductById(c.Request.Context(), &ecom.GetProductByIdRequest{
		Id: order.Product.Id,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	order.Product = &ecom.GetOrderByIdResponse_Product{
		Id:    product.Id,
		Title: product.Title,
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    order,
	})
}
