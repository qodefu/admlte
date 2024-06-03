package middleware

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

type RequestScope struct {
	req   *http.Request
	stack map[string][]templ.Component
}

func NewRequestScope(req *http.Request) RequestScope {
	return RequestScope{req, make(map[string][]templ.Component)}
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
	return ctx.Value(REQ_SCOPE).(RequestScope)
}

func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b

}
