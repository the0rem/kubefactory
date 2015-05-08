package main

import (

  "sync"

)

func Build() {
  wg := new(sync.WaitGroup)
  commands := []string{""}
    
  for _, str := range commands {
    wg.Add(1)
    go ExecCmd(str, wg)
  }
  
  wg.Wait()
}