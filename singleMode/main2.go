package singleton2

import "sync"

type singleton struct{}

var instance *singleton

var mu sync.Mutex

func GetInstance() *instance {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &singleton{} // unnecessary locking if instance already created
	}
	return instance
}
