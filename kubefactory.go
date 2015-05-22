/*

Kubefactory is a package for managing the orchestration and lifecycle management of software packages
running on kubernetes. It works closely with kubectl, kubernetes' CLI utility to make building and
managing templates easier, while enabling deployment to different environments a breeze. Extra tools
for generating templates, launching into containers and remote folder synchronisation are also included.

This app is used to enhance certain features of kubectl, not replace them.

*/
package main

import (
	"gopkg.in/alecthomas/kingpin.v1"
	"os"
	"path/filepath"
)

var (
	err error

	// Get path to current directory
	filePath, _ = filepath.Abs(filepath.Dir(os.Args[0]))

	pwd, _ = os.Getwd()

	/**
	 * Initialise main app and general options
	 */
	app         = kingpin.New("kubefactory", "Uses the power of the force to make Kubernetes do its biddings.")
	debug       = app.Flag("debug", "Enable debug mode.").Bool()
	appTest     = app.Flag("dryrun", "Take it for a test drive first.").Bool()
	environment = app.Flag("environment", "Environment to control. Default 'dev'").Default("dev").String()
	envFile     = app.Flag("envFile", "Environment file for connecting to kubernetes relative to the environment-specific directory. (Default: .kfenv)").Default(".kfenv").String()
	envDir      = app.Flag("envDir", "Root directory for environment directories").Default(pwd + "/environments/").String()
	// fleetctl    = app.Flag("fleetctl-endpoint", "Enpoint for fleetctl.").OverrideDefaultFromEnvar("FLEETCTL_ENDPOINT").URL()
	// kubernetes  = app.Flag("kubernetes-master", "Enpoint for kubernetes.").OverrideDefaultFromEnvar("KUBERNETES_MASTER").URL()

	/**
	 * Add command to generate the template environment
	 */
	initialise = app.Command("init", "Initalise a new deployment folder.")

	/**
	 * Add command to handle building deployment files from templates
	 */
	build          = app.Command("build", "Configure a new deployment from template files.")
	templateSource = build.Arg("templateSource", "Source directory for YAML deployment templates").Default(pwd + "/templates/").String()
	buildDest      = build.Arg("buildDest", "Destination directory for saving generated distribution files").Default(filePath + "/dist/").String()

	/**
	 * Add command to launch built configs
	 * @type {[type]}
	 */
	launch    = app.Command("launch", "Launch templates in an environment.")
	launchSrc = launch.Arg("launchSrc", "Source directory for loading template files").Default(pwd + "/dist/").String()

	/**
	 * Add command to enter a container via bash
	 */
	enter         = app.Command("enter", "Open an interactive terminal within a container. This is useful for gaining  access to containers when developing, or debugging critical issues in production. You can liken it to SSH access.")
	podName       = enter.Arg("podName", "Name of pod to enter. We're smart enough to figure out which one, so give us the start of the name.").Required().String()
	containerName = enter.Arg("containerName", "Name of container to enter. We're smart enough to figure out which one, so give us the start of the name. If it's the same as the pod, you don't need to add the container.").String()

	/**
	 * Add command to scale up new container versions while scaling down old versions
	 */
	upgrade         = app.Command("upgrade", "Deploy by scaling down old and scaling up new. Replaces the specified controller with new controller, updating one pod at a time to use the new PodTemplate. The new-controller.yaml must specify the same namespace as the existing controller and overwrite at least one (common) label in its replicaSelector.")
	currentRC       = upgrade.Arg("old", "Name of resource controller to update.").Required().String()
	newRC           = upgrade.Arg("new", "Filename of new resource controller to upgrade to.").Required().String()
	upgradeInterval = upgrade.Arg("interval", "Time between updates.").Required().String()

	/**
	 * Add command to handle synchronisation between local folders and remote folders
	 */
	link = app.Command("link", "Manage synchronised folders between local and remote environments.")

	linkList = link.Command("list", "Get current links.")

	linkAdd  = link.Command("add", "Add a new link.")
	linkKey  = linkAdd.Flag("SSH Key", "Path to SSH key for remote login.").ExistingFile()
	linkFrom = linkAdd.Arg("from", "Path to local folder.").Required().ExistingDir()
	linkTo   = linkAdd.Arg("to", "Remote destination (Must be of the form \"user@host:path/to/dir\"").Required().String()

	linkUp     = link.Command("up", "Enable a currently saved link.")
	linkDown   = link.Command("down", "Disable a currently saved link.")
	linkRemove = link.Command("remove", "Disable and remove a link.")
)

func main() {

	result := kingpin.MustParse(app.Parse(os.Args[1:]))

	// println(result)
	// println()

	switch result {

	// case generate.FullCommand():

	// 	println((*generate).FullCommand())
	// 	Generate()

	case initialise.FullCommand():

		Init(pwd, *envDir+"/"+*environment, *envDir+"/"+*environment+"/"+*envFile)

	case build.FullCommand():

		builder := new(builder)
		builder.Configure(*buildDest, *templateSource, *envDir+*environment+"/")
		builder.Build()

	case launch.FullCommand():

		deployment := new(launchParams)

		deployment.Configure(*launchSrc, *envFile, *envDir+*environment+"/"+*envFile)
		deployment.Launch()

	case enter.FullCommand():

		// Hamndle if only a pod is given
		if *containerName != "" {
			Enter(*podName, *containerName, *envDir+*environment+"/"+*envFile)
		} else {
			Enter(*podName, *podName, *envDir+*environment+"/"+*envFile)
		}

	case upgrade.FullCommand():

		Upgrade(*currentRC, *newRC, *upgradeInterval)

	case (*linkList).FullCommand():

		dirSync := new(dirSync)
		dirSync.List()

	case (*linkAdd).FullCommand():

		dirSync := new(dirSync)

		dirSync.Configure(*linkFrom, *linkTo)
		dirSync.Add()

	case (*linkRemove).FullCommand():

		dirSync := new(dirSync)
		dirSync.Remove()

	case (*linkUp).FullCommand():

		dirSync := new(dirSync)
		dirSync.Up()

	case (*linkDown).FullCommand():

		dirSync := new(dirSync)
		dirSync.Down()

	}
}
