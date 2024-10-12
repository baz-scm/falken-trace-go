package opentelemetry

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func TestAddCodeTags(t *testing.T) {
	// given
	ctx := context.Background()
	tracerProvider := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(tracerProvider)

	// when
	span := testFuncContext(ctx)

	attrs, err := getOpentelemetrySpanAttributes(span)
	if err != nil {
		t.Fatal(err)
	}

	// then
	assert.Len(t, attrs, 3)

	file := attrs[0]
	line := attrs[1]
	name := attrs[2]

	assert.Equal(t, "utils_test.go", file.Value.AsString())
	assert.EqualValues(t, 37, line.Value.AsInt64())
	assert.Equal(t, "testInnerFuncContext", name.Value.AsString())
}

func TestFalkenTracer_Start(t *testing.T) {
	// given
	ctx := context.Background()
	tracerProvider := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(tracerProvider)

	// when
	span := testFuncSpan(ctx)

	attrs, err := getOpentelemetrySpanAttributes(span)
	if err != nil {
		t.Fatal(err)
	}

	// then
	assert.Len(t, attrs, 3)

	file := attrs[0]
	line := attrs[1]
	name := attrs[2]

	assert.Equal(t, "utils_test.go", file.Value.AsString())
	assert.EqualValues(t, 42, line.Value.AsInt64())
	assert.Equal(t, "testFuncSpan", name.Value.AsString())
}
