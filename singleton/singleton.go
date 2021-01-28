package singleton

import (
	"sync"
)

type singleton struct{}

var single *singleton

var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		single = &singleton{}
	})
	return single
}
