package main

import (
	"fmt"
	"sync"
)

type logger struct {
}

type singleton struct {
	instance *logger
	lock     *sync.Mutex
}

var once sync.Once

func (s *singleton) initialize() {

	s.lock.Lock()
	defer s.lock.Unlock()

	once.Do(func() {
		fmt.Println("Intializing logger")
		s.instance = &logger{}
	})

	if s.instance != nil {
		fmt.Println("Logger already initialized.")
	}
}

func main() {
	obj := singleton{lock: &sync.Mutex{}}
	size := 5
	var wg sync.WaitGroup

	// you can also add these one at
	// a time if you need to

	wg.Add(size)

	for i := 0; i < size; i++ {
		go func(index int) {
			defer wg.Done()
			fmt.Printf("Trying to init. id:%d\n", index)
			obj.initialize()
		}(i)
	}

	wg.Wait()
}
