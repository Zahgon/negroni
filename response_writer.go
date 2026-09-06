package negroni

import (
	"io"
	"net/http"
)

type ResponseWriter interface {
	http.ResponseWriter

	Status() int

	Written() bool

	Size() int

	Before(func(ResponseWriter))
}

type beforeFunc func(ResponseWriter)

func NewResponseWriter(rw http.ResponseWriter) ResponseWriter {
	_ = "STUB: not implemented"
	return *new(ResponseWriter)
}

type responseWriter struct {
	http.ResponseWriter
	pendingStatus  int
	status         int
	size           int
	beforeFuncs    []beforeFunc
	callingBefores bool
}

func (rw *responseWriter) WriteHeader(s int) { _ = "STUB: not implemented"; return }

func (rw *responseWriter) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (rw *responseWriter) ReadFrom(r io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rw *responseWriter) Unwrap() http.ResponseWriter {
	_ = "STUB: not implemented"
	return *new(http.ResponseWriter)
}

func (rw *responseWriter) Status() int { _ = "STUB: not implemented"; return 0 }

func (rw *responseWriter) Size() int { _ = "STUB: not implemented"; return 0 }

func (rw *responseWriter) Written() bool { _ = "STUB: not implemented"; return false }

func (rw *responseWriter) Before(before func(ResponseWriter)) { _ = "STUB: not implemented"; return }

func (rw *responseWriter) callBefore() { _ = "STUB: not implemented"; return }
