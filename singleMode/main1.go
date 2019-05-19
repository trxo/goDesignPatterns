package singleton1

// singleton ...
type singleton struct{}

// private
var instance *singleton

// GetInstance ...
func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{} //no thread safe
	}
	return instance
}
