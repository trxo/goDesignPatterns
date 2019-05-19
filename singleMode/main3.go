package singleton3

import "sync"

type singleton struct{}

var instance *singleton

var mu sync.Mutex

func GetInstance() *instance {
	if instance == nil { // Not yet perfect. since it's not fully atomic
		mu.Lock()
		defer mu.Unlock()

		if instance == nil {
			instance = &singleton{} // unnecessary locking if instance already created
		}
	}
	return instance
}
