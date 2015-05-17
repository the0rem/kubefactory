# kubefactory

Kubefactory is a package for managing the orchestration and lifecycle management of software packages
running on kubernetes. It works closely with kubectl, kubernetes' CLI utility to make building and
managing templates easier, while enabling deployment to different environments a breeze. Extra tools
for generating templates, launching into containers and remote folder synchronisation are also included.

## TODO

 - [ ] Launch built templates
 - [ ] Enter containers
 - [ ] Finalise environment variables for kubectl integration
 - [ ] Write Docs
 - [ ] Kill resources handler


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

### Generate TODO

Generate a new deployment template.


### Build 

Builds a distribution release of your deployment using your deployment templates along with your environmental files.

#### Writing templates

Google kubernetes provides a comprehensive templating system for provisioning your resources on the cloud. The following docs will help you get started writing template files for kubernetes

Example templates - https://github.com/GoogleCloudPlatform/kubernetes/tree/master/examples
Template specifications - https://godoc.org/github.com/GoogleCloudPlatform/kubernetes/pkg/api/v1beta3

#### Building template tailored to your environment

You're going to need to build for different environments right? Of course! This tool wil help inject the dependencies of each build environment iwth the data that is required. This could be the dev files for the docker containers, diffferent folder mounts, whatever you wants, just append any section of your templates with keyName: #keyName# and kubefactory will track down the yaml file to inject for each environment. This must be in your /environments folder for the environment that you have selected

#### Configuring environmental dependencies into templates

One thing that this system doesn't cater for is the changes which need to be made depending on the deployment enironment. For example, you may wish to pass different environment variables, set different IP addresses, replication sizes, or pass a folder mount for development only. 

This can be acheived by identifying where environment dependencies should be injected.



```
 - string: #string#
```



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

```
  kubefactory enter <podName> [<containerName>]

  Args:
    <podName>          Name of pod to enter. We're smart enough to figure out which one, so give us the start of the name.
    [<containerName>]  Name of container to enter. We're smart enough to figure out which one, so give us the start of the name. If it's the same as the pod, you don't need to add the container.
```

### Supercede TODO

Deploy by scaling down old and scaling up new.

```
  kubefactory supercede <old> <new> <interval>
```

### Kill TODO

Take down pods resource controllers and services.

```
  kubefactory kill <targets>
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


## Generating templates

If you don't know how to reate your own template files, this is the tool to use. An interactive template generator which will help you create your build environment without any knowlegde required in Go










