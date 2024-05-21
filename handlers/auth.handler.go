package handlers

import (
	"net/http"

	"github.com/haledir/little_blog/services"
	"github.com/haledir/little_blog/sessions"
	"github.com/labstack/echo/v4"
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

	user, err := h.AuthService.AuthenticateUser(username, password)
	if err != nil {
		return c.Render(http.StatusUnauthorized, "login.html", map[string]interface{}{
			"Title":   "Login",
			"Message": "Invalid username and/or password",
		})
	}

	err = sessions.SetSessionValue(c, "username", user.Username)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/welcome?username="+username)
}

func (h *AuthHandler) HandleLogout(c echo.Context) error {
	session, err := sessions.GetSession(c)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusSeeOther, "/")
}
