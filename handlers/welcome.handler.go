package handlers

import (
	"github.com/haledir/little_blog/sessions"
	"github.com/labstack/echo/v4"
	"net/http"
)

type WelcomeHandler struct{}

func (h *WelcomeHandler) HandleWelcome(c echo.Context) error {
	session, err := sessions.GetSession(c)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/")
	}

	username := session.Values["username"]
	if username == nil {
		return c.Redirect(http.StatusSeeOther, "/")
	}

	return c.Render(http.StatusOK, "welcome.html", map[string]interface{}{
		"Title":    "Welcome",
		"Username": username,
	})
}
