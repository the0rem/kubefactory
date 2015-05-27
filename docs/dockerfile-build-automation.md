# Dockerfile build automation

Running on the assumption that Docker builds in modern day development revolve around a git repository, we can setup a build tool which can perform the following:

 - Respond to a web hook from a git repo when commits occur
 - Tag docker builds based off git commit tags/branch name
 - Push docker build to a repository

These features allow for the automation of the docker image lifecycle ready for use on Kubernetes. These features already exist with a tool rafecolton/docker-builder. Great now what!

We still need to bridge the gap in automation between a docker push and a kubefactory build/launch. Luckily Docker Repository allows you to configure notification endpoints in the config. The following is an example for configuring a callback to trigger a Kubernetes deployment.

There are some minor drawback to this, the callback has to be resolvable from the registry. That shouln't be an issue if you are injecting your local development into your cluster for dev work. Your production cluster will work just fine alongside this.

It does bring up another consideration however... you will need to be able to access kubefactory via API and also have it resolve relative to the Registry. 

This presents a fork in the road for using the kubefactory tool. Up until this point the tool has been used manually, auto automated on a CLI level only. We will need to look at giving kubefactory a Dockerfile for repeatable builds as well as advertising the image publically.

We will also need to attach API accessibility to kubefactory with its own post-action hooks.

We will also need to add logging capabilities so that we can monitor and debug an issues as they arise.

Because we're automating this process, we will also need to make sure the .kfenv config file only advertises the environment that its deployed in. What may be a good case for this is a command which can actually export the specific environment removing all others and setting the environment variable. This way each time a build and launch runs, we will be working on the correct details.