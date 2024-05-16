package handlers

import (
	"net/http"
	"strconv"

	"github.com/haledir/little_blog/services"
	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	ArticleService *services.ArticleService
}

func (h *ArticleHandler) HandleIndex(c echo.Context) error {
	articles, err := h.ArticleService.GetAllArticles()
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"Title":    "Blog Home",
		"Articles": articles,
	}
	return c.Render(http.StatusOK, "index.html", data)
}

func (h *ArticleHandler) HandleArticle(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid article ID")
	}

	article, err := h.ArticleService.GetArticleById(id)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	data := map[string]interface{}{
		"Title":   article.Title,
		"Content": article.Content,
	}
	return c.Render(http.StatusOK, "article.html", data)
}
