package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/longjoy/micro-go-book/ch6-discovery/discover"
	"github.com/longjoy/micro-go-book/ch6-discovery/string-service/config"
	"github.com/longjoy/micro-go-book/ch6-discovery/stringService/endpoint"
	"github.com/longjoy/micro-go-book/ch6-discovery/stringService/plugins"
	"github.com/longjoy/micro-go-book/ch6-discovery/stringService/service"
	"github.com/longjoy/micro-go-book/ch6-discovery/stringService/transport"
)

func main() {
	var (
		servicePort = flag.Int("servicePort", 2000, "service port")
		serviceHost = flag.String("serviceHost", "127.0.0.1", "service host")
		consulPort  = flag.Int("consulPort", 8500, "service port")
		consulHost  = flag.String("consulHost", "127.0.0.1", "consul host")
		serviceName = flag.String("serviceName", "stringService", "service name")

		ctx             = context.Background()
		errChan         = make(chan error)
		discoveryClient discover.DiscoveryClient
		instanceId      = uuid.NewString()
		err             error
		meta            = map[string]string{}
		logger          = log.Default()
	)

	flag.Parse()

	discoveryClient, err = discover.NewKitDiscoverClient(*consulHost, *consulPort)
	if err != nil {
		log.Fatal(err)
	}

	var svc = plugins.LoggingMiddleware(config.KitLogger)(service.StringService{})
	stringEp := endpoint.MakeStringEndpoint(svc)
	healthEp := endpoint.MakeHealthEndpoint(svc)
	eps := endpoint.StringEndpoints{
		StringEndpoint: stringEp,
		HealthEndpoint: healthEp,
	}

	r := transport.MakeHttpHandler(ctx, eps, config.KitLogger)
	go func() {

		if !discoveryClient.Register(*serviceName, instanceId, "/v1/health", *serviceHost, *servicePort, meta, logger) {
			log.Fatal("register failed")
		}
		errChan <- http.ListenAndServe(fmt.Sprintf("%s:%d", *serviceHost, *servicePort), r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	err = <-errChan
	discoveryClient.DeRegister(instanceId, logger)
	logger.Fatalf("exit %s", err)
	// discoveryClient.Register(*serviceName, instanceId, "", *serviceHost, *servicePort, meta, logger)

}
