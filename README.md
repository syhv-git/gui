# About
This a template for creating customized desktop applications. The primary API driver for this template is the Fyne go module. The documentation can be found [here](https://pkg.go.dev/fyne.io/fyne/v2) and the source code is found [here](https://github.com/fyne-io/fyne).

# Prerequisites
This package is built on Go v1.18. Run the following to ensure all dependent packages exist on your system.

For Debian:`sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev`

Next, open a terminal in the workspace root, and run:

```
go mod init
go get -u -t ./...
go mod tidy
```
> If you are using GoLand, make sure go modules integration is enabled in the Go settings

# Getting Started
To start the application, run the command: `go run main.go`. Add build flags depending on system specs and intended behavior.
