package middleware

import (
	"github.com/haledir/little_blog/sessions"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, err := sessions.GetSession(c)
		if err != nil || session.Values["username"] == nil {
			return c.Redirect(http.StatusSeeOther, "/")
		}
		return next(c)
	}
}
