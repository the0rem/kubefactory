package main

import (
	// "github.com/codeskyblue/go-sh"
	"sync"
)

func Supercede() {
	wg := new(sync.WaitGroup)
	commands := []string{""}

	for _, str := range commands {
		wg.Add(1)
		go ExecCmd(str, wg)
	}

	wg.Wait()
}
