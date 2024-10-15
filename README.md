# falken-trace-go

## Getting Started

### Requirements

- Go 1.22+

### Install

```shell
go get github.com/baz-scm/falken-trace-go
```

### Usage
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
