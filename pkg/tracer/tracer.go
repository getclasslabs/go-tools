package tracer

import "github.com/opentracing/opentracing-go"

func TraceIt(i *Infos, name string) opentracing.Span {
	return i.Tracer.StartSpan(name, opentracing.ChildOf(i.Span.Context()))
}
