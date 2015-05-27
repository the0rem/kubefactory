# Properly handling SSL
One of the issues with this method is the reliance of SSL keys residing on the load balancer. There are two paths to this:
 - reverse-proxy the requests and let the pods handle SSL termination
 - store SSL keys on load balancer and find a way to store keys

The first option, reverse-proxy is the simplest solution, however there are issues with passing client header information when this is the case.

The second option, storing keys on the load balancer, would require a few things to keep the keys secure. After all, we want to be very careful with who can access the private key for an SSL connection.

Just a quick note: This isn't an issue if you only need to serve one SSL certificate, or if that state of your SSL certificates does not change. However you will still need to find a way to manage the SSL certificates when SSL changes occur such as renewal. The system is designed to listed to etcd for updates, not the certs folder.

##
Etcd

Etcd is already well in the mix with the clustered distribution os key:value information. The are where it lacks is security of information. Anyone who can access this service can access the information. One solution could be to add another level of "encapsulation" in the form of encryption.

This would allow us to encrypt the SSL key information, and have the load balancer decrypt it and add it for configuration. 

This seems to be the best option so far, however we still need to figure out the following:
 - How do we manage key-pair creation including the case of re-keying?
 - How will a pod know the public key of the load balancer?
 - What will happen if the private key changes to already configured hosts?

Let's take a look:

### How do we manage key-pair creation including the case of re-keying?
At this stage we're putting a lot of trust in the load balancer. Not only is it responsible for routing to services, it will also hold the keys to data security across the cluster.

For this reason it looks like we should entrust the load balancer (or related service in its pod) with handling key creation. This service could trigger a re-key via API (More on legacy keys later).

### How will a pod know the public key of the load balancer?
With a service built for handling key-pair generation, we could extend the API to provide a call which returns the public key available for encryption.

Once we have a public key to encrypt our data with, we will have to find a way to inject it into our Kubernetes deployment template.

At the moment kubefactory can inject other partial templates for an environment into the base templates. We could follow this path and provide a tool which can query the encryption API and encrypt the cert file with the public key. We could then place this encrypted key-data into a partial environment template as part of the metadata for the controller.

This metadata would then be updated in etcd when the template is updated on the cluster. At this point, the watch on Kubertenes by the load balancer could look for the encrypted SSL key, retreive it for the host, unencode it, and save it in the cert folder under the hostname required.

It would be ideal if the kubefactory could take other inputs for environmental partials other than just .yaml files. You could for example reference a file with a she-bang which would execute code that would return the required valid YAML data for injection. This would allow us to inject the encrypted key data into the template at build time in an automated fahsion.

### What will happen if the private key changes to already configured hosts?

If we follow these other two paths, the only thing we would need to do is ensure that old private keys are kept when a new key-pair is generated. This would ensure that decrypting with gpg would be able to find the correct private key, however all new encryptions will have the latest public key returned via the API for new encryption instances.


### Required tools
To accomplish the encryption handling, a tool exists called Crypt (http://xordataexchange.github.io/crypt/). More research still needs to be made as to whether this is the best solution. If we use the solution outlined above, we can use straight GnuPGP to hande the all of the required components.