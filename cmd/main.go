package main

import (
	"html/template"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/haledir/little_blog/db"
	"github.com/haledir/little_blog/handlers"
	"github.com/haledir/little_blog/services"
)

func main() {

	database, err := db.InitDB("./local.db")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer database.Close()

	articleService := &services.ArticleService{DB: database}
	authService := &services.AuthService{DB: database}

	articleHander := &handlers.ArticleHandler{ArticleService: articleService}
	authHandler := &handlers.AuthHandler{AuthService: authService}
	welcomeHandler := &handlers.WelcomeHander{}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &handlers.TemplateRenderer{
		Templates: template.Must(template.New("").Funcs(template.FuncMap{
			"safeHTML": handlers.SafeHTML,
		}).ParseGlob("public/views/*.html")),
	}

	e.Renderer = t

	e.GET("/", articleHander.HandleIndex)
	e.GET("/article/:id", articleHander.HandleArticle)
	e.GET("/login", authHandler.HandleLogin)
	e.POST("/login", authHandler.HandleLoginSubmit)
	e.GET("/welcome", welcomeHandler.HandleWelcome)

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":50259"))
}
