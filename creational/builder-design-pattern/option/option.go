package option

import (
	"errors"
)

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

type ResourcePoolOption struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolOptFunc func(option *ResourcePoolOption)

func NewResourcePool(name string, opts ...ResourcePoolOptFunc) (*resourcePool, error) {
	if name == "" {
		return nil, errors.New("name can't be empty")
	}

	option := &ResourcePoolOption{
		maxTotal: defaultMaxTotal,
		maxIdle:  defaultMaxIdle,
		minIdle:  defaultMinIdle,
	}

	for _, opt := range opts {
		opt(option)
	}

	if option.minIdle < 0 {

	}
	if option.maxIdle < option.minIdle {

	}
	if option.maxTotal < option.maxIdle {

	}

	return &resourcePool{
		name:     name,
		maxTotal: option.maxTotal,
		maxIdle:  option.maxIdle,
		minIdle:  option.minIdle,
	}, nil
}
