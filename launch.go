package main

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"github.com/fatih/color"
	"io/ioutil"
	"path/filepath"
)

/**
 * [func description]
 * @param  {[type]} app *appConfig)   Launch(distDir string [description]
 * @return {[type]}     [description]
 */
func (app *appConfig) Launch(distDir string) {

	app.LaunchTemplatesFromFolder(distDir)
	app.GetClusterStatus()

}

/**
 * [func description]
 * @param  {[type]} app *appConfig)   LaunchTemplatesFromFolder(sourceDir string [description]
 * @return {[type]}        [description]
 */
func (app *appConfig) LaunchTemplatesFromFolder(sourceDir string) {

	// Get list of files in sourceDir
	files, err := ioutil.ReadDir(sourceDir)

	// Check for error finding files
	if err != nil {
		color.Red(fmt.Sprintf("%s", err))
	}

	// Deploy generated template files to kubernetes endpoint
	for _, file := range files {

		filename := sourceDir + file.Name()
		// If file is dir, recurse function
		if file.IsDir() {

			app.LaunchTemplatesFromFolder(filename + "/")
			continue

		}

		// Only handle file if a yaml
		if filepath.Ext(file.Name()) != ".yaml" {

			color.Yellow(fmt.Sprintf("File %s is not a template", filename))
			continue

		}

		// Run create command on specified server
		msg, err := sh.Command("kubectl", "--kubeconfig="+app.envFile, "create", "-f", filename).Output()

		if err != nil {
			color.Red(fmt.Sprintf("%s", err))
		}

		color.White(string(msg[:]))

	}

}

/**
 * [func description]
 * @param  {[type]} config *appConfig)   GetClusterStatus( [description]
 * @return {[type]}        [description]
 */
func (app *appConfig) GetClusterStatus() {

	var msg []byte

	msg, _ = sh.Command("kubectl", "--kubeconfig="+app.envFile, "cluster-info").Output()
	color.Green("Cluster Info")
	color.White(string(msg[:]))

	msg, _ = sh.Command("kubectl", "--kubeconfig="+app.envFile, "get", "minions").Output()
	color.Green("Minions")
	color.White(string(msg[:]))

	msg, _ = sh.Command("kubectl", "--kubeconfig="+app.envFile, "get", "services").Output()
	color.Green("Services")
	color.White(string(msg[:]))

	msg, _ = sh.Command("kubectl", "--kubeconfig="+app.envFile, "get", "replicationcontrollers").Output()
	color.Green("Replication Controllers")
	color.White(string(msg[:]))

	msg, _ = sh.Command("kubectl", "--kubeconfig="+app.envFile, "get", "pods").Output()
	color.Green("Pods")
	color.White(string(msg[:]))

}
