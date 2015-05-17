package main

import (
	"bufio"
	"fmt"
	"github.com/codeskyblue/go-sh"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

/**
 * Return an interactive bash terminal for the given pod and container
 * @param {string} pod
 * @param {string} container
 */
func Enter(pod, container string) {

	var podName string
	var containerName string

	// Look for pods starting with the given string
	// TODO: Test what field awk needs to get from each line
	msg, err := sh.Command("kubectl", "get", "pods").Command("grep", pod).Output()

	if err != nil {
		color.Red("Couldn't find the given pod")
		os.Exit(1)
	}

	// Break up results based on spaces
	pods := strings.Split(string(msg[:]), " ")

	/// If multiple pods are found, ask the user which one
	if len(pods) > 1 {

		// Ask user to select the right container
		color.White("We found more than one pod matching your name, which one would you like to choose?")

		for key, pod := range pods {
			color.White(fmt.Sprintf("[%d] %s", key, pod))
		}

		consolereader := bufio.NewReader(os.Stdin)
		input, err := consolereader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		podIndex, _ := strconv.Atoi(input)
		podName = pods[podIndex]

	} else {

		/// If only one pod is found, save it
		podName = pods[0]

	}

	// Look for containers in the pod within the given pod
	// TODO: Test what field awk needs to get from each line
	msg, err = sh.Command("kubectl", "get", "containers").Command("grep", podName).Command("grep", container).Output()

	if err != nil {
		color.Red("Couldn't find the given controller")
		os.Exit(1)
	}

	// Break up results based on spaces
	containers := strings.Split(string(msg[:]), " ")

	// If multiple containers are found, ask the user which one
	if len(containers) > 1 {

		// Ask user to select the right container
		color.White("We found more than one container matching your name, which one would you like to choose?")

		for key, container := range containers {
			color.White(fmt.Sprintf("[%d] %s", key, container))
		}

		consolereader := bufio.NewReader(os.Stdin)
		input, err := consolereader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		containerIndex, _ := strconv.Atoi(input)
		containerName = containers[containerIndex]

	} else {

		// If only one container is found, save it
		containerName = containers[0]

	}

	// Get the absolute path to command
	binary, lookErr := exec.LookPath("kubectl")

	if lookErr != nil {
		panic(lookErr)
	}

	// Set arguments for an interactive bash terminal
	args := []string{"exec", "-c", containerName, "-p", podName, "-i", "-t", "--", "bash", "-il"}

	// Add environment variables
	env := os.Environ()

	// Run the command and give to user
	execErr := syscall.Exec(binary, args, env)

	// Handle error
	if execErr != nil {
		panic(execErr)
	}

}
