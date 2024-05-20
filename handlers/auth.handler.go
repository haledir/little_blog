package handlers

import (
	"github.com/haledir/little_blog/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func (h *AuthHandler) HandleLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"Title": "Login",
	})
}

func (h *AuthHandler) HandleLoginSubmit(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	_, err := h.AuthService.AuthenticateUser(username, password)
	if err != nil {
		return c.Render(http.StatusUnauthorized, "login.html", map[string]interface{}{
			"Title":   "Login",
			"Message": "Invalid username and/or password",
		})
	}
	return c.Redirect(http.StatusSeeOther, "/welcome?username="+username)
}
