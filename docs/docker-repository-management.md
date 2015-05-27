# Repository Management

Repositories can either be stored publicly or privately. For most projects the public docker hub repository will do great. Private repositories are also easily available however you may want to host your own repository for security or cost-saving reasons.

This will guide you in setting up a private repository on your Kubernetes cluster while also going throught he considerations that are required.

The [Docker Registry HTTP API V2](docs/spec/api.md) is the latest iteration from Docker for managing your private Docker repository (now in Golang). The registry supports multiple data backends and is easy to setup for SSL.

For the purpose of this exercise I am going to use local storage, however the Kubernetes template is setup to allow for easy plugging in of backends. 

## Exposing the repository

Its worth considering where the repository should be accessible from. For security's sake, you would only want the services comsuming it to be able to access it (namely your Kubernetes & Docker build/deploy tools). For this reason you could keep the service internal to the cluster and provide an endpoint via SkyDNS. This would allow for internal DNS resolution of the repository while keeping it hidden elsewhere.

You might be thinking: That doesn't make sense, I need it across multiple clusters. If thats the case its easy enough to spin up in each cluster with a shared backend filesystem such as a key-value store (Look at the Docker Registry docs for more).

You would also want to have the repository accessible locally............?>??>?>>?/