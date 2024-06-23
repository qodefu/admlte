package components

import (
	"context"
	"fmt"
	"goth/internal/middleware"
	"io"

	"github.com/a-h/templ"
)

func PushScript(name string, script templ.ComponentScript) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		reqscope := middleware.ReqScope(ctx)
		// childcomp := templ.GetChildren(ctx)
		reqscope.Push(name, script)
		_, err := io.WriteString(w, "")
		return err
	})
}

func Push(name string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		reqscope := middleware.ReqScope(ctx)
		childcomp := templ.GetChildren(ctx)
		reqscope.Push(name, childcomp)
		_, err := io.WriteString(w, "")
		return err
	})
}

func Stack(name string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		reqscope := middleware.ReqScope(ctx)
		for _, c := range reqscope.Stack(name) {
			c.Render(ctx, w)
		}

		return nil
	})
}

func Js(script string, a ...any) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, fmt.Sprintf(script, a...))
		return err
	})
}
