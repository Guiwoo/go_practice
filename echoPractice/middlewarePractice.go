package echoPractice

import (
	"context"
	"crypto/subtle"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Custom struct {
	time time.Time
	path string
}

func NewCustom() *Custom {
	return &Custom{
		time: time.Now(),
		path: "",
	}
}

func (cc *Custom) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		log.Printf("[%s][%s] %s %s ", req.Method, "Before Access hanlder", req.URL, cc.time)
		err := next(c)
		if err != nil {
			log := c.Logger()
			log.Error("Can't Response Error")
		}
		res := c.Response()
		log.Printf("[%s][%s] %s %s ", req.Method, res.Status, req.URL, time.Now())

		return nil
	}
}

func ExMiddleware() {
	e := echo.New()

	cc := NewCustom()

	e.Use(cc.Process)

	e.GET("/products", func(c echo.Context) error {
		fmt.Println("Handler")
		return c.String(http.StatusOK, "Product List")
	})
	e.GET("/products/1", func(c echo.Context) error {
		fmt.Println("Handler")
		return c.String(http.StatusOK, "Product 1 details")
	})
	e.GET("/products/1/order", func(c echo.Context) error {
		fmt.Println("Handler")
		return c.String(http.StatusOK, "Product 1 to order")
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func serverSetting(e *echo.Echo) {
	e.GET("/products", func(c echo.Context) error {
		fmt.Println("Handler")
		return c.String(http.StatusOK, "Product List")
	})
	e.GET("/products/1", func(c echo.Context) error {
		fmt.Println("Handler")
		return c.String(http.StatusOK, "Product 1 details")
	})
	e.GET("/products/1/order", func(c echo.Context) error {
		fmt.Println("Handler")
		return c.String(http.StatusOK, "Product 1 to order")
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func BasicAuth() {
	e := echo.New()

	e.Use(middleware.BasicAuth(func(s string, s2 string, context echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(s), []byte("guiwoo")) == 1 &&
			subtle.ConstantTimeCompare([]byte(s2), []byte("123")) == 1 {
			return true, nil
		}
		return false, nil
	}))
	serverSetting(e)
}

func ExCorsMiddleware() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	serverSetting(e)
	e.Use(middleware.Gzip())
}

func ExJwt() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))

	e.GET("/", func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
		if !ok {
			return errors.New("JWT token missing or invalid")
		}
		claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
		if !ok {
			return errors.New("failed to cast claims as jwt.MapClaims")
		}
		return c.JSON(http.StatusOK, claims)
	})

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

type jwtCustomClaims struct {
	Name  string
	Admin bool
	jwt.RegisteredClaims
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username != "guiwoo" || password != "123" {
		return echo.ErrUnauthorized
	}

	claims := &jwtCustomClaims{
		"Guiwoo park",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("GuiwooSecret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func CustomJwt() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", login)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Can Access")
	})

	r := e.Group("/restricted")
	r.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("GuiwooSecret"),
	}))
	r.GET("", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)
		name := claims.Name
		return c.String(http.StatusOK, "welcome "+name+" !")
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func RedirectMiddleware() {
	e := echo.New()
	e.Pre(middleware.HTTPSRedirect())

	serverSetting(e)
}

func GracefulShutdown() {
	e := echo.New()
	e.Logger.SetLevel(2)
	e.GET("/", func(c echo.Context) error {
		time.Sleep(5 * time.Second)
		return c.JSON(http.StatusOK, "OK")
	})

	go func() {
		if err := e.Start(":8080"); err != nil && err == http.ErrServerClosed {
			e.Logger.Fatal("Shutting down")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("Before Shutdown")
	if err := e.Shutdown(ctx); err != nil {
		fmt.Println("hit")
		e.Logger.Fatal(err)
	}
}

func Ex11() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// create a new HTTP client with the context
	client := &http.Client{Timeout: time.Second * 10}

	// send an HTTP GET request to an external API with the client
	req, err := http.NewRequest("GET", "https://api.example.com/data", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// process the response
	fmt.Println("Response status:", resp.Status)
}
