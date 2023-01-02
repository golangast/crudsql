package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"

	. "sql/db"
	. "sql/generate"
)

//go:embed templates
var tmplMainGo embed.FS

func main() {
	//get all the files from the embedding
	files, err := getAllFilenames(&tmplMainGo)
	if err != nil {
		fmt.Print(err)
	}
	//generate the db
	Gendatabase("db")
	Createtable()
	//start the server
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseFS(tmplMainGo, files...)),
	}
	e.Renderer = t
	e.GET("/form", GetForm)
	e.GET("/getdb", GetdbNames)
	e.GET("/putform", Putform)
	e.GET("/deleteform", Deleteform)

	e.POST("/post", PostForm)
	e.POST("/put", Putdb)
	e.POST("/delete", Deletedb)

	e.Logger.Fatal(e.Start(":8080"))
}

// show form
func GetForm(c echo.Context) error {

	return c.Render(http.StatusOK, "form.html", map[string]interface{}{})
}

// post form data
func PostForm(c echo.Context) error {

	fname := c.FormValue("fname")
	lname := c.FormValue("lname")

	DbConnection()
	//save it in the db
	Addsavedata(fname, lname)

	return c.Render(http.StatusOK, "show.html", map[string]interface{}{
		"fname": fname,
		"lname": lname,
	})
}

// grab data
func GetdbNames(c echo.Context) error {

	names := Getdb()

	return c.Render(http.StatusOK, "showdb.html", map[string]interface{}{
		"names": names,
	})
}

// show updated database data
func Putform(c echo.Context) error {

	names := Getdb()
	return c.Render(http.StatusOK, "updatedata.html", map[string]interface{}{
		"names": names,
	})
}

// update the table
func Putdb(c echo.Context) error {

	fname := c.FormValue("fname")
	lname := c.FormValue("lname")

	UpdateNames(fname, lname)
	names := Getdb()

	return c.Render(http.StatusOK, "showdb.html", map[string]interface{}{
		"names": names,
	})

}

// delete name
func Deletedb(c echo.Context) error {

	fname := c.FormValue("fname")
	lname := c.FormValue("lname")

	DeleteNames(fname, lname)

	names := Getdb()

	return c.Render(http.StatusOK, "showdb.html", map[string]interface{}{
		"names": names,
	})

}
func Deleteform(c echo.Context) error {

	return c.Render(http.StatusOK, "deleteform.html", map[string]interface{}{})

}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// walk through embedded file
func getAllFilenames(efs *embed.FS) (files []string, err error) {
	if err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
}
