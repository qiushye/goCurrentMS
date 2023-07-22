package service

type Service interface {
	Health() bool
	Concat(a, b string) (string, error)
	Diff(a, b string) (string, error)
}

type StringService struct{}

func (s StringService) Concat(a, b string) (string, error) {
	return "concat", nil
}

func (s StringService) Diff(a, b string) (string, error) {
	return "diff", nil
}

func (s StringService) Health() bool {
	return true
}

type ServiceMiddleware func(Service) Service
