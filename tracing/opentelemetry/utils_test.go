// Used to store intermediate functions used in actual tests
// WARNING: changing code can result in errors in other test files!

package opentelemetry

import (
	"context"
	"errors"
	"reflect"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func testFuncTags(ctx context.Context) trace.Span {
	tracer := otel.Tracer("test-tracer")
	_, span := tracer.Start(ctx, "test-span")
	defer span.End()

	addCodeTags(span)

	return span
}

func testFuncContext(ctx context.Context) trace.Span {
	tracer := otel.Tracer("test-tracer")
	ctx, span := tracer.Start(ctx, "test-span")
	defer span.End()

	testInnerFuncContext(ctx)

	return span
}

func testInnerFuncContext(ctx context.Context) {
	AddCodeTags(ctx)
}

func testFuncSpan(ctx context.Context) trace.Span {
	tracer := Tracer("test-tracer")
	_, span := tracer.Start(ctx, "test-span")
	defer span.End()

	return span
}

func getOpentelemetrySpanAttributes(span trace.Span) ([]attribute.KeyValue, error) {
	spanValue := reflect.ValueOf(span)

	// Accessing the Span's `Attributes()` method
	method := spanValue.MethodByName("Attributes")
	if !method.IsValid() {
		return nil, errors.New("Method 'Attributes' not found")
	}

	// Call the method
	results := method.Call(nil)

	attributes := results[0].Interface().([]attribute.KeyValue)
	return attributes, nil
}
