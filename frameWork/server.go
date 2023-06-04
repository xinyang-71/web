package frameWork

import (
	"net"
	"net/http"
)

var _ Server = &HttpServer{}

type HandleFunc func(ctx Context)

type Server interface {
	http.Handler
	Start(addr string) error

	// 增加一些路由注册功能
	AddRoute(method string, path string, handleFunc HandleFunc)
}

type HttpServer struct {
	*router
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		router: newRouter(),
	}
}

func (h *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	h.Serve(ctx)
}

func (h *HttpServer) Serve(ctx *Context) {
}

//func (h *HttpServer) AddRoute(method string, path string, handleFunc HandleFunc) {
//
//}

func (h *HttpServer) Get(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodGet, path, handleFunc)
}

func (h *HttpServer) Post(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodPost, path, handleFunc)
}

func (h *HttpServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return http.Serve(l, h)
}
