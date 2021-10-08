# Datadog APM Settings

## About Datadog APM​

https://www.datadoghq.com/product/apm/



## Start the tracer

`tracer.Start` starts the tracer with the given set of options.

```go
import (
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

tracer.Start(
	tracer.WithService(config.DataDog.Service),
	tracer.WithEnv(config.DataDog.Env),
	tracer.WithServiceVersion(config.DataDog.Version),
	tracer.WithAnalytics(true),
	tracer.WithLogger(ddlogger),
)
defer tracer.Stop()
```

In the deployment environment, the following three environment variables can be used.
- `DD_SERVICE`
- `DD_ENV`
- `DD_VERSION`

It might be better to be able to switch tracer off in the local environment.



## Start the span for incoming requests

`gintrace.Middleware` do the following.
​
- starts the span for a incoming request.
- set the span to `gin.Context.Request.Context()`
​
```go
import (
  gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)
​
​
func main() {
  app := gin.New()
  app.Use(gintrace.Middleware(fmt.Sprintf("%s-%s", config.DataDog.Service, config.DataDog.Env)))
}
```

When using multiple gin middleware, it is preferable to start span as close
to the front as possible.
- To get accurate measurement data
- The subsequent middlewares can output the logs linked with span
​

In particular, `panic` or runtime errors should be recovered
before returning the process to `gintrace.Middleware`.
This is because `gintrace.Middleware` does some tagging after the `c.Next()`.

​https://github.com/DataDog/dd-trace-go/blob/v1.33.0/contrib/gin-gonic/gin/gintrace.go#L56-L64
​
​
​
## Link the application logs with span
​
By adding the following fields to logs, you can link the application logs with
the span.
​
```go
span, isSpanExist := tracer.SpanFromContext(c.Request.Context())
if isSpanExist {
	zapFields = append(zapFields,
		zap.Uint64("dd.trace_id", span.Context().TraceID()),
		zap.Uint64("dd.span_id", span.Context().SpanID()))
}
​
zapLogger.Info("message", zapFields...)
```

Since it would be difficult to implement the above at all logging points, it might better to set the request logger to context as shown below.

```go
// in pkgLogger package
func SetRequestLogger(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, requestLoggerKey, l)
}

// in SetRequestLogger middleware
c.Request = c.Request.WithContext(pkgLogger.SetRequestLogger(reqCtx, reqLogger))
```
​
​

## Link GORM logs with span
​
The GORM default logger is not designed with Datadog APM in mind.

https://github.com/go-gorm/gorm/blob/v1.21.15/logger/logger.go
​​

An example solution is shown below.
1. Add a customized GORM logger based on the GORM default logger.
2. The customized logger use the request logger.

But this solution will need to follow future changes to the GORM default logger.

​

## Link panic logs with span
​
The `gin.Recovery` middleware doesn't support `zap.Logger`.

https://github.com/gin-gonic/gin/blob/v1.7.4/recovery.go

​
​There is a library available for `zap`, but if you use it as is, you will get
logs that are not linked with span.

https://github.com/gin-contrib/zap/blob/v0.0.1/zap.go#L65
​

An example solution is shown below.
1. Add a customized middleware based on the `gin-contrib/zap` middleware.
2. The customized middleware use the request logger instead of the normal
`zap.Logger`.

But this solution will need to follow future changes to `gin-contrib/zap`.
​


## Create custom span

If you want to make your own measurements, such as API execution for external
systems or complex calculation logic, you can do the following.
​
```go
span, ctx := tracer.StartSpanFromContext(ctx, "child.span.name")
defer span.Finish()

someProcess()
```
