package main

import (

  "sync"

)

func Enter() {
  
  commands := []string{"kubectl -c " +  + " -p " +  + " -it"}

  wg := new(sync.WaitGroup)
    
  for _, str := range commands {
    wg.Add(1)
    go ExecCmd(str, wg)
  }
  
  wg.Wait()
}