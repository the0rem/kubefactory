# kubefactory

usage: kubefactory [<flags>] <command> [<flags>] [<args> ...]

Uses the power of the force to make Kubernetes do its biddings.

Flags:
  --help               Show help.
  --debug              Enable debug mode.
  --dryrun             Take it for a test drive first.
  --environment="dev"  Environment to control. Default 'dev'
  --fleetctl-endpoint=http://172.17.8.101:4001  
                       Enpoint for fleetctl.
  --kubernetes-master=http://172.17.8.101:8080  
                       Enpoint for kubernetes.

Commands:
  help [<command>]
    Show help for a command.

  generate
    Generate a new deployment template.

  build
    Configure a new deployment from template files.

  launch
    Launch templates in an environment.

  enter <container>
    Deploy templates to environment.

  supercede <old> <new> <interval>
    Deploy by scaling down old and scaling up new.

  kill <targets>
    Take down pods resource controllers and services.

  link list
    Get current links.

  link remove
    Remove a link (Data will persist on remote.

  link add [<flags>] <local> <remote>
    Add a new link.
