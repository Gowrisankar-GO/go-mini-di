package di

type container struct {
	registry *registry
}

func New() *container {
	return &container{
		registry: newRegistry(),
	}
}
