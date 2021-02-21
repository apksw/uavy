package base

type (
	Service interface {
		Name() string
		Service() Service
		SetService(Service)
	}
)

type (
	BaseService struct {
		name    string
		service Service
	}
)

func NewService(name string) *BaseService {
	name = genName(name, "service")

	return &BaseService{
		name: name,
	}
}

func (bs BaseService) Name() string {
	return bs.name
}

func (bs BaseService) Service() Service {
	return bs.service
}

func (bs BaseService) SetService(s Service) {
	bs.service = s
}
