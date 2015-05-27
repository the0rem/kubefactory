# Writing templates

Google kubernetes provides a comprehensive templating system for provisioning your resources on the cloud. The following docs will help you get started writing template files for kubernetes

Example templates - https://github.com/GoogleCloudPlatform/kubernetes/tree/master/examples
Template specifications - https://godoc.org/github.com/GoogleCloudPlatform/kubernetes/pkg/api/v1beta3

The template specifications are the best place for understanding how templates can be built however there is an initial learning curve if you're not familiar with the documentation style.

## Building template tailored to your environment

To have your templates coordinate with environmental differences you will need to have the following:

 - 1. A reference in the template to inject another template file
 - 2. An environment file to inject into the templates

This will allow you to state when a template (an environmental template) should be injected into another template (your base template) when required. From here you can define which environment you would like to build for and it will handle the "mixing" of your template files.


You're going to need to build for different environments right? Of course! This tool wil help inject the dependencies of each build environment iwth the data that is required. This could be the dev files for the docker containers, diffferent folder mounts, whatever you wants, just append any section of your templates with keyName: #keyName# and kubefactory will track down the yaml file to inject for each environment. This must be in your /environments folder for the environment that you have selected

#### Configuring environmental dependencies into templates

One thing that the kubernetes system doesn't cater for is the changes which need to be made depending on the deployment enironment. For example, you may wish to pass different environment variables, set different IP addresses, replication sizes, or pass a folder mount for development only. 

This can be acheived by identifying where environment dependencies should be injected. Once you can do this, you can then add the environment variables for your environment


```
 - parameter: #environmentFilename#
```

![Injecting Environmental Dependencies](https://raw.githubusercontent.com/the0rem/kubefactory/master/docs/images/environment-inject.png)