package middleware

import (
	"context"
	"net/http"
)

type RequestScope struct {
	req *http.Request
}

func NewRequestScope(req *http.Request) RequestScope {
	return RequestScope{req}
}

func (thing RequestScope) IsUrl(url string) bool {
	return thing.req.RequestURI == url

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
