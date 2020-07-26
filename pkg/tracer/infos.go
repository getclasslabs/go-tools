package tracer

import "github.com/opentracing/opentracing-go"

type Infos struct {
	Span   opentracing.Span
	Tracer opentracing.Tracer
}

func (i *Infos) LogError(err error) {
	i.Span.SetTag("error", true)
	i.Span.SetTag("errorMsg", err.Error())
}
