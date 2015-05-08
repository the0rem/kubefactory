package main

import (

  "sync"

)

type dirSync struct {
  from, to string
}

func (link *dirSync) configure() {
  
  link.from = "yomama"
  link.to = "mymama"
  
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