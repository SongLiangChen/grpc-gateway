package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	cnf := cors.DefaultConfig()
	cnf.AllowOriginFunc = func(origin string) bool {
		return true
	}
	cnf.AllowCredentials = true
	return cors.New(cnf)
}
