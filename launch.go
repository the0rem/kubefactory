package main

import (
	"github.com/codeskyblue/go-sh"
	"github.com/fatih/color"
	"path/filepath"
	// "github.com/joho/godotenv/autoload"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
)

type launchParams struct {
	distDir          string
	environment      string
	envFile          string
	sourceFilePrefix string
}

func (params *launchParams) Configure(distDir, environment, envFile string) {

	params.distDir = distDir
	params.environment = environment
	params.envFile = envFile

}

func (params *launchParams) Launch() {

	var msg []byte

	// Get list of files in dist
	files, err := ioutil.ReadDir(params.distDir)

	// Check for error finding files
	if err != nil {
		color.Red(fmt.Sprintf("%s", err))
	}

	// Get env variables
	err = godotenv.Load(params.envFile)

	if err != nil {
		color.Red(fmt.Sprintf("%s", err))
	}

	params.LaunchTemplatesFromFolder(files, "")

	// Output status of machine
	msg, err = sh.Command("kubectl", "cluster-info").Output()
	color.Green("Cluster Info")
	color.White(string(msg[:]))

	msg, err = sh.Command("kubectl", "get", "minions").Output()
	color.Green("Minions")
	color.White(string(msg[:]))

	msg, err = sh.Command("kubectl", "get", "services").Output()
	color.Green("Services")
	color.White(string(msg[:]))

	msg, err = sh.Command("kubectl", "get", "replicationcontrollers").Output()
	color.Green("Replication Controllers")
	color.White(string(msg[:]))

	msg, err = sh.Command("kubectl", "get", "pods").Output()
	color.Green("Pods")
	color.White(string(msg[:]))

}

func (params *launchParams) LaunchTemplatesFromFolder(files []os.FileInfo, subDir string) {

	// Deploy generated template files to kubernetes endpoint
	for _, file := range files {

		filename := params.distDir + subDir + file.Name()
		// If file is dir, recurse function
		if file.IsDir() {

			buildFiles, _ := ioutil.ReadDir(params.distDir + subDir + file.Name())
			params.LaunchTemplatesFromFolder(buildFiles, "/"+file.Name())
			continue

		}

		if filepath.Ext(file.Name()) != ".yaml" {

			color.Yellow(fmt.Sprintf("File %s is not a template", filename))
			continue

		}

		// Run create command on specified server
		msg, err := sh.Command("kubectl", "--insecure-skip-tls-verify", os.Getenv("LINK_INSECURE_SKIP_TLS_VERIFY"), "--username", os.Getenv("LINK_USERNAME"), "--password", os.Getenv("LINK_PASSWORD"), "--server", os.Getenv("LINK_SERVER"), "create", "-f", filename).Output()

		if err != nil {
			color.Red(fmt.Sprintf("%s", err))
		}

		output := string(msg[:])

		color.White(output)

	}

}
