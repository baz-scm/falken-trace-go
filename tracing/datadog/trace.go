package datadog

import (
	"context"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// AddCodeTags attaches code metadata to the current span.
// It should only be used, when the span is automatically created by Datadog,
// like for web frameworks handler functions
func AddCodeTags(ctx context.Context) {
	span, found := tracer.SpanFromContext(ctx)
	if !found {
		return
	}

	addCodeTags(span)
}

// StartSpanFromContext wraps Datadog's tracer.StartSpanFromContext and attaches code metadata to the span.
func StartSpanFromContext(ctx context.Context, operationName string, opts ...tracer.StartSpanOption) (tracer.Span, context.Context) {
	span, ctx := tracer.StartSpanFromContext(ctx, operationName, opts...)

	addCodeTags(span)

	return span, ctx
}

// StartSpan wraps Datadog's tracer.StartSpan and attaches code metadata to the span.
func StartSpan(operationName string, opts ...tracer.StartSpanOption) tracer.Span {
	span := tracer.StartSpan(operationName, opts...)

	addCodeTags(span)

	return span
}
