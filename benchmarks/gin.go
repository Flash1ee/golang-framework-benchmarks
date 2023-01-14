package benchmarks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	h := gin.New()
	h.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World")
	})
	h.POST("/", func(c *gin.Context) {
		var req Request[[]interface{}]
		if err := c.Bind(&req); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if len(req.Data) == 0 {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, req.Data[len(req.Data)-1])
	})
	h.GET("/param/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s", name)
	})

	RegisterHandler("gin", h)
}
