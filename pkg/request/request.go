package request

import (
	"context"
	"fmt"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

const ContextKey = "infos"

type controller func(http.ResponseWriter, *http.Request)

func PreRequest(h controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := opentracing.GlobalTracer()
		tCtx, _ := t.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))

		span := t.StartSpan(fmt.Sprintf("%s %s", r.Method, r.URL.Path), ext.RPCServerOption(tCtx))
		defer span.Finish()
		ext.HTTPMethod.Set(span, r.Method)
		ext.HTTPUrl.Set(span, r.URL.String())

		ctx := context.WithValue(r.Context(), ContextKey, &tracer.Infos{Span: span, Tracer: t})

		w.Header().Set("Content-Type", "application/json")
		h(w, r.WithContext(ctx))
	}
}
