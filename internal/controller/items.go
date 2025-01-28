package controller

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"webproj/internal/service"
)

type ItemHandler struct {
	ctx      context.Context
	services *service.Services
}

func NewItemHandler(ctx context.Context, services *service.Services) *ItemHandler {
	return &ItemHandler{
		ctx:      ctx,
		services: services,
	}
}

func (h *ItemHandler) NewItem(ctx echo.Context) error {
	var itemValue service.ItemValues
	err := ctx.Bind(&itemValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("could not decode input data"))
	}

	err = ctx.Validate(&itemValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("could not validate input data"))
	}

	err = h.services.Items.NewItem(ctx.Request().Context(), &itemValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusCreated, "Item successfully created")
}

func (h *ItemHandler) Delete(ctx echo.Context) error {
	itemId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid id",
		})
	}
	err = h.services.Items.Delete(ctx.Request().Context(), itemId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, "Item successfully deleted")
}

func (h *ItemHandler) GetAll(ctx echo.Context) error {
	i, err := h.services.Items.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get items")
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"items": i,
	})
}
func (h *ItemHandler) Get(ctx echo.Context) error {
	itemId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, echo.Map{
			"Message": "invalid id",
		})
	}
	i, err := h.services.Items.Get(ctx.Request().Context(), itemId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get item by ID")
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"item": i,
	})
}

func (h *ItemHandler) Update(ctx echo.Context) error {
	return nil
}
