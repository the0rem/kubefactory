package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type generatorOptions struct {
	templateSource string
	envSource      string
}

func (options *generatorOptions) Generate(templateSource, envSource) {

	consolereader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name : ")

	input, err := consolereader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(input)

	// wg := new(sync.WaitGroup)
	// commands := []string{""}

	// for _, str := range commands {
	// 	wg.Add(1)
	// 	go ExecCmd(str, wg)
	// }

	// wg.Wait()
}

func start() {

	consolereader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the template generator. To get started, answer from the options below:")
	fmt.Println("[0] Create a new file")
	fmt.Println("[1] Edit an existing file")

	input, err := consolereader.ReadString('\n')

}
