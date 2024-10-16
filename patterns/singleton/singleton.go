package singleton

import "sync"

type Singleton struct {
}

var (
	instance *Singleton
	once     sync.Once
)

// GetInstance returns singleton
func GetInstance() *Singleton {

	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
