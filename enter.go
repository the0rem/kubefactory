package main

import (

  "sync"

)

func Enter() {
  
  pod, container := "", ""


  commands := []string{"kubectl -c " + container + " -p " + pod + " -it"}

  wg := new(sync.WaitGroup)
    
  for _, str := range commands {
    wg.Add(1)
    go ExecCmd(str, wg)
  }
  
  wg.Wait()
}