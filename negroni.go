package negroni

import (
	"net/http"
)

const (
	DefaultAddress = ":8080"
)

func init() {
	initFeaturePicker()
}

type Handler interface {
	ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type HandlerFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (h HandlerFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

type middleware struct {
	handler Handler

	nextfn func(rw http.ResponseWriter, r *http.Request)
}

func newMiddleware(handler Handler, next *middleware) middleware {
	_ = "STUB: not implemented"
	return *new(middleware)
}

func (m middleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func Wrap(handler http.Handler) Handler { _ = "STUB: not implemented"; return *new(Handler) }

func WrapFunc(handlerFunc http.HandlerFunc) Handler {
	_ = "STUB: not implemented"
	return *new(Handler)
}

type Negroni struct {
	middleware middleware
	handlers   []Handler
}

func New(handlers ...Handler) *Negroni { _ = "STUB: not implemented"; return nil }

func (n *Negroni) With(handlers ...Handler) *Negroni { _ = "STUB: not implemented"; return nil }

func Classic() *Negroni { _ = "STUB: not implemented"; return nil }

func (n *Negroni) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (n *Negroni) Use(handler Handler) { _ = "STUB: not implemented"; return }

func (n *Negroni) UseFunc(handlerFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) {
	_ = "STUB: not implemented"
	return
}

func (n *Negroni) UseHandler(handler http.Handler) { _ = "STUB: not implemented"; return }

func (n *Negroni) UseHandlerFunc(handlerFunc func(rw http.ResponseWriter, r *http.Request)) {
	_ = "STUB: not implemented"
	return
}

func (n *Negroni) Run(addr ...string) { _ = "STUB: not implemented"; return }

func detectAddress(addr ...string) string { _ = "STUB: not implemented"; return "" }

func (n *Negroni) Handlers() []Handler { _ = "STUB: not implemented"; return nil }

func build(handlers []Handler) middleware { _ = "STUB: not implemented"; return *new(middleware) }

func voidMiddleware() middleware { _ = "STUB: not implemented"; return *new(middleware) }
