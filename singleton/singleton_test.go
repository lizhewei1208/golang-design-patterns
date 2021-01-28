package singleton

import (
	"sync"
	"testing"
)

var wg sync.WaitGroup

const count = 100

func TestSingleton(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()
	if instance1 != instance2 {
		t.Fatal("instance is not equal")
	}
}

func TestSingleton2(t *testing.T) {
	start := make(chan int)
	instances := [count]*singleton{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {

			//协程阻塞，等待channel被关闭才能继续运行
			<-start
			instances[i] = GetInstance()
			wg.Done()
		}(i)
	}

	//关闭channel，所有协程同时开始运行，实现并行(parallel)
	close(start)

	wg.Wait()
	for i := 1; i < count; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
