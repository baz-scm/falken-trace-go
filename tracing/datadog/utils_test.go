// Used to store intermediate functions used in actual tests
// WARNING: changing code can result in errors in other test files!

package datadog

import (
	"context"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func testFuncTags() {
	span := tracer.StartSpan("test-span")
	defer span.Finish()

	addCodeTags(span)
}

func testFuncContext(ctx context.Context) {
	span, ctx := tracer.StartSpanFromContext(ctx, "test-span")
	defer span.Finish()

	testInnerFuncContext(ctx)
}

func testInnerFuncContext(ctx context.Context) {
	AddCodeTags(ctx)
}

func testFuncSpanContext(ctx context.Context) {
	span, _ := StartSpanFromContext(ctx, "test-span")
	defer span.Finish()
}

func testFuncSpan() {
	span := StartSpan("test-span")
	defer span.Finish()
}
