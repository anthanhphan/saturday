package route

import (
	"github.com/anthanhphan/saturday/http/constant/method"
	"github.com/gin-gonic/gin"
)

type GroupRoute struct {
	Prefix      string
	Middlewares []func(*gin.Context)
	Routes      []Route
}

type Route struct {
	Path        string
	Method      method.Method
	Handler     func(*gin.Context)
	Middlewares []func(*gin.Context)
}

// CombineHandler merges route middlewares and the main handler into a single slice.
//
// Returns:
//   - []gin.HandlerFunc: Combined slice of middleware and handler functions
//
// Examples:
//
//	route := Route{
//	    Handler: mainHandler,
//	    Middlewares: []func(*gin.Context){auth, logger},
//	}
//	handlers := route.CombineHandler()
//	// handlers contains [auth, logger, mainHandler]

func (route Route) CombineHandler() []gin.HandlerFunc {
	var handler []gin.HandlerFunc
	for _, middleware := range route.Middlewares {
		handler = append(handler, middleware)
	}
	handler = append(handler, route.Handler)
	return handler
}

// AddRoute registers a single route with the Gin engine based on its method.
//
// Parameters:
//   - g: The Gin engine to add the route to
//
// Examples:
//
//	route := Route{Path: "/users", Method: method.GET, Handler: handleUsers}
//	route.AddRoute(engine)
//	// Registers GET /users endpoint

// AddRoute adds a route to the gin.Engine based on the Route's Method and Path.
//
// It takes a pointer to a gin.Engine as the parameter.
// It does not return anything.
func (route Route) AddRoute(g *gin.Engine) {
	switch route.Method {
	case method.GET:
		g.GET(route.Path, route.CombineHandler()...)
	case method.POST:
		g.POST(route.Path, route.CombineHandler()...)
	case method.PUT:
		g.PUT(route.Path, route.CombineHandler()...)
	case method.PATCH:
		g.PATCH(route.Path, route.CombineHandler()...)
	case method.DELETE:
		g.DELETE(route.Path, route.CombineHandler()...)
	case method.HEAD:
		g.HEAD(route.Path, route.CombineHandler()...)
	case method.OPTIONS:
		g.OPTIONS(route.Path, route.CombineHandler()...)
	}
}

// AddGroupRoute registers a group of routes with the Gin engine.
//
// Parameters:
//   - g: The Gin engine to add the route group to
//
// Examples:
//
//	group := GroupRoute{
//	    Prefix: "/api",
//	    Routes: []Route{{Path: "/users", Method: method.GET, Handler: handleUsers}},
//	}
//	group.AddGroupRoute(engine)
//	// Registers GET /api/users endpoint

// AddGroupRoute adds a group route to the gin.Engine instance.
//
// Parameters:
// - g: a pointer to a gin.Engine instance.
//
// Return type: none
func (route GroupRoute) AddGroupRoute(g *gin.Engine) {
	groupRoute := g.Group(route.Prefix)
	for _, middleware := range route.Middlewares {
		groupRoute.Use(middleware)
	}
	for _, route := range route.Routes {
		switch route.Method {
		case method.GET:
			groupRoute.GET(route.Path, route.CombineHandler()...)
		case method.POST:
			groupRoute.POST(route.Path, route.CombineHandler()...)
		case method.PUT:
			groupRoute.PUT(route.Path, route.CombineHandler()...)
		case method.PATCH:
			groupRoute.PATCH(route.Path, route.CombineHandler()...)
		case method.DELETE:
			groupRoute.DELETE(route.Path, route.CombineHandler()...)
		case method.HEAD:
			groupRoute.HEAD(route.Path, route.CombineHandler()...)
		case method.OPTIONS:
			groupRoute.OPTIONS(route.Path, route.CombineHandler()...)
		}
	}
}
