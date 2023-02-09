package main

import (
	"quizapp/src/package/routers"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go routers.Init()

	wg.Wait()
}
