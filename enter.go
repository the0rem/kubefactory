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
func (app *appConfig) Enter(pod, container string) {

	var podName string
	var containerName string
	var podIndex int
	var containerIndex int

	// Look for pods starting with the given string
	msg, err := sh.Command("kubectl", "--kubeconfig="+app.envFile, "get", "pods").Command("grep", pod).Output()

	if err != nil {
		color.Red(fmt.Sprintf("Couldn't find the given pod %s", err))
		os.Exit(1)
	}

	// Break up results based on spaces
	pods := strings.Split(string(msg[:]), "\n")

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

		podIndex, _ = strconv.Atoi(input)

	}

	podFields := strings.Fields(pods[podIndex])

	podName = podFields[0]

	// Break up results based on commas
	containers := strings.Split(podFields[2], ",")

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

		containerIndex, _ = strconv.Atoi(input)

	}

	containerName = containers[containerIndex]

	color.White(fmt.Sprintf("Set pod as: %s, container as %s", podName, containerName))

	// Get the absolute path to command
	binary, lookErr := exec.LookPath("kubectl")

	if lookErr != nil {
		panic(lookErr)
	}

	// Set arguments for an interactive bash terminal
	args := []string{"--kubeconfig=" + app.envFile, "exec", "-c", containerName, "-p", podName, "-i", "-t", "--", "bash", "-il"}

	// Add environment variables
	env := os.Environ()

	// Run the command and give to user
	execErr := syscall.Exec(binary, args, env)

	// Handle error
	if execErr != nil {
		panic(execErr)
	}

}
