package route

import (
	"net/http"

	"github.com/anthanhphan/saturday/http/resp"
	"github.com/gin-gonic/gin"
)

type GinOption func(*gin.Engine)

// NewGinEngine creates and configures a new Gin engine instance.
//
// Parameters:
//   - options: Variable number of GinOption functions to configure the engine
//
// Returns:
//   - *gin.Engine: A configured Gin engine instance
//
// Examples:
//
//	engine := NewGinEngine()  // creates basic engine
//	engine := NewGinEngine(AddHealthCheckRoute(), SetStrictSlash(true))
func NewGinEngine(options ...GinOption) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	for _, option := range options {
		option(r)
	}

	return r
}

// AddMiddlewares creates a GinOption that adds middleware functions to the engine.
//
// Parameters:
//   - ms: Variable number of middleware functions
//
// Returns:
//   - GinOption: Function that applies the middlewares
//
// Examples:
//
//	logger := func(c *gin.Context) { /* logging logic */ }
//	auth := func(c *gin.Context) { /* auth logic */ }
//	engine := NewGin(AddMiddlewares(logger, auth))
func AddMiddlewares(ms ...func(c *gin.Context)) GinOption {
	return func(g *gin.Engine) {
		for _, m := range ms {
			g.Use(m)
		}
	}
}

// AddGroupRoutes creates a GinOption that adds route groups to the engine.
//
// Parameters:
//   - gr: Slice of GroupRoute configurations
//
// Returns:
//   - GinOption: Function that applies the route groups
//
// Examples:
//
//	groups := []GroupRoute{
//	    {Path: "/api/v1", Routes: []Route{...}},
//	}
//	engine := NewGin(AddGroupRoutes(groups))
func AddGroupRoutes(gr []GroupRoute) GinOption {
	return func(g *gin.Engine) {
		for _, r := range gr {
			r.AddGroupRoute(g)
		}
	}
}

// AddRoutes creates a GinOption that adds individual routes to the engine.
//
// Parameters:
//   - rs: Slice of Route configurations
//
// Returns:
//   - GinOption: Function that applies the routes
//
// Examples:
//
//	routes := []Route{
//	    {Method: "GET", Path: "/users", Handler: handleUsers},
//	}
//	engine := NewGin(AddRoutes(routes))
func AddRoutes(rs []Route) GinOption {
	return func(g *gin.Engine) {
		for _, r := range rs {
			r.AddRoute(g)
		}
	}
}

// SetStrictSlash creates a GinOption that configures URL trailing slash behavior.
//
// Parameters:
//   - strict: When true, removes extra slashes from URLs
//
// Returns:
//   - GinOption: Function that configures slash handling
//
// Examples:
//
//	engine := NewGin(SetStrictSlash(true))
//	// /api/users/ and /api/users will route to the same handler
func SetStrictSlash(strict bool) GinOption {
	return func(g *gin.Engine) {
		g.RemoveExtraSlash = strict
	}
}

// SetMaximumMultipartSize creates a GinOption that sets the maximum multipart form size.
//
// Parameters:
//   - size: Maximum size in bytes
//
// Returns:
//   - GinOption: Function that sets the size limit
//
// Examples:
//
//	maxSize := 8 << 20  // 8 MB
//	engine := NewGin(SetMaximumMultipartSize(maxSize))
func SetMaximumMultipartSize(size int64) GinOption {
	return func(g *gin.Engine) {
		g.MaxMultipartMemory = size
	}
}

// AddGinOptions combines multiple GinOptions into a single option.
//
// Parameters:
//   - options: Variable number of GinOption functions
//
// Returns:
//   - GinOption: Function that applies all options
//
// Examples:
//
//	combined := AddGinOptions(
//	    AddHealthCheckRoute(),
//	    SetStrictSlash(true),
//	)
//	engine := NewGin(combined)
func AddGinOptions(options ...GinOption) GinOption {
	return func(e *gin.Engine) {
		for _, o := range options {
			o(e)
		}
	}
}

// AddHealthCheckRoute creates a GinOption that adds a health check endpoint.
//
// Returns:
//   - GinOption: Function that adds the /health-check route
//
// Examples:
//
//	engine := NewGin(AddHealthCheckRoute())
//	// Adds GET /health-check endpoint returning {"server_status": "healthy"}
func AddHealthCheckRoute() GinOption {
	return func(g *gin.Engine) {
		g.GET("/health-check", func(ctx *gin.Context) {
			data := map[string]interface{}{
				"server_status": "healthy",
			}
			resp.ResponseSuccess(ctx, resp.NewSuccessResp("success", data))
		})
	}
}

// AddRouteNotFoundHandler creates a GinOption that configures the 404 handler.
//
// Returns:
//   - GinOption: Function that sets the not found handler
//
// Examples:
//
//	engine := NewGin(AddRouteNotFoundHandler())
//	// Returns standardized 404 error for unmatched routes
func AddRouteNotFoundHandler() GinOption {
	return func(g *gin.Engine) {
		g.NoRoute(func(g *gin.Context) {
			panic(resp.NewErrorResp(http.StatusNotFound, nil, "api not found"))
		})
	}
}
