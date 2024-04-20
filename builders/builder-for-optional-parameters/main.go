package main

import (
	"crypto/tls"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	server := NewWebServer(db).
		ListenOn(":8080").
		TLS("server.crt", "server.key").
		Logger(&log.Logger{}).
		Middlewares(
			middleware.Logger,
			middleware.RequestID,
			middleware.Heartbeat("/ping"),
		)

	// Configure the web server
	if err := server.Start(); err != nil {
		panic(err)
	}

	// Later, stop the web server
	err = server.Stop()
	if err != nil {
		panic(err)
	}
}

type Middleware func(next http.Handler) http.Handler

type WebServer struct {
	address     string
	tlsConfig   *tls.Config
	logger      *log.Logger
	middlewares []Middleware
	isRunning   bool
}

func NewWebServer(db *gorm.DB) *WebServer {
	return &WebServer{
		address: ":8000",
	}
}

func (ws *WebServer) ListenOn(address string) *WebServer {
	ws.address = address
	return ws
}

func (ws *WebServer) TLS(certFile, keyFile string) *WebServer {
	// Initialize and configure the TLS configuration
	ws.tlsConfig = &tls.Config{ /* ... */ }
	return ws
}

func (ws *WebServer) Logger(logger *log.Logger) *WebServer {
	ws.logger = logger
	return ws
}

func (ws *WebServer) Middlewares(middleware ...Middleware) *WebServer {
	ws.middlewares = middleware
	return ws
}

func (ws *WebServer) Start() error {
	if ws.isRunning {
		return errors.New("web server already running")
	}

	// Start the web server using the configured settings
	// ...

	ws.isRunning = true
	return nil
}

func (ws *WebServer) Stop() error {
	// Implement stopping the web server
	// ...
	return nil
}
