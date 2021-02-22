package base

type (
	Service interface {
		Name() string
	}
)

type (
	BaseService struct {
		name string
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

func (bs BaseService) SetName(name string) {
	bs.name = name
}
