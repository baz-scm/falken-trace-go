package datadog

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/mocktracer"
)

func TestAddCodeTags(t *testing.T) {
	// given
	mt := mocktracer.Start()
	defer mt.Stop()
	ctx := context.Background()

	// when
	testFuncContext(ctx)

	// then
	spans := mt.FinishedSpans()

	assert.Len(t, spans, 1)
	tags := spans[0].Tags()

	assert.GreaterOrEqual(t, len(tags), 3)

	assert.Equal(t, "utils_test.go", tags["code.filepath"])
	assert.Equal(t, "27", tags["code.lineno"])
	assert.Equal(t, "testInnerFuncContext", tags["code.func"])
}

func TestStartSpanFromContext(t *testing.T) {
	// given
	mt := mocktracer.Start()
	defer mt.Stop()
	ctx := context.Background()

	// when
	testFuncSpanContext(ctx)

	// then
	spans := mt.FinishedSpans()

	assert.Len(t, spans, 1)
	tags := spans[0].Tags()

	assert.GreaterOrEqual(t, len(tags), 3)

	assert.Equal(t, "utils_test.go", tags["code.filepath"])
	assert.Equal(t, "31", tags["code.lineno"])
	assert.Equal(t, "testFuncSpanContext", tags["code.func"])
}

func TestStartSpan(t *testing.T) {
	// given
	mt := mocktracer.Start()
	defer mt.Stop()

	// when
	testFuncSpan()

	// then
	spans := mt.FinishedSpans()

	assert.Len(t, spans, 1)
	tags := spans[0].Tags()

	assert.GreaterOrEqual(t, len(tags), 3)

	assert.Equal(t, "utils_test.go", tags["code.filepath"])
	assert.Equal(t, "36", tags["code.lineno"])
	assert.Equal(t, "testFuncSpan", tags["code.func"])
}
