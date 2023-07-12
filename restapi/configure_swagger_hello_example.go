// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ma91n/rfc7807gen/models"
	"github.com/ma91n/rfc7807gen/restapi/operations"
	"github.com/ma91n/rfc7807gen/restapi/operations/example"
)

//go:generate swagger generate server --target ..\..\rfc7807 --name SwaggerHelloExample --spec ..\swagger.yaml --principal interface{}

func configureFlags(api *operations.SwaggerHelloExampleAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SwaggerHelloExampleAPI) http.Handler {
	api.ServeError = errors.ServeError
	api.UseSwaggerUI()

	api.SetDefaultConsumes("application/json")
	api.SetDefaultProduces("application/json")

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.ExampleHelloHandler = example.HelloHandlerFunc(func(params example.HelloParams) middleware.Responder {
		fmt.Println("here")
		return example.NewHelloOK().WithPayload(&models.Hello{
			Message: "test ok",
		})
	})

	api.PreServerShutdown = func() {}
	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return RfcMiddleware(handler)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

type captureResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewCaptureResponseWriter(w http.ResponseWriter) *captureResponseWriter {
	return &captureResponseWriter{w, http.StatusOK}
}

func (lrw *captureResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
	//lrw.ResponseWriter.Header().Set("Content-Type", "application/problem+json")
}

func RfcMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mediaTypeCtx := context.WithValue(r.Context(), int8(2), "application/json") // 2:ctxResponseFormat

		r = r.WithContext(mediaTypeCtx)

		lrw := NewCaptureResponseWriter(w)
		next.ServeHTTP(lrw, r)
	})
}
