package main

import "fmt"

type logger struct {
}

type singleton struct {
	instance *logger
}

func (s *singleton) initialize() {
	if s.instance != nil {
		fmt.Println("Logger already initialized.")
		return
	}

	fmt.Println("Intializing logger")
	s.instance = &logger{}
}

func main() {
	obj := singleton{}

	obj.initialize()
	obj.initialize()
	obj.initialize()
}
