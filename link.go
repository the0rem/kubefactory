package main

import (
	"sync"
)

/**
 *
 */
type dirSync struct {
	from, to string
}

/**
 * [func description]
 * @param  {[type]} link *dirSync)     configure(from, to string [description]
 * @return {[type]}      [description]
 */
func (link *dirSync) configure(from, to string) {

	link.from = from
	link.to = to

}

func (link *dirSync) List() {
	wg := new(sync.WaitGroup)
	commands := []string{""}

	for _, str := range commands {
		wg.Add(1)
		go ExecCmd(str, wg)
	}

	wg.Wait()
}

func (link *dirSync) Remove() {
	wg := new(sync.WaitGroup)
	commands := []string{""}

	for _, str := range commands {
		wg.Add(1)
		go ExecCmd(str, wg)
	}

	wg.Wait()
}

func (link *dirSync) Add() {
	wg := new(sync.WaitGroup)
	commands := []string{""}

	for _, str := range commands {
		wg.Add(1)
		go ExecCmd(str, wg)
	}

	wg.Wait()
}
