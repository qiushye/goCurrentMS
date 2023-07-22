package endpoint

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/longjoy/micro-go-book/ch6-discovery/stringService/service"
)

type StringEndpoints struct {
	StringEndpoint endpoint.Endpoint
	HealthEndpoint endpoint.Endpoint
}

type StringRequest struct {
	RequestType string
	A           string
	B           string
}

type StringResponse struct {
	Res string
}

type HealthResponse struct {
	Status bool
}

func MakeStringEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(StringRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request")
		}

		var res string
		switch req.RequestType {
		case "concat":
			res, err = svc.Concat(req.A, req.B)
			if err != nil {
				return nil, err
			}
		case "diff":
			res, err = svc.Diff(req.A, req.B)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("invalid request_type")
		}

		return StringResponse{
			Res: res,
		}, nil
	}

}

func MakeHealthEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status := svc.Health()
		return HealthResponse{
			Status: status,
		}, nil
	}
}
