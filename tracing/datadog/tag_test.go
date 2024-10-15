package datadog

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/mocktracer"
)

func TestAddCodeTagsPrivate(t *testing.T) {
	// given
	mt := mocktracer.Start()
	defer mt.Stop()

	// when
	testFuncTags()

	// then
	spans := mt.FinishedSpans()

	assert.Len(t, spans, 1)
	tags := spans[0].Tags()

	assert.GreaterOrEqual(t, len(tags), 3)

	assert.NotEmpty(t, tags["code.filepath"])
	assert.NotEmpty(t, tags["code.lineno"])
	assert.NotEmpty(t, tags["code.func"])
}
