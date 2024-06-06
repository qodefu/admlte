package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/a-h/templ"
)

type RequestScope struct {
	req    *http.Request
	writer http.ResponseWriter
	stack  map[string][]templ.Component
}

func NewRequestScope(req *http.Request, w http.ResponseWriter) RequestScope {
	return RequestScope{req, w, make(map[string][]templ.Component)}
}

func (thing RequestScope) Context() context.Context {
	return thing.req.Context()
}

type TriggerMsg struct {
	Event   string `json:"-"`
	Message string `json:"message"`
	Tags    string `json:"tags"`
	Data    any    `json:"data,omitempty"`
}

func (thing RequestScope) HxTrigger(events ...TriggerMsg) {
	var data = make(map[string][]TriggerMsg)
	for _, e := range events {
		data[e.Event] = append([]TriggerMsg{}, e)
	}
	bytes, _ := json.Marshal(data)

	thing.Response().Header().Set("HX-Trigger", string(bytes))
}

func (thing RequestScope) QueryParam(key string) string {
	return thing.req.URL.Query().Get(key)
}

func (thing RequestScope) Request() *http.Request {
	return thing.req
}

func (thing RequestScope) Response() http.ResponseWriter {
	return thing.writer
}
func (thing RequestScope) IsUrl(url string) bool {
	return thing.req.RequestURI == url
}

func (thing RequestScope) Push(name string, comp templ.Component) {
	thing.stack[name] = append(thing.stack[name], comp)
}

func (thing RequestScope) Stack(name string) []templ.Component {
	return thing.stack[name]
}

func ReqScope(ctx context.Context) RequestScope {
	ptr := ctx.Value(req_scope{}).(*RequestScope)
	return *ptr
}

func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b

}
