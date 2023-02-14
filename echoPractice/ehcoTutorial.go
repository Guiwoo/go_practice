package echoPractice

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

type Product struct {
	ID    string `query:"id""`
	Name  string `query:"name"`
	Price int    `query:"price"`
}

func (p *Product) String() string {
	return fmt.Sprintf("ID is :%s , Name is :%s, Price is :%d", p.ID, p.Name, p.Price)
}

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

type UserDTO struct {
	Name    string
	Email   string
	IsAdmin bool
}

/**
Path parameters
Query parameters (only for GET/DELETE methods)
Request body

이 순서로 바인딩 되서 가장 마지막 꺼가 바인딩됨 만약 path 로 보내는데 body 가 비었다 ? 그건 empty parameter 임
*/

func Ex03() {
	logic := func(u UserDTO) {
		fmt.Println("User logic by UserDTO", u)
	}

	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		u := new(User)
		err := c.Bind(u)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		user := UserDTO{
			Name:    u.Name,
			Email:   u.Email,
			IsAdmin: false,
		}
		logic(user)
		return c.JSON(http.StatusOK, u)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func Ex02x() {
	e := echo.New()

	var product Product
	e.GET("/", func(c echo.Context) error {
		err := c.Bind(&product)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusOK, product.String())
	})
	e.Logger.Fatal(e.Start(":8080"))
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
