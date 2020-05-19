package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func testGet(c *gin.Context) {
	data := map[string]interface{}{
		"name": "abelit",
		"age":  28,
	}
	c.JSON(http.StatusOK, data)
}
