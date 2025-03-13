# Usage

`roverctl` is a CLI with various (nested) options. To view all available commands, run:
```bash
roverctl help
```

To generate autocompletion for your shell, run:
```bash
roverctl completion
```

## Specify a Rover

Many commands require a Rover to be specified either through a *Rover id* (the index between 1-20 that you find on the Rover) or a *host* (to override the default IP address used to connect). You can also specify a Roverd *username* and *password* if you want to override the default username and password used. Examples include:

```bash
# Open roverctl-web for Rover 12
roverctl --rover 12

# Open roverctl-web for a Rover at 192.168.0.112
roverctl --host 192.168.0.112

# Open roverctl-web for Rover 12 with a custom username and password
roverctl --username admin --password welcome123
``` 


## Open `roverctl-web`

To open the `roverctl-web` interface, you can just run `roverctl` followed by the Rover you want to connect to. **Docker needs to be installed**.

### Example Usage
```bash
# Open roverctl-web for Rover 12
roverctl --rover 12

# Open roverctl-web for a Rover at 192.168.0.112
roverctl --host 192.168.0.112

# Open roverctl-web for Rover 12 with a custom username and password
roverctl --username admin --password welcome123
```

## View Build Info

To view information about your `roverctl` and/or `roverd` installation, you can use the `roverctl info` command. When no Rover is specified, `roverd`-specific information will not be fetched.

### Example Usage
```bash
# View info about roverctl
roverctl info

# View info about Rover 12
roverctl info -r 12

# View info about a Rover at 192.168.0.112
roverctl info --host 192.168.0.112
```


## Initialize a Service

To create a new service in your current working directory (a new subfolder will be created), you can use the `roverctl service init` command followed by the language you want to create your service in. You will need to specify a valid `--name` and `--source` for your service.

### Example Usage
```bash
# See which languages are available
roverctl service init

# Initialize a go service called "my-distance-sensor" at "github.com/ielaajezdev/my-distance-sensor"
roverctl service init go --name my-distance-sensor --source github.com/ielaajezdev/my-distance-sensor
``` 

## Upload a Service

Service directories can be uploaded to the Rover using the `roverctl upload` command. You can specify a Rover as shown [above](#specify-a-rover). You can upload one or multiple directories simultaneously and enable a file watcher to continuously sync your local changes to the Rover.

### Example Usage
```bash
# Upload your current working directory (.) to Rover 12
roverctl upload . -r 12

# Upload your current working directory (.) to Rover 12 and keep watching for changes
roverctl upload . -r 12 --watch

# Upload multiple directories to Rover 12
roverctl upload /dir1 /dir2 -r 12

# Upload multiple direcotires to a Rover at 192.168.0.112
roverctl upload /dir1 /dir2 --host 192.168.0.112
```

## View Pipeline

To view the current enabled pipeline on the Rover, you can use the `roverctl pipeline` command. You can specify a Rover as shown [above](#specify-a-rover). 

### Example Usage
```bash
# View the pipeline of Rover 12
roverctl pipeline -r 12

# View the pipeline of a Rover at 192.168.0.112
roverctl pipeline --host 192.168.0.112
```

## Start or Stop Pipeline

To start or stop the current enabled pipeline on the Rover, you can use the `roverctl pipeline [start/stop]` command. You can specify a Rover as shown [above](#specify-a-rover). 

### Example Usage
```bash
# Start the pipeline of Rover 12
roverctl pipeline start -r 12

# Start the pipeline of a Rover at 192.168.0.112
roverctl pipeline start --host 192.168.0.112

# Stop the pipeline of Rover 12
roverctl pipeline stop -r 12

# Stop the pipeline of a Rover at 192.168.0.112
roverctl pipeline stop --host 192.168.0.112
```

## View Services

To view the currently installed servies on the Rover, you can use the `roverctl services` command. You can specify a Rover as shown [above](#specify-a-rover). 

### Example Usage
```bash
# View the services installed on Rover 12
roverctl services -r 12

# View the servies installed on a Rover at 192.168.0.112
roverctl services --host 192.168.0.112
```

## Enable or Disable Services
To enable or disable services in the pipeline on the Rover, you can use the `roverctl pipeline [enable/disable]` command. You can specify a Rover as shown [above](#specify-a-rover). Services to be enabled must be **fully qualified** (i.e. by author, name and version). When disabling, only the author and name need to be specified.

### Example Usage
```bash
# Enable the vu-ase imaging service (version 1.2.3) in the pipeline of Rover 12
roverctl pipeline enable vu-ase imaging 1.2.3 -r 12

# Enable the vu-ase imaging service (version 1.2.3) in the pipeline of a Rover at 192.168.0.112
roverctl pipeline enable vu-ase imaging 1.2.3 --host 192.168.0.112

# Disable the vu-ase imaging service in the pipeline of Rover 12
roverctl pipeline disable vu-ase imaging -r 12

# Disable the vu-ase imaging service in the pipeline of a Rover at 192.168.0.112
roverctl pipeline disable vu-ase imaging --host 192.168.0.112
```

## View Service Logs
To view logs of current and previous runs of a specific service, you can use the `roverctl logs` command. You can specify a Rover as shown [above](#specify-a-rover). Services must be **fully qualified** (i.e. by author, name and version).

### Example Usage
```bash
# View logs for the the vu-ase imaging service (version 1.2.3) of Rover 12
roverctl logs vu-ase imaging 1.2.3 -r 12

# View logs for the vu-ase imaging service (version 1.2.3) of a Rover at 192.168.0.112
roverctl logs vu-ase imaging 1.2.3 --host 192.168.0.112
```
