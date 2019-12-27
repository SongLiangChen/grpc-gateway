package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type Handler interface {
	NewEngine(ctx context.Context, network, addr string, opts []runtime.ServeMuxOption, handleFuncs ...gin.HandlerFunc) (*gin.Engine, error)
}

func RegisterHandler(key string, handler Handler) {
	if _, ok := handlers[key]; ok {
		panic(fmt.Sprintf("handler[%v] already register!", key))
	}
	handlers[key] = handler
}

func NewHandler(key string, ctx context.Context, network, addr string, opts []runtime.ServeMuxOption, handleFuncs ...gin.HandlerFunc) (*gin.Engine, error) {
	if _, ok := handlers[key]; !ok {
		panic(fmt.Sprintf("handler[%v] not exist!", key))
	}

	return handlers[key].NewEngine(ctx, network, addr, opts, handleFuncs...)
}

var (
	handlers map[string]Handler
)

func init() {
	handlers = make(map[string]Handler)
}
