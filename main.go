package main

import (
	"github.com/gin-gonic/gin"
	"grpc-gateway/common"
	"grpc-gateway/config"
	"grpc-gateway/gateway"
	"grpc-gateway/middleware"
	"grpc-gateway/proto/user-service"
)

func main() {
	if err := config.InitConfig(); err != nil {
		println(err.Error())
		return
	}
	cnf := config.GetConfig()

	common.ConfigLogger(cnf.LogCnf.Dir, cnf.LogCnf.Name, cnf.LogCnf.KeepDay, cnf.LogCnf.RateHours)

	gw := gateway.NewGateway()

	gw.HandleFunc("/swagger/", middleware.SwaggerServer())

	if err := gw.RegisterEndServer(&gateway.EndServerConfig{
		RelativePath: "/01/",
		Network:      "tcp",
		Addr:         "127.0.0.1:18080",
		// Opts:         []runtime.ServeMuxOption{runtime.WithForwardResponseOption(forwardResponseOption)},
		Handlers:     []gin.HandlerFunc{middleware.Recover, middleware.Cors()},
		RegisterFunc: user_service.RegisterUserServiceHandler,
	}); err != nil {
		println(err.Error())
		return
	}

	defer gw.Close()

	gw.Run(cnf.AppCnf.Addr)
}
