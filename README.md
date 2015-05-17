# kubefactory

usage: kubefactory [<flags>] <command> [<flags>] [<args> ...]

Uses the power of the force to make Kubernetes do its biddings.

Flags:
  --help               Show help.
  --debug              Enable debug mode.
  --dryrun             Take it for a test drive first.
  --environment="dev"  Environment to control. Default 'dev'
  --envFile="/Users/shaunsmekel/go/src/github.com/the0rem/kubefactory/.kfenv_dev"  
                       Environment file for connecting to kubernetes. (Default: /Users/shaunsmekel/go/src/github.com/the0rem/kubefactory/.kfenv_dev)
  --fleetctl-endpoint=http://172.17.8.101:4001  
                       Enpoint for fleetctl.
  --kubernetes-master=http://172.17.8.101:8080  
                       Enpoint for kubernetes.

Commands:
  help [<command>]
    Show help for a command.

  generate
    Generate a new deployment template.

  build [<envDir> [<templateSource> [<buildDest>]]]
    Configure a new deployment from template files.

  launch [<launchSrc>]
    Launch templates in an environment.

  enter <podName> [<containerName>]
    Open an interactive terminal within a container. This is useful for gaining access to containers when developing, or debugging critical issues in
    production. You can liken it to SSH access.

  supercede <old> <new> <interval>
    Deploy by scaling down old and scaling up new.

  kill <targets>
    Take down pods resource controllers and services.

  link list
    Get current links.

  link add [<flags>] <from> <to>
    Add a new link.

  link up
    Enable a currently saved link.

  link down
    Disable a currently saved link.

  link remove
    Disable and remove a link.


## Installing a cluster

  ### OpenStack

  ### AWS

  ### Azure

  ### Rackspace

  ### Juju

  ### Vagrant

  ### Libvirt


## Generating template 

If you don't know how to reate your own template files, this is the tool to use. An interactive template generator which will help you create your build environment without any knowlegde required in Go

## Building template to your environment

You're going to need to build for different environments right? Of course! This tool wil help inject the dependencies of each build environment iwth the data that is required. This could be the dev files for the docker containers, diffferent folder mounts, whatever you wants, just append any section of your templates with keyName: #keyName# and kubefactory will track down the yaml file to inject for each environment. This must be in your /environments folder for the environment that you have selected









