package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func saveUser(c echo.Context) error {
	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "id:"+id+", name:"+name+", email:"+email)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "id:"+id)
}

func saveMultipart(c echo.Context) error {
	name := c.FormValue("name")

	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}
	fmt.Println(avatar.Filename)

	src, err := avatar.Open()
	if err != nil {
		return err
	}
	//close file after function return
	defer src.Close()

	dst, err := os.Create("static/" + avatar.Filename)
	if err != nil {
		return err
	}
	//close file after function return
	defer dst.Close()

	//copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}

func main() {
	e := echo.New()

	//global middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//group middleware
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(
		func(username, password string, c echo.Context) (bool, error) {
			if username == "admin" && password == "123456" {
				return true, nil
			}
			return false, nil
		}))

	//route middleware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}

	//lambda funct
	e.GET("/", func(c echo.Context) error {
		return c.String(
			http.StatusOK,
			"Hello, World!")
	})
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser, track)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.POST("/save/avatar", saveMultipart)
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":10020"))
}
