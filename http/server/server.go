package server

import (
	"time"

	"github.com/anthanhphan/saturday/http/route"
	"github.com/gin-gonic/gin"
)

// Option is a function type that modifies HttpServer configuration.
type Option func(*HttpServer)

// HttpServer represents the HTTP server configuration and components.
//
// Fields:
//   - Name: Server name identifier
//   - Port: Port number to listen on
//   - StrictSlash: Whether to enforce strict slash handling in routes
//   - Routes: List of individual routes
//   - GroupRoutes: List of grouped routes
//   - Middlewares: Global middleware functions
//   - GinOptions: Gin-specific configuration options
//   - OnCloseFunc: Function to execute on server shutdown
//   - GracefulShutdownTimeout: Maximum time to wait for graceful shutdown
type HttpServer struct {
	Name                    string
	Port                    int64
	StrictSlash             bool
	Routes                  []route.Route
	GroupRoutes             []route.GroupRoute
	Middlewares             []func(ctx *gin.Context)
	GinOptions              []route.GinOption
	OnCloseFunc             func()
	GracefulShutdownTimeout time.Duration
}

// NewHttpServer creates a new HTTP server instance with the provided options.
//
// Parameters:
//   - options: Variable number of Option functions to configure the server
//
// Returns:
//   - *HttpServer: A new configured HTTP server instance
//
// Example:
//
//	server := NewHttpServer(
//	    AddName("api-server"),
//	    AddPort(8080),
//	)
func NewHttpServer(options ...Option) *HttpServer {
	s := &HttpServer{}
	for _, option := range options {
		option(s)
	}
	return s
}

// AddName returns an Option to set the server name.
//
// Parameters:
//   - name: String identifier for the server
//
// Returns:
//   - Option: Function that modifies the HttpServer name
//
// Example:
//
//	server := NewHttpServer(AddName("api-server"))
func AddName(name string) Option {
	return func(server *HttpServer) {
		server.Name = name
	}
}

// AddPort returns an Option to set the server port.
//
// Parameters:
//   - port: Port number to listen on
//
// Returns:
//   - Option: Function that modifies the HttpServer port
//
// Example:
//
//	server := NewHttpServer(AddPort(8080))
func AddPort(port int64) Option {
	return func(server *HttpServer) {
		server.Port = port
	}
}

// AddMiddlewares returns an Option to append middleware functions.
//
// Parameters:
//   - middleware: Slice of Gin middleware functions to add
//
// Returns:
//   - Option: Function that adds middlewares to the HttpServer
//
// Example:
//
//	middlewares := []func(*gin.Context){
//	    gin.Logger(),
//	    gin.Recovery(),
//	}
//	server := NewHttpServer(AddMiddlewares(middlewares))
func AddMiddlewares(middleware []func(c *gin.Context)) Option {
	return func(server *HttpServer) {
		server.Middlewares = append(server.Middlewares, middleware...)
	}
}

// AddRoutes adds individual routes to the server configuration.
//
// Parameters:
//   - route: Slice of Route configurations to add
//
// Example:
//
//	routes := []route.Route{
//	    {Path: "/health", Method: "GET", Handler: healthCheck},
//	}
//	server.AddRoutes(routes)
func (server *HttpServer) AddRoutes(route []route.Route) {
	server.Routes = route
}

// AddGroupRoutes adds grouped routes to the server configuration.
//
// Parameters:
//   - groupRoute: Slice of GroupRoute configurations to add
//
// Example:
//
//	groupRoutes := []route.GroupRoute{
//	    {Prefix: "/api/v1", Routes: []route.Route{{Path: "/users", Method: "GET", Handler: getUsers}}},
//	}
//	server.AddGroupRoutes(groupRoutes)
func (server *HttpServer) AddGroupRoutes(groupRoute []route.GroupRoute) {
	server.GroupRoutes = groupRoute
}

// SetStrictSlash returns an Option to configure strict slash handling.
//
// Parameters:
//   - strict: Boolean indicating whether to enable strict slash handling
//
// Returns:
//   - Option: Function that sets the StrictSlash configuration
//
// Example:
//
//	server := NewHttpServer(SetStrictSlash(true))
func SetStrictSlash(strict bool) Option {
	return func(server *HttpServer) {
		server.StrictSlash = strict
	}
}

// SetGracefulShutdownTimeout returns an Option to set the graceful shutdown timeout duration.
//
// Parameters:
//   - t: Duration to wait for server shutdown
//
// Returns:
//   - Option: Function that sets the shutdown timeout
//
// Example:
//
//	server := NewHttpServer(SetGracefulShutdownTimeout(5 * time.Second))
func SetGracefulShutdownTimeout(t time.Duration) Option {
	return func(server *HttpServer) {
		server.GracefulShutdownTimeout = t
	}
}
