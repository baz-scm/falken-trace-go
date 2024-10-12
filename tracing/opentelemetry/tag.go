package opentelemetry

import (
	"github.com/baz-scm/falken-trace-go/internal"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func addCodeTags(span trace.Span) {
	file, line, name := internal.ExtractCodeData()

	if file != "" {
		span.SetAttributes(
			attribute.String("code.filepath", file),
		)
	}
	if line != 0 {
		span.SetAttributes(
			attribute.Int("code.lineno", line),
		)
	}
	if name != "" {
		span.SetAttributes(
			attribute.String("code.func", name),
		)
	}
}
