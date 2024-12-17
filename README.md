<div align="center">
   <img align="center" width="128px" src="https://avatars.githubusercontent.com/u/140384842?s=200&v=4" />
   <h1 align="center"><b>falken-trace-go</b></h1>
   <p align="center">
      Enhance OpenTelemetry with pinpointed code-level observability for Python.
      <br />
      <a href="https://github.com/baz-scm/"><strong>Baz on GitHub »</strong></a>
      <br />
      <br />
      <b>Install</b>
      <br />
      <code>go get github.com/baz-scm/falken-trace-go</code>
   </p>
</div>

---

## 🚀 What is Falken Trace?

Falken Trace extends OpenTelemetry for Python by pinpointing **file names, function names**, and **line numbers** that generate spans. It addresses gaps in default observability implementations, making tracing faster and more actionable.

Discovered while building [Baz](https://baz.co), Falken Trace streamlines debugging and enhances insights into codebase behaviors.

---

## Install

```shell
go get github.com/baz-scm/falken-trace-go
```


## Usage
#### With dd-trace-go
Just replace the `tracer.StartSpanFromContext` calls with `falken.StartSpanFromContext`.

```go
import (
    "context"

-   "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
+   falken "github.com/baz-scm/falken-trace-go/tracing/datadog"
)

func doSomething(ctx context.Context) {
-   span, _ := tracer.StartSpanFromContext(ctx, "doSomething")
+   span, _ := falken.StartSpanFromContext(ctx, "doSomething")
    defer span.Finish()
    
    ...
}
```


#### With otel
Just replace the `otel.Tracer` calls with `falken.Tracer`.

```go
import (
    "context"

-   "go.opentelemetry.io/otel"
+   falken "github.com/baz-scm/falken-trace-go/tracing/opentelemetry"
)

var (
-   tracer = otel.Tracer("tracer")
+   tracer = falken.Tracer("tracer")
)

func doSomething(ctx context.Context) {
    _, span := falken.Start(ctx, "doSomething")
    defer span.End()
    
    ...
}
```

# 🔗 Learn More
Python library: https://github.com/baz-scm/falken-trace-py

Blog post: [Extending OpenTelemetry to Pinpoint Code Elements](https://baz.co/resources/extending-opentelemetry-to-pinpoint-code-elements-our-journey-to-close-the-gap)
