package gateway

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"grpc-gateway/common"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

var (
	ErrPathRegistered    = errors.New("register to the same path")
	ErrRegisterFuncEmpty = errors.New("register func empty")
)

type Gateway struct {
	// mux is a global serveMux
	mux *http.ServeMux

	server *http.Server

	ctx    context.Context
	cancel context.CancelFunc

	registeredRelativePath map[string]struct{}
}

func (g *Gateway) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	g.mux.HandleFunc(pattern, handler)
}

func (g *Gateway) dial(network, addr string) (*grpc.ClientConn, error) {
	switch network {
	case "tcp":
		return grpc.DialContext(g.ctx, addr, grpc.WithInsecure())
	case "unix":
		d := func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", addr, timeout)
		}
		return grpc.DialContext(g.ctx, addr, grpc.WithInsecure(), grpc.WithDialer(d))
	default:
		return nil, fmt.Errorf("unsupported network type %q", network)
	}
}

type EndServerConfig struct {
	// All request will router to this server if RelativePath == "/".
	RelativePath string

	Network  string                   // tcp, unix
	Addr     string                   // ip:port
	Opts     []runtime.ServeMuxOption // use for grpc-gateway/runtime.NewServeMux
	Handlers []gin.HandlerFunc        // gin middleware

	// RegisterFunc register the http handlers to "mux", see *.pb.gw.go file for detail
	RegisterFunc func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
}

// RegisterEndServer register a new end server
func (g *Gateway) RegisterEndServer(c *EndServerConfig) error {
	if _, ok := g.registeredRelativePath[c.RelativePath]; ok {
		return ErrPathRegistered
	}

	if c.RegisterFunc == nil {
		return ErrRegisterFuncEmpty
	}

	conn, err := g.dial(c.Network, c.Addr)
	if err != nil {
		return err
	}

	mux := runtime.NewServeMux(c.Opts...)
	if err = c.RegisterFunc(g.ctx, mux, conn); err != nil {
		conn.Close()
		return err
	}

	engine := gin.New()
	engine.Use(c.Handlers...)
	engine.Use(gatewayLog(mux))

	engine.Any("/*any", func(c *gin.Context) {
		mux.ServeHTTP(c.Writer, c.Request)
	})

	go func() {
		<-g.ctx.Done()
		if err := conn.Close(); err != nil {
			fmt.Printf("Failed to close a client connection to the gRPC server: %v\n", err)
		}
	}()

	g.mux.Handle(c.RelativePath, engine)

	g.registeredRelativePath[c.RelativePath] = struct{}{}

	return nil
}

// Run init http.Server and call server.ListenAndServe.
// Note: this method will NOT execute in a new goroutine.
func (g *Gateway) Run(addr string) error {
	g.server = &http.Server{
		Addr:    addr,
		Handler: g.mux,
	}

	if err := g.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (g *Gateway) Close() error {
	defer g.cancel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	println("Deal unfinished request...wait...")
	if err := g.server.Shutdown(ctx); err != nil {
		println("Server Shutdown:", err)
	}

	return nil
}

func NewGateway() *Gateway {
	ctx, cancel := context.WithCancel(context.Background())

	return &Gateway{
		mux:    http.NewServeMux(),
		ctx:    ctx,
		cancel: cancel,

		registeredRelativePath: make(map[string]struct{}),
	}
}

func gatewayLog(mux *runtime.ServeMux) gin.HandlerFunc {
	once := sync.Once{}

	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "body", common.JsonBytesToMap(buf))
		ctx = context.WithValue(ctx, "url", c.Request.URL.String())
		ctx = context.WithValue(ctx, "header", c.Request.Header)
		ctx = context.WithValue(ctx, "beginTime", time.Now())

		c.Request = c.Request.WithContext(ctx)
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(buf))

		once.Do(func() {
			runtime.WithForwardResponseOption(func(i context.Context, writer http.ResponseWriter, message proto.Message) error {
				var (
					body         = i.Value("body")
					header       = i.Value("header")
					urlStr, _    = i.Value("url").(string)
					beginTime, _ = i.Value("beginTime").(time.Time)
				)

				log.WithFields(log.Fields{
					"requestBody": body,
					"header":      header,
					"url":         urlStr,
					"response":    message,
					"ttl":         time.Since(beginTime).String(),
				}).Info("")

				return nil
			})(mux)
		})
	}
}
