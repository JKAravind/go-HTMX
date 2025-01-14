package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseFiles("views/index.html")),
	}
}

type UserCredential struct {
	Name string
	Mail string
}

type UserCredentials = []UserCredential

type Data struct {
	UserCredentials UserCredentials
}

func addData() Data {
	return Data{
		UserCredentials: []UserCredential{
			addCredential("JohnDoe", "jd@gmail.com"),
			addCredential("JaneDoe", "ja@gmail.com"),
		},
	}
}

func addCredential(name string, mail string) UserCredential {
	return UserCredential{
		Name: name,
		Mail: mail,
	}
}

func main() {
	e := echo.New()
	e.Renderer = newTemplate()
	data := addData()

	e.GET("/", func(c echo.Context) error {

		return c.Render(200, "index", data)
	})

	e.Start(":8084")

}
