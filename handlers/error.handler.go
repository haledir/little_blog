package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok && he.Code == http.StatusNotFound {
		c.Render(http.StatusNotFound, "404.html", nil)
		return
	}

	c.Echo().DefaultHTTPErrorHandler(err, c)
}
