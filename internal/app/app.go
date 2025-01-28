package app

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
	"webproj/internal/config"
	"webproj/internal/controller"
	"webproj/internal/lib"
	mineMW "webproj/internal/middleware"
	"webproj/internal/repository"
	pg "webproj/internal/repository/PG"
	"webproj/internal/service"
)

func Run() {
	ctx := context.Background()

	cfg := config.Get()

	slog.Info("Connection to DB")
	db, err := pg.Init()
	if err != nil {
		slog.Error("Connection failed to DB", "error", err)
	}
	defer db.Close()

	repos := repository.NewRepositories(db)

	services := service.NewServices(service.Deps{
		Repos:          repos,
		AccessTokenTTL: cfg.AccessTokenTTL,
	})

	//handlers
	userHandler := controller.NewUserHandler(ctx, services)
	itemHandler := controller.NewItemHandler(ctx, services)
	orderHandler := controller.NewOrderHandler(ctx, services)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path}` + "\n",
	}))
	e.Validator = lib.NewCustomValidator()

	//routes
	e.POST("/sign-up", userHandler.SignUp)
	e.POST("/sign-in", userHandler.SignIn)

	//user routes
	UserRoute := e.Group("/users")
	UserRoute.GET("/:id", userHandler.GetById)
	UserRoute.DELETE("/:id", userHandler.Delete)

	//item routes
	ItemRoute := e.Group("/items")
	ItemRoute.POST("", itemHandler.NewItem)
	ItemRoute.GET("", itemHandler.GetAll)
	ItemRoute.GET("/:id", itemHandler.Get)
	ItemRoute.DELETE("/:id", itemHandler.Delete)

	//order routes
	OrderRoute := e.Group("/orders")
	OrderRoute.GET("", orderHandler.GetAll, mineMW.AuthMW)
	OrderRoute.GET("/:id", orderHandler.GetById, mineMW.AuthMW)
	OrderRoute.POST("", orderHandler.MakeOrder, mineMW.AuthMW)

	// Start server
	err = e.StartServer(&http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  cfg.ReadTimeOut,
		WriteTimeout: cfg.WriteTimeOut,
	})
	if err != nil {
		slog.Error("failed to start the server", "error", err)
	}

}
