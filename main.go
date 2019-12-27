package main

import (
	"context"
	"grpc-gateway/handler"
	"grpc-gateway/middleware"
	"net/http"

	_ "grpc-gateway/handler/echo-service-handler"
	_ "grpc-gateway/handler/user-service-handler"
)

type ServerEndpoint struct {
	Key     string
	Network string
	Addr    string
}

var (
	grpcServerEndpoint = []ServerEndpoint{
		{
			"01",
			"tcp",
			"127.0.0.1:18080",
		},
	}
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// gin.SetMode(gin.ReleaseMode)

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", middleware.SwaggerServer())

	for _, ep := range grpcServerEndpoint {
		h, err := handler.NewHandler(ep.Key, ctx, ep.Network, ep.Addr, nil,
			middleware.Recover, middleware.Cors())
		if err != nil {
			println(err.Error())
			return
		}

		mux.Handle("/"+ep.Key+"/", h)
	}

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := s.ListenAndServe(); err != nil {
		println(err.Error())
	}
}
