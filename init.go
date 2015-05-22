package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

var configTemplate = []byte(`
apiVersion: v1
kind: Config
preferences: {}

# Set the default context to use
current-context: dev

# Provide a link between a cluster and user
contexts:
  - name: dev
    context:
      # Name of cluster profile
      cluster: dev
      # Name of user profile
      user: dev

# Define the available Kubernetes clusters for the environment
clusters:
  # Lists the information for your Kubernetes clusers (we will only setup one per environment)
  # This information dictates:
  # - where the API server endpoint is (server)
  # - what version the API is compatible with (api-version)
  # - whether to require a valid certificate (insecure-skip-tls-verify)
  # - what certificate autority to expect: 
  # -- (certificate-authority) for a file path to the certificate
  # -- (certificate-authority-data) for a base64 encoding of the certificate data
  - name: dev
    cluster:

      # Server is the address of the kubernetes cluster (https://hostname:port).
      server: https://uri-to-kubernetes-api:listening-port/

      # APIVersion is the preferred api version for communicating with the kubernetes cluster (v1beta1, v1beta2, v1beta3, etc).
      api-version: v1beta3

      # InsecureSkipTLSVerify skips the validity check for the server's certificate. 
      # This will make your HTTPS connections insecure.
      insecure-skip-tls-verify: false

      # Use one of the options below for CA validation
      # CertificateAuthority is the path to a cert file for the certificate authority.
      certificate-authority:
      # CertificateAuthorityData contains PEM-encoded certificate authority certificates. 
      # Overrides CertificateAuthority.
      certificate-authority-data: 


# Define the available users for the environment
users:
  # 
  - name: dev
    user:

      # The following are the available authentication options using the Kubernetes API.
      # Configure the correct option which is configured on your Kubernetes API.

      # Option 1
      # Token-based authentication
      # Token is the bearer token for authentication to the kubernetes cluster.
      token:


      # Option 2
      # Password-based authentication
      # Username is the username for basic authentication to the kubernetes cluster.
      username: 
      # Password is the password for basic authentication to the kubernetes cluster.
      password: 


      # Option 3
      # Key-based authentication
      # ClientCertificate is the path to a client cert file for TLS.
      client-certificate:
      # ClientCertificateData contains PEM-encoded data from a client cert file for TLS. Overrides ClientCertificate.
      client-certificate-data: 
      # ClientKey is the path to a client key file for TLS.
      client-key:
      # ClientKeyData contains PEM-encoded data from a client key file for TLS. Overrides ClientKey.
      client-key-data: 


      # Option 4
      # AuthPath is the path to a kubernetes auth file (~/.kubernetes_auth).  
      # If you provide an AuthPath, the other options specified are ignored
      #auth-path:
`)

func Init(pwd, envDir, envFile string) {

	dirs := []string{pwd + "/dist", pwd + "/templates", envDir + "/partials"}

	for _, directory := range dirs {

		// Initialise Directories
		err = os.MkdirAll(directory, 0755)

		if err != nil {
			color.Red(fmt.Sprintf("Could not create folder: %s", directory))
			os.Exit(1)
		}

		color.Green(fmt.Sprintf("Created folder: %s", directory))

	}

	// Initialise environment file
	err := ioutil.WriteFile(envFile, configTemplate, 0644)

	if err != nil {
		color.Red(fmt.Sprintf("Could not create environment file: %s", envFile))
		os.Exit(1)
	}

	color.Green(fmt.Sprintf("Created environment file: %s", envFile))

}
