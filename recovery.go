package negroni

import (
	"net/http"
	"text/template"
)

const (
	NoPrintStackBodyString = "500 Internal Server Error"

	panicText = "PANIC: %s\n%s"
	panicHTML = `<html>
<head><title>PANIC: {{.RecoveredPanic}}</title></head>
<style type="text/css">
html, body {
	font-family: Helvetica, Arial, Sans;
	color: #333333;
	background-color: #ffffff;
	margin: 0px;
}
h1 {
	color: #ffffff;
	background-color: #f14c4c;
	padding: 20px;
	border-bottom: 1px solid #2b3848;
}
.block {
	margin: 2em;
}
.panic-interface {
}

.panic-stack-raw pre {
	padding: 1em;
	background: #f6f8fa;
	border: dashed 1px;
}
.panic-interface-title {
	font-weight: bold;
}
</style>
<body>
<h1>Negroni - PANIC</h1>

<div class="panic-interface block">
	<h3>{{.RequestDescription}}</h3>
	<span class="panic-interface-title">Runtime error:</span> <span class="panic-interface-element">{{.RecoveredPanic}}</span>
</div>

{{ if .Stack }}
<div class="panic-stack-raw block">
	<h3>Runtime Stack</h3>
	<pre>{{.StackAsString}}</pre>
</div>
{{ end }}

</body>
</html>`
	nilRequestMessage = "Request is nil"
)

var panicHTMLTemplate = template.Must(template.New("PanicPage").Parse(panicHTML))

type PanicInformation struct {
	RecoveredPanic interface{}
	Stack          []byte
	Request        *http.Request
}

func (p *PanicInformation) StackAsString() string { _ = "STUB: not implemented"; return "" }

func (p *PanicInformation) RequestDescription() string { _ = "STUB: not implemented"; return "" }

type PanicFormatter interface {
	FormatPanicError(rw http.ResponseWriter, r *http.Request, infos *PanicInformation)
}

type TextPanicFormatter struct{}

func (t *TextPanicFormatter) FormatPanicError(rw http.ResponseWriter, r *http.Request, infos *PanicInformation) {
	_ = "STUB: not implemented"
	return
}

type HTMLPanicFormatter struct{}

func (t *HTMLPanicFormatter) FormatPanicError(rw http.ResponseWriter, r *http.Request, infos *PanicInformation) {
	_ = "STUB: not implemented"
	return
}

type Recovery struct {
	Logger           ALogger
	PrintStack       bool
	LogStack         bool
	PanicHandlerFunc func(*PanicInformation)
	StackAll         bool
	StackSize        int
	Formatter        PanicFormatter

	ErrorHandlerFunc func(interface{})
}

func NewRecovery() *Recovery { _ = "STUB: not implemented"; return nil }

func (rec *Recovery) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}
