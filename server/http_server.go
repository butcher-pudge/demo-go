package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-learning-demo/auth"
	authController "go-learning-demo/auth/controller"
	"go-learning-demo/auth/middleware"
	authMongo "go-learning-demo/auth/repository/mongo"
	authService "go-learning-demo/auth/service"
	"go-learning-demo/bookmark"
	bookmarkController "go-learning-demo/bookmark/controller"
	bookmarkMongo "go-learning-demo/bookmark/repository/mongo"
	bookmarkService "go-learning-demo/bookmark/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer         *http.Server
	authController     *authController.AuthController
	authService        authService.IAuthService
	bookmarkController *bookmarkController.BookmarkController
}

func NewApp() *App {
	db := initDB()

	userRepository := authMongo.NewUserRepository(db, viper.GetString("mongo.user_collection"))
	bookmarkRepository := bookmarkMongo.NewBookmarkRepository(db, viper.GetString("mongo.bookmark_collection"))

	service := authService.NewAuthService(userRepository)
	return &App{
		authController:     authController.NewAuthController(service),
		authService:        service,
		bookmarkController: bookmarkController.NewBookmarkController(bookmarkService.NewBookmarkService(bookmarkRepository)),
	}
}

func (app *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger())

	auth.RegisterHttpEndpoints(router, app.authController)

	//API endpoints
	authMiddleware := middleware.NewAuthMiddleware(app.authService)
	apiRouterGroup := router.Group("/api", authMiddleware)

	bookmark.RegisterHttpEndpoints(apiRouterGroup, app.bookmarkController)

	// HTTP Server
	app.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return app.httpServer.Shutdown(ctx)
}
