package user_service_handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"grpc-gateway/handler"
	"grpc-gateway/proto/user-service"
)

type UserServiceHandler struct {
}

func (*UserServiceHandler) NewEngine(ctx context.Context, network, addr string, opts []runtime.ServeMuxOption, handleFuncs ...gin.HandlerFunc) (*gin.Engine, error) {
	conn, err := handler.Dial(ctx, network, addr)
	if err != nil {
		return nil, err
	}

	mux := runtime.NewServeMux(opts...)
	if err = user_service.RegisterUserServiceHandler(ctx, mux, conn); err != nil {
		conn.Close()
		return nil, err
	}

	engine := gin.Default()
	engine.Use(handleFuncs...)

	engine.Any("/*any", func(c *gin.Context) {
		mux.ServeHTTP(c.Writer, c.Request)
	})

	go func() {
		<-ctx.Done()
		if err := conn.Close(); err != nil {
			fmt.Printf("Failed to close a client connection to the gRPC server: %v\n", err)
		}
	}()

	return engine, nil
}

func init() {
	handler.RegisterHandler("01", &UserServiceHandler{})
}
