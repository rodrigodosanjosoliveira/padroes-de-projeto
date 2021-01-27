package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Criando a nstância agora.")
			singleInstance = &single{}
		} else {
			fmt.Println("Instância já foi criada.")
		}
	} else {
		fmt.Println("Instância já foi criada.")
	}
	return singleInstance
}
