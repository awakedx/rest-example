package controller

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"webproj/internal/service"
)

type OrderHandler struct {
	ctx      context.Context
	servives *service.Services
}

func NewOrderHandler(ctx context.Context, service *service.Services) *OrderHandler {
	return &OrderHandler{
		ctx:      ctx,
		servives: service,
	}
}

func (h *OrderHandler) MakeOrder(ctx echo.Context) error {
	var input service.InputOrder
	err := ctx.Bind(&input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("could not decode input data"))
	}

	input.UserId, err = service.GetUserIdClaims(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	orderId, err := h.servives.Orders.MakeOrder(ctx.Request().Context(), &input)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	} else {
		return ctx.JSON(http.StatusCreated, echo.Map{
			"OrderId": orderId,
		})
	}
}
func (h *OrderHandler) GetAll(ctx echo.Context) error {
	userId, err := service.GetUserIdClaims(ctx)
	if err != nil {
		return err
	}
	orders, err := h.servives.Orders.GetAll(ctx.Request().Context(), userId)
	return ctx.JSON(http.StatusOK, echo.Map{
		"Orders": orders,
	})
}
func (h *OrderHandler) GetById(ctx echo.Context) error {
	orderId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid id",
		})
	}
	userId, err := service.GetUserIdClaims(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	order, err := h.servives.Orders.GetById(ctx.Request().Context(), orderId, userId)
	if err != nil {
		if err.Error() == "invalid order id" {
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"Error": err.Error(),
			})
		}
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"Order": order,
	})
}
