package benchmarks

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetEchoApp() *echo.Echo {
	h := echo.New()
	h.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	h.POST("/", func(c echo.Context) error {
		var req Request[[]interface{}]
		if err := c.Bind(&req); err != nil {
			c.Response().WriteHeader(http.StatusBadRequest)
			c.Response().Write([]byte(err.Error()))
			return err
		}

		if len(req.Data) == 0 {
			c.Response().WriteHeader(http.StatusBadRequest)
			return nil
		}

		c.JSON(http.StatusOK, req.Data[len(req.Data)-1])
		return nil
	})
	h.GET("/param/:name", func(c echo.Context) error {
		name := c.Param("name")
		c.JSON(http.StatusOK, fmt.Sprintf("Hello, %s", name))
		return nil
	})

	RegisterHandler("echo", h)
	return h
}

func StartEcho() {
	DeleteHandler("echo")
	h := GetEchoApp()
	go h.Start(":3001")
}
