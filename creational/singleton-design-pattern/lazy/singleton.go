package lazy

import (
	"sync"
)

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstanceDirect() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func GetInstanceCheckNil() *singleton {
	if instance == nil {
		once.Do(func() {
			instance = &singleton{}
		})
	}
	return instance
}
