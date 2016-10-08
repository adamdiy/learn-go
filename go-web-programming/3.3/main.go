package main

import (
	"net/http"
)

type Server struct {
	Addr           string
	Handler        Handler
	ReadTimeOut    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
	TLSConfig      *tls.Config
	TLSNextProto   map[string]func(*Server, *tls.Conn, Handler)
	ConnState      func(net.Conn, ConnState)
	ErrorLog       *log.Logger
}

func main() {

}
