package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anthanhphan/saturday/http/constant/method"
	"github.com/anthanhphan/saturday/http/middlewares"
	"github.com/anthanhphan/saturday/http/resp"
	"github.com/anthanhphan/saturday/http/route"
	"github.com/anthanhphan/saturday/http/server"
	"github.com/anthanhphan/saturday/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Start(srv *server.HttpServer) {
	log := zap.L().With(zap.String("prefix", "start")).Sugar()

	srv.Middlewares = append(srv.Middlewares, middlewares.RequestId())
	srv.Middlewares = append(srv.Middlewares, middlewares.Recover())

	// Setup route
	r := route.NewGinEngine(
		route.AddMiddlewares(srv.Middlewares...),
		route.AddHealthCheckRoute(),
		route.AddRouteNotFoundHandler(),
		route.SetStrictSlash(srv.StrictSlash),
		route.SetMaximumMultipartSize(10000000),
		route.AddGroupRoutes(srv.GroupRoutes),
		route.AddRoutes(srv.Routes),
		route.AddGinOptions(srv.GinOptions...),
	)

	// Create an HTTP server instance with the Gin engine
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", srv.Port),
		Handler: r,
	}

	// Channel to listen for interrupt signals (e.g., CTRL+C, SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go func() {
		log.Infof("http server started on port %v", srv.Port)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("could not start server: %v\n", err)
		}
	}()

	// Wait for interrupt signal
	<-quit
	log.Infof("shutting down server...")

	// Graceful shutdown with a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Infof("server exited gracefully")
}

func main() {
	logInstance, undo := logger.InitLogger(&logger.Config{
		DisableCaller:     false,
		DisableStacktrace: true,
		EnableDevMode:     true,
		Level:             logger.LevelInfo,
		Encoding:          logger.EncodingConsole,
	})
	defer func() {
		_ = logInstance.Sync()
	}()
	defer undo()

	httpServer := server.NewHttpServer(
		server.AddPort(int64(5000)),
		server.AddName("Test Server"),
		server.SetStrictSlash(true),
		server.SetGracefulShutdownTimeout(time.Duration(10000)),
	)

	httpServer.AddRoutes([]route.Route{
		{
			Path:    "/hello-world",
			Method:  method.GET,
			Handler: SayHelloWorld,
		},
	})

	Start(httpServer)
}

func SayHelloWorld(ctx *gin.Context) {
	resp.ResponseSuccess(ctx, resp.NewSuccessResp("Hello World!", nil))
}
