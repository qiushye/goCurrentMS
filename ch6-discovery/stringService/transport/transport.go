package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/longjoy/micro-go-book/ch6-discovery/stringService/endpoint"
)

func MakeHttpHandler(ctx context.Context, endpoints endpoint.StringEndpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	// r.Handle("/v1/:a/:b/:request_type",endpoint.M)
	r.Methods("POST").Path("/v1/{a}/{b}/{request_type}").Handler(kithttp.NewServer(
		endpoints.StringEndpoint,
		decodeStringRequest,
		encodeResp,
		options...,
	))

	r.Methods("GET").Path("/v1/health").Handler(kithttp.NewServer(
		endpoints.HealthEndpoint,
		decodeHealthRequest,
		encodeResp,
		options...,
	))

	return r
}

func decodeStringRequest(ctx context.Context, req *http.Request) (request interface{}, err error) {
	vars := mux.Vars(req)
	requestType, ok := vars["request_type"]
	if !ok {
		return nil, fmt.Errorf("no type")
	}

	a, ok := vars["a"]
	if !ok {
		return nil, fmt.Errorf("no a")
	}

	b, ok := vars["b"]
	if !ok {
		return nil, fmt.Errorf("no b")
	}

	return endpoint.StringRequest{
		RequestType: requestType,
		A:           a,
		B:           b,
	}, nil

}

func decodeHealthRequest(ctx context.Context, req *http.Request) (request interface{}, err error) {
	return
}

func encodeResp(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	err := json.NewEncoder(w).Encode(resp)
	return err
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
