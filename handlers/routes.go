package handlers

import (
	"github.com/haledir/little_blog/middleware"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, authHandler *AuthHandler, articleHandler *ArticleHandler, welcomeHandler *WelcomeHandler) {
	e.GET("/", articleHandler.HandleIndex)
	e.GET("/article/:id", articleHandler.HandleArticle)

	e.GET("/login", authHandler.HandleLogin)
	e.POST("/login", authHandler.HandleLoginSubmit)

	protectedGroup := e.Group("", middleware.AuthMiddleware)
	protectedGroup.GET("/welcome", welcomeHandler.HandleWelcome)
	protectedGroup.GET("/logout", authHandler.HandleLogout)
}
