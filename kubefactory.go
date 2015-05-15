/*

Kubefactory is a package for managing the orchestration and lifecycle management of software packages
running on kubernetes. It works closely with kubectl, kubernetes' CLI utility to make building and
managing templates easier, while enabling deployment to different environments a breeze. Extra tools
for generating templates, launching into containers and remote folder synchronisation are also included.

*/
package main

import (
	"gopkg.in/alecthomas/kingpin.v1"
	"os"
	"path/filepath"
)

var (

	// Get path to current directory
	filePath, err = filepath.Abs(filepath.Dir(os.Args[0]))

	/**
	 * Initialise main app and general options
	 */
	app         = kingpin.New("kubefactory", "Uses the power of the force to make Kubernetes do its biddings.")
	debug       = app.Flag("debug", "Enable debug mode.").Bool()
	appTest     = app.Flag("dryrun", "Take it for a test drive first.").Bool()
	environment = app.Flag("environment", "Environment to control. Default 'dev'").Default("dev").String()
	envFile     = app.Flag("envFile", "Environment file for connecting to kubernetes. (Default: "+filePath+"/.kfenv_dev)").Default(filePath + "/.kfenv_dev").String()
	fleetctl    = app.Flag("fleetctl-endpoint", "Enpoint for fleetctl.").OverrideDefaultFromEnvar("FLEETCTL_ENDPOINT").URL()
	kubernetes  = app.Flag("kubernetes-master", "Enpoint for kubernetes.").OverrideDefaultFromEnvar("KUBERNETES_MASTER").URL()

	/**
	 * Add command to handle generating templates
	 */
	generate = app.Command("generate", "Generate a new deployment template.")

	/**
	 * Add command to handle building deployment files from templates
	 */
	build          = app.Command("build", "Configure a new deployment from template files.")
	envDir         = build.Arg("envDir", "Root directory for environment directories").Default(filePath + "/environments/").String()
	templateSource = build.Arg("templateSource", "Source directory for YAML deployment templates").Default(filePath + "/templates/").String()
	buildDest      = build.Arg("buildDest", "Destination directory for saving generated distribution files").Default(filePath + "/dist/").String()

	/**
	 * Add command to launch built configs
	 * @type {[type]}
	 */
	launch    = app.Command("launch", "Launch templates in an environment.")
	launchSrc = launch.Arg("launchSrc", "Source directory for loading template files").Default(filePath + "/dist/").String()

	/**
	 * Add command to enter a container via bash
	 */
	enter     = app.Command("enter", "Deploy templates to environment.")
	container = enter.Arg("container", "Name of container to enter.").Required().String()

	/**
	 * Add command to scale up new container versions while scaling down old versions
	 */
	supercede = app.Command("supercede", "Deploy by scaling down old and scaling up new.")
	old       = supercede.Arg("old", "Name of resource to kill.").Required().String()
	rise      = supercede.Arg("new", "Name of resource to rise up.").Required().String()
	speed     = supercede.Arg("interval", "Time between kills.").Required().Int()

	/**
	 * Add command to kill specific resources
	 */
	kill    = app.Command("kill", "Take down pods resource controllers and services.")
	targets = kill.Arg("targets", "Name of resources to kill.").Required().Strings()

	/**
	 * Add command to handle synchronisation between local folders and remote folders
	 */
	link = app.Command("link", "Manage synchronised folders between local and remote environments.")

	linkList = link.Command("list", "Get current links.")

	linkAdd = link.Command("add", "Add a new link.")
	linkKey = linkAdd.Flag("SSH Key", "Path to SSH key for remote login.").ExistingFile()
	local   = linkAdd.Arg("local", "Path to local folder.").Required().ExistingDir()
	remote  = linkAdd.Arg("remote", "Remote .").Required().String()

	linkRemove = link.Command("remove", "Remove a link (Data will persist on remote).")
)

func main() {

	result := kingpin.MustParse(app.Parse(os.Args[1:]))

	// println(result)
	// println()

	switch result {

	case generate.FullCommand():

		println((*generate).FullCommand())
		// Generate()

	case build.FullCommand():

		builder := new(builder)
		builder.Configure(*buildDest, *templateSource, *envDir+*environment+"/partials/")
		builder.Build()

	case launch.FullCommand():

		deployment := new(launchParams)

		deployment.Configure(*launchSrc, *envFile+*environment)
		deployment.Launch()

	case enter.FullCommand():

		println((*enter).FullCommand())
		println(*container)
		// Enter()

	case supercede.FullCommand():

		println((*supercede).FullCommand())
		// Supercede()

	case kill.FullCommand():

		println((*kill).FullCommand())
		// Kill([]string{"rc1","service1"})

	case (*linkList).FullCommand():

		println((*linkList).FullCommand())
		//   dirSync := new (dirSync)

		//   dirSync.configure()
		// dirSync.List()

	case (*linkAdd).FullCommand():

		println((*linkAdd).FullCommand())
		//   dirSync := new (dirSync)

		//   dirSync.configure()
		// dirSync.Add()

	case (*linkRemove).FullCommand():

		println((*linkRemove).FullCommand())
		//   dirSync := new (dirSync)

		//   dirSync.configure()
		// dirSync.Remove()

	}
}
