package singleton3

import (
	"sync"
)

type singleton struct{}

var instance *singleton

var once sync.Once

func GetInstance() *instance {
	once.Do(func() {
		instance = &instance{}
	})

	return instance
}
