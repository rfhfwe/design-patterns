package main

import (
	"fmt"
	"sync"
)

type single struct {
}

var once sync.Once

var singInstance *single

func getInstance() *single {
	if singInstance == nil {
		once.Do(func() {
			singInstance = &single{}
		})
	} else {
		fmt.Println("Single instance already created.")
	}
	return singInstance
}
