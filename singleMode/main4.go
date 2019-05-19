package singleton3

import "sync/atomic"

type singleton struct{}

var instance *singleton

var initialized uint32

func GetInstance() *instance {
	if atomic.LoadInt32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()
	if initialized == 0 {
		instance = &singleton{}
		atomic.StoreInt32(&initialized, 2)
	}

	return instance
}
