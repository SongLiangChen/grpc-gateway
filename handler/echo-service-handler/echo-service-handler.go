package echo_service_handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"grpc-gateway/handler"
	"grpc-gateway/proto/echo-service"
)

type EchoServiceHandler struct {
}

func (*EchoServiceHandler) NewEngine(ctx context.Context, network, addr string, opts []runtime.ServeMuxOption, handleFuncs ...gin.HandlerFunc) (*gin.Engine, error) {
	conn, err := handler.Dial(ctx, network, addr)
	if err != nil {
		return nil, err
	}

	mux := runtime.NewServeMux(opts...)
	if err = echo_service.RegisterEchoServiceHandler(ctx, mux, conn); err != nil {
		conn.Close()
		return nil, err
	}

	engine := gin.New()
	engine.Use(handleFuncs...)
	engine.Use(func(c *gin.Context) {
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
	handler.RegisterHandler("echo", &EchoServiceHandler{})
}
