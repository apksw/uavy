// NOTE: WIP, a lot in common with Service
package base

type (
	Worker interface {
		Name() string
		Init() bool  // NOTE: Better return an error?
		Start() bool // NOTE: Better return an error?
	}
)

type (
	BaseWorker struct {
		name     string
		didInit  bool
		didStart bool
	}
)

func NewWorker(name string) *BaseWorker {
	name = genName(name, "worker")

	return &BaseWorker{
		name: name,
	}
}

func (bw BaseWorker) Name() string {
	return bw.name
}

func (bw BaseWorker) SetName(name string) {
	bw.name = name
}

func (bw BaseWorker) Init() bool {
	return bw.didInit
}

func (bw BaseWorker) Start() bool {
	return bw.didStart
}
