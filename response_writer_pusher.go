//go:build go1.8
// +build go1.8

package negroni

import (
	"net/http"
)

func (rw *responseWriter) Push(target string, opts *http.PushOptions) error {
	_ = "STUB: not implemented"
	return nil
}
