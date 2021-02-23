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

func NewService(name string) *BaseService {
	return &BaseService{
		BaseWorker: NewWorker(name),
	}
}
