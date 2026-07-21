package negroni

import (
	"bufio"
	"net"
)

const (
	flusher = 1 << iota
	hijacker
	closeNotifier
)

type (
	flusherFeature       struct{ *responseWriter }
	hijackerFeature      struct{ *responseWriter }
	closeNotifierFeature struct{ *responseWriter }
)

func (f flusherFeature) Flush() { _ = "STUB: not implemented"; return }

func (f hijackerFeature) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

func (f closeNotifierFeature) CloseNotify() <-chan bool { _ = "STUB: not implemented"; return nil }

var featurePicker = make([]func(writer *responseWriter) ResponseWriter, 8)

func initFeaturePicker() { _ = "STUB: not implemented"; return }

func wrapFeature(w *responseWriter) ResponseWriter {
	_ = "STUB: not implemented"
	return *new(ResponseWriter)
}
