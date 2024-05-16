package handlers

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	Templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

// SafeHTML is a custom template function to mark a string as safe HTML
func SafeHTML(s string) template.HTML {
	return template.HTML(s)
}
