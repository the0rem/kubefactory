package main

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

func ExecCmd(cmd string, wg *sync.WaitGroup) {

	fmt.Println("command is ", cmd)

	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Printf("%s", out)

	wg.Done() // Need to signal to waitgroup that this goroutine is done

}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
