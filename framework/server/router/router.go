package router

import (
	"net/http"
)

func New() *Mux {
	return NewMux()
}

// TODO: Map the `root_path` type handle to the controllers, that way we can
// potentially declare Before/After action(route) via these routes in a map.

type Router interface {
	http.Handler
	Routes
	Use(middlewares ...func(http.Handler) http.Handler)
	With(middlewares ...func(http.Handler) http.Handler) Router
	Group(fn func(r Router)) Router

	// TODO: Should the BeforeRoute and AfterRoute (in Rails Before and After
	// Action be a middleware or just built into the router?)
	// Either we need to auto-magicaly generate the path variables, that are
	// familiar to Rails developers:
	// (i.e. root_path, user_login_path, register_user_path, etc)
	// Or as in Rails we can have them be defined during routes declaration.
	// Such as
	Route(pattern string, fn func(r Router)) Router
	Mount(pattern string, h http.Handler)
	Handle(pattern string, h http.Handler)
	HandleFunc(pattern string, h http.HandlerFunc)
	Method(method, pattern string, h http.Handler)
	MethodFunc(method, pattern string, h http.HandlerFunc)
	Connect(pattern string, h http.HandlerFunc)
	Delete(pattern string, h http.HandlerFunc)
	Get(pattern string, h http.HandlerFunc)
	Head(pattern string, h http.HandlerFunc)
	Options(pattern string, h http.HandlerFunc)
	Patch(pattern string, h http.HandlerFunc)
	Post(pattern string, h http.HandlerFunc)
	Put(pattern string, h http.HandlerFunc)
	Trace(pattern string, h http.HandlerFunc)

	NotFound(h http.HandlerFunc)
	MethodNotAllowed(h http.HandlerFunc)
}

type Routes interface {
	Routes() []Route
	Middlewares() Middlewares
	Match(rctx *Context, method, path string) bool
}

type Middlewares []func(http.Handler) http.Handler
