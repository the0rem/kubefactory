/*

Kubefactory is a package for managing the orchestration and lifecycle management of software packages
running on kubernetes. It works closely with kubectl, kubernetes' CLI utility to make building and
managing templates easier, while enabling deployment to different environments a breeze. Extra tools
for generating templates, launching into containers and remote folder synchronisation are also included.

*/
package main

import (
	"os"
	// "strings"
  // "fmt" 
	"gopkg.in/alecthomas/kingpin.v1"
  // "github.com/GoogleCloudPlatform/kubernetes/pkg/kubectl/cmd"

)

var (

	/**
	 * Initialise main app and general options
	 */
	app         = kingpin.New("kubefactory", "Uses the power of the force to make Kubernetes do its biddings.")
	debug       = app.Flag("debug", "Enable debug mode.").Bool()
	appTest     = app.Flag("dryrun", "Take it for a test drive first.").Bool()
	environment = app.Flag("environment", "Environment to control. Default 'dev'").Default("dev").String()
	fleetctl    = app.Flag("fleetctl-endpoint", "Enpoint for fleetctl.").OverrideDefaultFromEnvar("FLEETCTL_ENDPOINT").URL()
	kubernetes  = app.Flag("kubernetes-master", "Enpoint for kubernetes.").OverrideDefaultFromEnvar("KUBERNETES_MASTER").URL()

	/**
	 * Add command to handle generating templates
	 */
	generate = app.Command("generate", "Generate a new deployment template.")

	/**
	 * Add command to handle building deployment files from templates
	 */
	build = app.Command("build", "Configure a new deployment from template files.")

	/**
	 * Add command to launch built configs
	 * @type {[type]}
	 */
	launch = app.Command("launch", "Launch templates in an environment.")

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

	links      = link.Command("list", "Get current links.")
	linkRemove = link.Command("remove", "Remove a link (Data will persist on remote).")

	addlink = link.Command("add", "Add a new link.")
	linkKey = addlink.Flag("SSH Key", "Path to SSH key for remote login.").ExistingFile()
	local   = addlink.Arg("local", "Path to local folder.").Required().ExistingDir()
	remote  = addlink.Arg("remote", "Remote .").Required().String()
)

func main() {

	// switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	// case launch.FullCommand():
	//   println(*environment)

	// }

	switch app.Parse() {

	case "generate":
		Generate()
	case "build":
		Build()
	case "launch":
		Launch()
	case "enter":
		Enter()
	case "supercede":
		Supercede()
	case "kill":
		Kill()
	case "link list":
		Link.List()
	case "link add":
		Link.Add()
	case "link remove":
		Link.Remove()
	}
}
