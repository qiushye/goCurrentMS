package client

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/longjoy/micro-go-book/ch13-seckill/pb"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

func TestUserClientImpl_CheckUser(t *testing.T) {
	client, _ := NewUserClient("user", nil, genTracerAct(nil))

	if response, err := client.CheckUser(context.Background(), nil, &pb.UserRequest{
		Username: "xuan",
		Password: "xuan",
	}); err == nil {
		fmt.Println(response.Result)
	} else {
		fmt.Println(err.Error())
	}
}

func genTracerAct(tracer *zipkin.Tracer) *zipkin.Tracer {
	if tracer != nil {
		return tracer
	}
	zipkinUrl := "http://127.0.0.1:9411/api/v2/spans"
	zipkinRecorder := "localhost:12344"
	// collector, err := zipkin.NewHTTPCollector(zipkinUrl)
	// if err != nil {
	// 	log.Fatalf("zipkin.NewHTTPCollector err: %v", err)
	// }

	reporter := zipkinhttp.NewReporter(zipkinUrl)

	// recorder := zipkin.NewRecorder(collector, false, zipkinRecorder, "user-client")
	zEP, _ := zipkin.NewEndpoint("user-client", zipkinRecorder)
	zipkinTracer, err := zipkin.NewTracer(
		reporter, zipkin.WithLocalEndpoint(zEP), zipkin.WithNoopTracer(false),
	)
	if err != nil {
		log.Fatalf("zipkin.NewTracer err: %v", err)
	}

	// res, err := zipkin.NewTracer(
	// 	reporter, zipkin.ClientServerSameSpan(true),
	// )
	// if err != nil {
	// 	log.Fatalf("zipkin.NewTracer err: %v", err)
	// }
	return zipkinTracer

}
