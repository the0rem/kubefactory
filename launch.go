package main

import (
	"github.com/codeskyblue/go-sh"
	"github.com/fatih/color"
	// "github.com/joho/godotenv/autoload"
	"github.com/joho/godotenv"
	"io/ioutil"
)

type launchParams struct {
	distDir          string
	environment      string
	envFile          string
	sourceFilePrefix string
}

func (params *launchParams) Configure(distDir, environment string) {

	params.distDir = distDir
	params.environment = environment
	params.envFile = ".kfenv_" + environment

}

func (params *launchParams) Launch() {

	var msg []byte

	// Get list of files in dist
	files, err := ioutil.ReadDir(params.distDir)

	// Check for error finding files
	if err != nil {

	}

	// Get env variables
	err = godotenv.Load(params.envFile)

	if err != nil {

	}

	// Create a session
	session := sh.NewSession()

	// Deploy generated template files to kubernetes endpoint
	for _, value := range files {

		if value.IsDir() {
			continue
		}

		// TODO: add correct env files
		filename := params.distDir + value.Name()
		msg, err = session.Command("kubectl", "--kubeconfig=params.envFile", "create", "-f", filename).Output()

		if err != nil {

		}

		output := string(msg[:])

		color.White(output)

	}

	// Output status of machine
	msg, err = session.Command("kubectl", "cluster-info").Output()
	msg, err = session.Command("kubectl", "get", "minions").Output()
	msg, err = session.Command("kubectl", "get", "services").Output()
	msg, err = session.Command("kubectl", "get", "replicationcontrollers").Output()
	msg, err = session.Command("kubectl", "get", "pods").Output()

}

/*
# Parse arguments
usage() {
    echo "Usage: $0 [-h] [-v] -e ENV"
    echo "  -h --help  Help. Display this message and quit."
    echo "  -v --version Version. Print version number and quit."
    echo "  -d --dir Specify the distribution directory."
    echo "  -e --env Specify configuration file FILE."
    echo "  -u --update Whether to update a current deployment"
    exit
}

distDir="${PWD}/dist"
environment="dev"
envFile=""
sourceFilePrefix=".kubectl_"

update=0
deployDelay=10

while (( $# > 0 ))
do
    option="$1"
    shift

    case $option in
    -h|--help)
        usage
        exit 0
        ;;
    -v|--version)
        echo "$0 version $version"
        exit 0
        ;;
    -e|--env)  # Example with an operand
        environment="$1"
        shift
        ;;
    -d|--dir)  # Example with an operand
        distDir="$1"
        shift
        ;;
    -u|--update)  # Example with an operand
        update=1
        ;;
    -*)
        echo "Invalid option: '$opt'" >&2
        exit 1
        ;;
    *)
        # end of long options
        break;
        ;;
   esac

done

if [ ! $envFile ]; then
    envFile=${sourceFilePrefix}${environment}
fi

if [ ! -f $envFile ]; then
    echo "Environment file '$envFile' cannot be sourced" >&2
    exit 1
fi

source $envFile

# We only want to update
if [ $update -eq 1 ]; then
    for file in $distDir/*.yaml; do

        isReplicationController=$(grep 'kind: ReplicationController' $file)

        if [ -z $isReplicationController ]; then
            # Runs a graceful update of the deployed applicaiton
            kubectl rolling-update "$deploymentName" --update-period="$deployDelay"s -f "$newDeploymentFile"
        fi

    done
else

    # Build tempaltes using environmental variables
    for file in $distDir/*.yaml; do

        kubectl create -f $file
        echo "Creating $file"

    done
fi

kubectl cluster-info
kubectl get minions
kubectl get services
kubectl get replicationcontrollers
kubectl get pods
*/
