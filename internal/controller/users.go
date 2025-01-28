package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webproj/internal/service"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	ctx      context.Context
	servives *service.Services
}

func NewUserHandler(ctx context.Context, services *service.Services) *UserHandler {
	return &UserHandler{
		ctx:      ctx,
		servives: services,
	}
}

func (h *UserHandler) SignUp(ctx echo.Context) error {
	var signUpInput service.SignUpInput
	err := ctx.Bind(&signUpInput)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("could not decode user data"))
	}
	err = ctx.Validate(&signUpInput)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	err = h.servives.Users.SignUp(ctx.Request().Context(), &signUpInput)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusCreated, "user successfully created")
}

func (h *UserHandler) SignIn(ctx echo.Context) error {
	var signInInput service.SignInInput
	err := ctx.Bind(&signInInput)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("could not decode input data"))
	}
	err = ctx.Validate(&signInInput)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	accToken, err := h.servives.Users.SignIn(ctx.Request().Context(), &signInInput)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"Error": err.Error(),
		})
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = accToken
	cookie.Expires = time.Now().Add(h.servives.Users.GetAccTokenTTL())
	cookie.HttpOnly = true
	ctx.SetCookie(cookie)
	return ctx.JSON(http.StatusOK, "Successfully authorized")
}

func (h *UserHandler) Delete(ctx echo.Context) error {
	userId, err := uuid.Parse(ctx.Param("id"))
	fmt.Println(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = h.servives.Users.DeleteUser(ctx.Request().Context(), userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, "user deleted")
}

func (h *UserHandler) GetById(ctx echo.Context) error {
	userId, err := uuid.Parse(ctx.Param("id"))
	fmt.Println(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	user, err := h.servives.Users.GetById(ctx.Request().Context(), userId)
	return ctx.JSON(http.StatusOK, echo.Map{
		"User": user,
	})
}
