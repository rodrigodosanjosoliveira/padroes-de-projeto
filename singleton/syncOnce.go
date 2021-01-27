package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleOnce struct {
}

var singleInstanceOnce *singleOnce

func getInstanceOnce() *singleOnce {
	if singleInstanceOnce == nil {
		once.Do(
			func() {
				fmt.Println("Criando a instância agora.")
				singleInstanceOnce = &singleOnce{}
			})
	} else {
		fmt.Println("Instância já foi criada.")
	}

	return singleInstanceOnce
}
