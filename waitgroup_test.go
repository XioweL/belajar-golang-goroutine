package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1) // Pindahkan Add(1) ke sini
		go RunAsynchronous(group)
	}

	group.Wait() // Tunggu semua goroutine selesai
	fmt.Println("Selesai")
}
