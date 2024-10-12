package opentelemetry

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func TestAddCodeTagsPrivate(t *testing.T) {
	// given
	ctx := context.Background()
	tracerProvider := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(tracerProvider)

	// when
	span := testFuncTags(ctx)

	attrs, err := getOpentelemetrySpanAttributes(span)
	if err != nil {
		t.Fatal(err)
	}

	// then
	file := attrs[0]
	line := attrs[1]
	name := attrs[2]

	assert.NotEmpty(t, file.Value.AsString())
	assert.NotEmpty(t, line.Value.AsInt64())
	assert.NotEmpty(t, name.Value.AsString())
}
