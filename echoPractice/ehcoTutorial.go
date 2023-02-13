package echoPractice

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func Ex02x() {

}

func Ex01() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {

		cookie := new(http.Cookie)
		cookie.Name = "abc"
		cookie.Value = "abcdefgRinbowAttack"
		cookie.Expires = time.Now().Add(time.Hour)

		c.SetCookie(cookie)
		return c.String(http.StatusOK, "Hello World")
	})
	e.POST("/test", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		return c.String(http.StatusOK, fmt.Sprintf("Save name : %s , email : %s", name, email))
	})
	e.GET("/test", func(c echo.Context) error {
		abc, err := c.Cookie("abc")
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, fmt.Sprint("Cookie abc is ", abc))
	})
	e.Logger.Fatal(e.Start(":8080"))
}
