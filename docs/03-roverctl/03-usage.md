# Usage

To use, simply run `roverctl` in your shell of choice after installation. You will be greeted with a screen of available options. The first time you boot, you will be asked to pick an author name, which will be added to all code you write and upload to the Rover. It is recommended to first set up a connection with a `roverd` instance, but you can already get started writing a service when your Rover is offline.

Much `roverctl` functionality is based on the public REST API exposed by the `roverd` process, which runs on the Rover. To *connect to a Roverd* you are actually connecting to a *`roverd` instance* instance running on that Rover. Before connecting, make sure that your Rover is powered on and both you and the Rover are connected to the ASE labs WiFi network.

## Connecting

1. Open `roverctl`
2. Select `connect`

Upon connecting you need to specify the index of your Rover, which should be enough to get you up and running. For advanced options, press <kbd>ctrl</kbd>+<kbd>a</kbd>. The username and password combination are not necessarily identical to your SSH credentials.

## Uploading a service

1. Open `roverctl`
2. Select `Upload to Rover`

By default, `roverctl` will upload the service it finds in your current working directory. If no service directory is found, you can select one manually with the commands shown on the upload screen.