package main

import (
	"sync"
	// "github.com/codeskyblue/go-sh"
)

func Enter() {

	pod, container := "", ""

	// Search for a pod/replication controller with the entered name
	// Get a container in the found item

	commands := []string{"kubectl -c " + container + " -p " + pod + " -it"}

	wg := new(sync.WaitGroup)

	for _, str := range commands {
		wg.Add(1)
		go ExecCmd(str, wg)
	}

	wg.Wait()
}
