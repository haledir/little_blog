package handlers

import (
	"net/http"

	"github.com/haledir/little_blog/services"
	"github.com/labstack/echo/v4"
)

type EditorHandler struct {
	ArticleService *services.ArticleService
}

func (h *EditorHandler) HandleEditor(c echo.Context) error {
	return c.Render(http.StatusOK, "editor.html", map[string]interface{}{
		"Title": "Create a new Article",
	})
}

func (h *EditorHandler) HandleSaveArticle(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")

	err := h.ArticleService.CreateArticle(title, content)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save article")
	}
	return c.Redirect(http.StatusSeeOther, "/dashboard")
}
