# kubefactory

Kubefactory is a package for managing the orchestration and lifecycle management of software packages
running on kubernetes. It works closely with kubectl, kubernetes' CLI utility to make building and
managing templates easier, while enabling deployment to different environments a breeze. Extra tools
for generating templates, launching into containers and remote folder synchronisation are also included.

This package aims to solve the logistics around deploying your applications to different environments. By integrating environmental dependencies into packages, you can use the same template bases to deploy for all cycles through development to production.


## TODO

 - [ ] Add example templates to expose full features of Kubernetes
 - [ ] Build SkyDNS deployment template
 - [ ] Add logstash/kabana deployment template for logging
 - [ ] Add API access for commands
 - [ ] Add logging for commands
 - [ ] Add post-event hooks for commands
 - [ ] Integrate rafecolton/docker-builder into a docker build and build API hook command (currently has features for queueing and serving the API which we should bring upstream)
 - [ ] Add feature to package a Dockerfile with dependencies to run kubefactory for a specific environment
 - [ ] Add releases for architecture environments

# Goal

 - Leverage kubernetes with other tools to create a framework for building containerised applications from development through to production


## Usage

```
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
```

## Commands
  
### Help

Show help for a command.

```
  kubefactory help [<command>]
```

### Init

Initialise the default project folder structure including an example config file.


### Build 

Builds an environment-specific release of your deployment using your deployment templates along with your environmental files.

```
  kubefactory build [<envDir> [<templateSource> [<buildDest>]]]

  Args:
    [<envDir>]          Root directory for environment directories
    [<templateSource>]  Source directory for YAML deployment templates
    [<buildDest>]       Destination directory for saving generated distribution files
```


### Launch 

Launch templates in an environment.

```
  kubefactory launch [<launchSrc>]

  Args:
    [<launchSrc>]  Source directory for loading template files
```

### Enter 

Open an interactive terminal within a container. This is useful for gaining access to containers when developing, or debugging critical issues in production. You can liken it to SSH access.

Don't know the container name? Don't worry. This interactive command will help get you there.

```
  kubefactory enter <podName> [<containerName>]

  Args:
    <podName>          Name of pod to enter. We're smart enough to figure out which one, so give us the start of the name.
    [<containerName>]  Name of container to enter. We're smart enough to figure out which one, so give us the start of the name. If it's the same as the pod, you don't need to add the container.
```

### Upgrade

This tool will handle pushing out deployment updates of pods under replication controllers. It is an overlay of kubectl rolling-update which uses timestamps as metadata in the build command to handle identifying new template deployments.

This is a simple process to identify a change in the state of the containers.

The goal is to eventually use the docker commit hash

```
  kubefactory upgrade <old> <new> <interval>
```

## Links 

Manage synchronised folders between local and remote environments.

### Mac OS X

Prerequisites: 
 - https://github.com/ggreer/fsevents-tools

### Linux

Prerequisites: 
 - https://github.com/drunomics/syncd


### Get Links

Get current links.

```
  kubefactory link list
```

### Add Link

Add a new link.

```
  kubefactory link add [<flags>] <from> <to>

  Flags:
    --SSH Key=SSH KEY  Path to SSH key for remote login.

  Args:
    <from>  Path to local folder.
    <to>    Remote destination (Must be of the form "user@host:path/to/dir"
```

### Resume a link

Enable a currently saved link.

```
  kubefactory link up
```

### Disable a link

Disable a currently saved link.

```
  kubefactory link down
```

### link remove

Disable and remove a link.


## Installing a CoreOS Kubernetes cluster
https://github.com/GoogleCloudPlatform/kubernetes/tree/master/docs/getting-started-guides

### OpenStack
https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/getting-started-guides/juju.md

### AWS
https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/getting-started-guides/aws-coreos.md

### Azure
https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/getting-started-guides/azure.md

### Juju
https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/getting-started-guides/juju.md

### Vagrant
https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/getting-started-guides/vagrant.md
https://github.com/pires/kubernetes-vagrant-coreos-cluster












