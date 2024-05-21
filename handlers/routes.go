package handlers

import (
	"github.com/haledir/little_blog/middleware"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, authHandler *AuthHandler, articleHandler *ArticleHandler, welcomeHandler *WelcomeHandler, editorHandler *EditorHandler) {
	e.GET("/", articleHandler.HandleIndex)
	e.GET("/article/:id", articleHandler.HandleArticle)

	e.GET("/login", authHandler.HandleLogin)
	e.POST("/login", authHandler.HandleLoginSubmit)

	protectedGroup := e.Group("", middleware.AuthMiddleware)
	protectedGroup.GET("/welcome", welcomeHandler.HandleWelcome)
	protectedGroup.GET("/dashboard", welcomeHandler.HandleDashboard)
	protectedGroup.GET("/editor", editorHandler.HandleEditor)
	protectedGroup.POST("/save-article", editorHandler.HandleSaveArticle)
	protectedGroup.GET("/logout", authHandler.HandleLogout)
}
