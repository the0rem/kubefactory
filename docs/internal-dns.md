# Internal DNS management

SkyDNS has been the recommended tool by Kubernetes for DNS resolution within the cluster. This help simplifies the routing of dependencies so that instead of relying on environment variables passed into a container, we can use reliable DNS resolution.

Kubernetes takes care in injecting the DNS resolver information into the docker containers on creation doing a lot of the heavy lifting for us.

