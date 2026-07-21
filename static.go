package negroni

import (
	"net/http"
)

type Static struct {
	Dir http.FileSystem

	Prefix string

	IndexFile string
}

func NewStatic(directory http.FileSystem) *Static { _ = "STUB: not implemented"; return nil }

func (s *Static) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}
