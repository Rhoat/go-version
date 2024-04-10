# go-version
This is a simple go version manager. It allows you to print the current version of the application dynamically by taking advantage of the -ldflags option in the go build command.

The main idea is to have a centralize place to manage the version of the application. This is useful when you have multiple microservices and you want to have a consistent way to manage the version of the application.