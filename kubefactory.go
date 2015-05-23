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
)

/**
 *
 */
type appConfig struct {
	debug       bool
	appTest     bool
	environment string
	envFile     string
	envDir      string
	workingDir  string
}

var (
	/**
	 * Error variable
	 */
	err error

	/**
	 * Get path to current directory
	 * @type {string}
	 */
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

	/**
	 * Add command to generate the template environment
	 */
	initialise = app.Command("init", "Initalise a new deployment folder.")

	/**
	 * Add command to handle building deployment files from templates
	 */
	build          = app.Command("build", "Configure a new deployment from template files.")
	templateSource = build.Arg("templateSource", "Source directory for YAML deployment templates").Default(pwd + "/templates/").String()
	buildDest      = build.Arg("buildDest", "Destination directory for saving generated distribution files").Default(pwd + "/dist/").String()

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

/**
 * [main description]
 * @return {[type]} [description]
 */
func main() {

	result := kingpin.MustParse(app.Parse(os.Args[1:]))

	// Initalise app config data
	appConfig := appConfig{
		debug:       *debug,
		appTest:     *appTest,
		envDir:      *envDir + *environment,
		envFile:     *envDir + *environment + "/" + *envFile,
		environment: *environment,
		workingDir:  pwd,
	}

	switch result {

	/**
	 * Initialise project folder structure
	 */
	case initialise.FullCommand():

		appConfig.Init()

	/**
	 * Build templates for environment
	 */
	case build.FullCommand():

		builder := new(builder)
		builder.Configure(*buildDest, *templateSource, appConfig.envDir)
		builder.Build()

	/**
	 * Launch built templates to environment
	 */
	case launch.FullCommand():

		appConfig.Launch(*launchSrc)

	/**
	 * Enter a container
	 */
	case enter.FullCommand():

		// Hamndle if only a pod is given
		if *containerName != "" {
			appConfig.Enter(*podName, *containerName)
		} else {
			appConfig.Enter(*podName, *podName)
		}

	/**
	 * Upgrade replication controller pods
	 */
	case upgrade.FullCommand():

		appConfig.Upgrade(*currentRC, *newRC, *upgradeInterval)

	/**
	 * List directory links
	 */
	case (*linkList).FullCommand():

		dirSync := new(dirSync)
		dirSync.List()

	/**
	 * Add directory link
	 */
	case (*linkAdd).FullCommand():

		dirSync := new(dirSync)
		dirSync.Configure(*linkFrom, *linkTo)
		dirSync.Add()

	/**
	 * Remove directory link
	 */
	case (*linkRemove).FullCommand():

		dirSync := new(dirSync)
		dirSync.Remove()

	/**
	 * Enable directory link
	 */
	case (*linkUp).FullCommand():

		dirSync := new(dirSync)
		dirSync.Up()

	/**
	 * Disable directory link
	 */
	case (*linkDown).FullCommand():

		dirSync := new(dirSync)
		dirSync.Down()

	}
}
