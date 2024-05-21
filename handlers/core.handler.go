package handlers

import (
	"errors"
	"html/template"
	"io"

	"github.com/haledir/little_blog/sessions"
	"github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	Templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	session, err := sessions.GetSession(c)
	if err != nil {
		return err
	}

	if data == nil {
		data = map[string]interface{}{}
	}

	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return errors.New("invalid data map")
	}

	dataMap["IsLoggedIn"] = session.Values["username"] != nil
	dataMap["Username"] = session.Values["username"]

	return t.Templates.ExecuteTemplate(w, name, data)
}

// SafeHTML is a custom template function to mark a string as safe HTML
func SafeHTML(s string) template.HTML {
	return template.HTML(s)
}
