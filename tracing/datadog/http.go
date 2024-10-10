package datadog

import (
	"net/http"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
)

// WrapClient wraps Datadog's httptrace.WrapClient and attaches code metadata to the spans.
func WrapClient(c *http.Client, opts ...httptrace.RoundTripperOption) *http.Client {
	opts = append(opts, httptrace.WithBefore(func(req *http.Request, span ddtrace.Span) {
		addCodeTagsRecursive(span)
	}))
	return httptrace.WrapClient(c, opts...)
}
