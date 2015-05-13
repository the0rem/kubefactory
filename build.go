package main

import (
	"os"
	"sync"
)

type builder struct {
	buildDest      string
	templateSource string
	envSource      string
}

func (build *builder) Build(distDest, templateSource, envSource string) {

  // Ensure all required directory entries exist
  directories := []string{distDest, templateSource, envSource}

  for _, directory := range directories {
    if _, err exists(directory) && err != nil {
      fmt.Printf("%s", err)
    }  
  }

  // Set variables to point to folders
  templateFiles, err := ioutil.ReadDir(templateSource)
  envFiles, err := ioutil.ReadDir(envSource)
	err := Chdir(distDest)

  // Loop through template files for processing
  for _, file := range templateFiles {

    // If file is a .yaml send for processing
    
    // Add file to dist
    
  }

	wg := new(sync.WaitGroup)
	commands := []string{""}

	for _, str := range commands {
		wg.Add(1)
		go ExecCmd(str, wg)
	}

	wg.Wait()
}

func (build *builder)  {

}

