package builder

import "errors"

const (
	defaultMaxTotal = 10
	defaultMaxIdle  = 9
	defaultMinIdle  = 1
)

type resourcePool struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

type resourcePoolBuilder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func NewResourcePoolBuilder(name string) *resourcePoolBuilder {
	return &resourcePoolBuilder{
		name:     name,
		maxTotal: defaultMaxTotal,
		maxIdle:  defaultMaxIdle,
		minIdle:  defaultMinIdle,
	}
}

func (b *resourcePoolBuilder) MaxTotal(n int) *resourcePoolBuilder {
	b.maxTotal = n
	return b
}

func (b *resourcePoolBuilder) MaxIdle(n int) *resourcePoolBuilder {
	b.maxIdle = n
	return b
}
func (b *resourcePoolBuilder) MinIdle(n int) *resourcePoolBuilder {
	b.minIdle = n
	return b
}

func (b *resourcePoolBuilder) Build() (*resourcePool, error) {
	if b.name == "" {
		return nil, errors.New("name can't be empty")
	}
	if b.minIdle < 0 {

	}
	if b.maxIdle < b.minIdle {

	}
	if b.maxTotal < b.maxIdle {

	}

	return &resourcePool{
		name:     b.name,
		maxTotal: b.maxTotal,
		maxIdle:  b.maxIdle,
		minIdle:  b.minIdle,
	}, nil
}
