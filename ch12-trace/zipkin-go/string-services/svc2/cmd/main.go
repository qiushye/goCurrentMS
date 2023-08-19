//go:build go1.7
// +build go1.7

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/longjoy/micro-go-book/ch12-trace/zipkin-go/string-services/svc2"

	reportHttp "github.com/openzipkin/zipkin-go/reporter/http"

	"github.com/openzipkin/zipkin-go"
	//"github.com/openzipkin-contrib/zipkin-go-opentracing/examples/string-services/svc2"
)

const (
	// Our service name.
	serviceName = "svc2"

	// Host + port of our service.
	hostPort = "127.0.0.1:61002"

	// Endpoint to send Zipkin spans to.
	zipkinHTTPEndpoint = "http://127.0.0.1:9411/api/v1/spans"

	// Debug mode.
	debug = false

	// same span can be set to true for RPC style spans (Zipkin V1) vs Node style (OpenTracing)
	sameSpan = true

	// make Tracer generate 128 bit traceID's for root spans.
	traceID128Bit = true
)

//svc2
func main() {
	// create collector.
	// collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	// if err != nil {
	// 	fmt.Printf("unable to create Zipkin HTTP collector: %+v\n", err)
	// 	os.Exit(-1)
	// }

	// // create recorder.
	// recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)

	// // create tracer.
	// tracer, err := zipkin.NewTracer(
	// 	recorder,
	// 	zipkin.ClientServerSameSpan(sameSpan),
	// 	zipkin.TraceID128Bit(traceID128Bit),
	// )
	// if err != nil {
	// 	fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
	// 	os.Exit(-1)
	// }

	// explicitly set our tracer to be the default tracer.
	// opentracing.InitGlobalTracer(tracer)

	reporter := reportHttp.NewReporter(zipkinHTTPEndpoint)

	// create recorder.
	// recorder := zipkin.New(collector, debug, hostPort, serviceName)

	// create tracer.
	tracer, err := zipkin.NewTracer(
		reporter,
		zipkin.WithTraceID128Bit(traceID128Bit),
	)
	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
		os.Exit(-1)
	}

	// create the service implementation
	service := svc2.NewService()

	// create the HTTP Server Handler for the service
	handler := svc2.NewHTTPHandler(*tracer, service)

	// start the service
	fmt.Printf("Starting %s on %s\n", serviceName, hostPort)
	http.ListenAndServe(hostPort, handler)
}
