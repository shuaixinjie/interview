package designpattren

import "time"

// 创建者模式

// go语言适合的函数选项模式

type Server struct {
	host string
	port int
}

// New 常见的建造者模式，可读性很差
func New(host string, port int) *Server {
	return &Server{
		host: "127.0.0.1",
		port: 8080,
	}
}

// Option 函数选项模式
type Option func(server *Server)

func New2(options ...Option) *Server {
	svr := new(Server)
	for _, f := range options {
		f(svr)
	}
	return svr
}

func WithHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

// 函数选项模式，一种更为简洁的写法

type Server2 struct {
	host    string
	port    int
	timeout time.Duration
}

func NewServer2(addr string, options ...func(server *Server2)) (*Server2, error) {
	svr := &Server2{
		host: addr,
	}
	for _, f := range options {
		f(svr)
	}
	return svr, nil
}
