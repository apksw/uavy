package base

type (
	Service interface {
		Worker
	}
)

type (
	BaseService struct {
		*BaseWorker
	}
)

func NewService(name string, tracingLevel string) *BaseService {
	return &BaseService{
		BaseWorker: NewWorker(name, tracingLevel),
	}
}
