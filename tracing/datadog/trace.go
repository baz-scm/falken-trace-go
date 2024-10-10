package datadog

import (
	"context"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

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
