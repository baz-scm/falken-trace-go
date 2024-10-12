package datadog

import (
	"strconv"

	"github.com/baz-scm/falken-trace-go/internal"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func addCodeTags(span tracer.Span) {
	file, line, name := internal.ExtractCodeData()

	if file != "" {
		span.SetTag("code.filepath", file)
	}
	if line != 0 {
		span.SetTag("code.lineno", strconv.Itoa(line))
	}
	if name != "" {
		span.SetTag("code.func", name)
	}
}

// It doesn't add the tags recursively,
// but rather leverages the recursive variant of ExtractCodeData (internal.ExtractCodeDataRecursive)
func addCodeTagsRecursive(span tracer.Span) {
	file, line, name := internal.ExtractCodeDataRecursive()

	if file != "" {
		span.SetTag("code.filepath", file)
	}
	if line != 0 {
		span.SetTag("code.lineno", line)
	}
	if name != "" {
		span.SetTag("code.func", name)
	}
}
