package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type WelcomeHander struct{}

func (h *WelcomeHander) HandleWelcome(c echo.Context) error {
	username := c.QueryParam("username")
	return c.Render(http.StatusOK, "welcome.html", map[string]interface{}{
		"Title":    "Welcome",
		"Username": username,
	})
}
