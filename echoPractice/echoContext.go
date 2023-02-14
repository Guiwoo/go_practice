package echoPractice

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"html/template"
	"io"
	"net/http"
	"sync"
	"time"
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo() {
	fmt.Println("FOoo")
}
func (c *CustomContext) Bar() {
	fmt.Println("Bar")
}

func Ex04() {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			cc.Foo()
			cc.Bar()
			return next(cc)
		}
	})

	e.GET("/", func(c echo.Context) error {
		cc := c.(*CustomContext)
		cc.Foo()
		cc.Bar()
		return c.String(http.StatusOK, " Done")
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "whatimade"
	cookie.Value = "sbwlstm"
	cookie.Expires = time.Now().Add(time.Minute)
	c.SetCookie(cookie)

	return c.String(http.StatusOK, "Cookie set is done")
}

func readCookie(c echo.Context) error {
	cookie, err := c.Cookie("whatimade")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, fmt.Sprintln(cookie.Name, cookie.Value))
}

func readAllCookies(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Println(cookie.Name, cookie.Value)
	}
	return c.String(http.StatusOK, "Look at the console")
}

func Ex05() {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
		}
	})

	e.GET("/", func(c echo.Context) error {
		return writeCookie(c)
	})

	e.GET("/2", func(c echo.Context) error {
		fmt.Println(c.Cookie("whatimade"))
		return readCookie(c)
	})
	e.GET("/3", func(c echo.Context) error {
		return readAllCookies(c)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func Ex06RetrieveData() {
	e := echo.New()

	//Form Data
	e.GET("/form", func(c echo.Context) error {
		name := c.FormValue("name")
		return c.String(http.StatusOK, name)
	})

	e.GET("/queryParam", func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.String(http.StatusOK, name)
	})

	e.GET("/pathParam/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, name)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

type (
	User2 struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Ex07Validate() {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/user2", func(c echo.Context) error {
		u := new(User2)
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := c.Validate(u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, u)
	})
	e.Logger.Fatal(e.Start(":8080"))
}

type Template struct {
	template *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func Ex08Response() {
	e := echo.New()

	render := &Template{
		template: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = render

	e.GET("/string", func(c echo.Context) error {
		return c.String(http.StatusOK, "String Return")
	})

	e.GET("/htmlref", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<div><h1>Hello World</h1><button>Join us</button></div>")
	})

	e.GET("/htmlRefAsFile", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello.html", "Data")
	})

	//RenderTemplate

	e.Logger.Fatal(e.Start(":8080"))
}

type (
	Stats struct {
		Uptime time.Time
		mutex  sync.RWMutex
	}
)

func NewStats() *Stats {
	return &Stats{
		Uptime: time.Now(),
	}
}

func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		log := c.Logger()
		log.Info("Before Access Service Running this middleware")
		return nil
	}
}

func ExRouter() {
	e := echo.New()

	e.Logger.SetLevel(log.DEBUG)

	g := e.Group("/user")
	s := NewStats()
	g.Use(s.Process)

	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "/user")
	})
	e.GET("/user/profile", func(c echo.Context) error {
		return c.String(http.StatusOK, "/user/profile")
	})
	e.GET("/user/profile/pics", func(c echo.Context) error {
		return c.String(http.StatusOK, "/user/profile/pics")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
