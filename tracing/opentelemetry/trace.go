package opentelemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// AddCodeTags attaches code metadata to the current span.
// It should only be used, when the span is automatically created by an OpenTelemetry integration,
// like for web frameworks handler functions
func AddCodeTags(ctx context.Context) {
	span := trace.SpanFromContext(ctx)
	if !span.SpanContext().IsValid() {
		// make sure it is a real Span
		return
	}

	addCodeTags(span)
}

// Tracer wraps OpenTelemetry's otel.Tracer and returns a FalkenTracer instance to work as a proxy.
func Tracer(name string, opts ...trace.TracerOption) FalkenTracer {
	innerTracer := otel.Tracer(name, opts...)
	return FalkenTracer{innerTracer}
}

type FalkenTracer struct {
	innerTracer trace.Tracer
}

// Start wraps OpenTelemetry's tracer.Start and attaches code metadata to the span.
func (t FalkenTracer) Start(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	ctx, span := t.innerTracer.Start(ctx, name, opts...)

	addCodeTags(span)

	return ctx, span
}
