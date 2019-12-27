package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic: [%v]\n", stack())

			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()

	c.Next()
}

func stack() string {
	var buf [2 << 10]byte
	return string(buf[:runtime.Stack(buf[:], true)])
}
