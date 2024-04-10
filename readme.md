# go-version
This is a simple go version manager. It allows you to print the current version of the application dynamically by taking advantage of the -ldflags option in the go build command.

The main idea is to have a centralize place to manage the version of the application. This is useful when you have multiple microservices and you want to have a consistent way to manage the version of the application.

Here is what the version printed will look like:
```bash
testing v1.0.0 RELEASE (43d01a72cbf53869a18029b86274ce971eaf6c95) linux/amd64 - BuildDate: 20240410015203
# fmt.Sprintf("%s v%s %s (%s) %s - BuildDate: %s", vi.ApplicationName, vi.BuildVersion, vi.BuildTag, vi.CommitHash, GetTarget(), vi.BuildDate)
```

## Background
ldflags stands for linker flags, and is used to pass in flags to the underlying linker in the Go toolchain.
While there are many different link flags, We are going to focus on the -X flag, which allows you to Set the value of the string variable in importpath named name to value. This is only effective if the variable is declared in the source code either uninitialized or initialized to a constant string expression. -X will not work if the initializer makes a function call or refers to other variables. More information can be found in the go tool documentation, [here](https://golang.org/cmd/link/).

syntax: 
```bash
go build -ldflags="-X 'package_path.variable_name=new_value'"
```

Because we can take advantage of the package path, we can import this package into any of our projects and set the version of the application dynamically. For more information check the below example for usage. 

# Installation 
```bash
go get github.com/pragmaticengineering/go-version
```

# Usage
Setup your application so that you can print the version, This can be done anywhere but for simplicity, I am going to do it in the main.go file.
```go
package main

import (
	"fmt"

	version "github.com/pragmaticengineering/go-version"
)

func main() {
	// Set version
	fmt.Printf("Version: %s\n", version.Version)
}
```
Make sure to use a go mod init and go mod tidy to make sure the package is added to your go.mod file.

```bash
go mod init example
go mod tidy
```

Also make sure you initialize this as a git repository so that you can pull the commit hash.

```bash
git init
```

Build your application with the version information.
The variables that you can set are located in the ldflags.go file

For simplicity sake, I am goign to set all the values in the build command, but you can make a better script to set the values from other functions and pass them to the build command.

```bash
go build -o cmd/example/testing -ldflags "-X 'github.com/pragmaticengineering/go-version.BuildVersion=1.0.0' \
-X 'github.com/pragmaticengineering/go-version.ApplicationName="testing"' \
-X 'github.com/pragmaticengineering/go-version.CommitHash=$(git rev-parse HEAD)' \
-X 'github.com/pragmaticengineering/go-version.BuildDate=$(date +%Y%m%d%H%M%S)' \
-X 'github.com/pragmaticengineering/go-version.BuildTag=RELEASE'" \
cmd/example/main.go
```

go ahead and run the application to confirm that the version is being printed correctly with your flag values.
```bash 
./cmd/example/testing
Version: testing v1.0.0 RELEASE (43d01a72cbf53869a18029b86274ce971eaf6c95) linux/amd64 - BuildDate: 20240410015203
```

## Pipeline
For pipelining here is what I would recommend:
1. Create a script that will pull the version of the application and automatically increment the version number when you merge into main. This script will also make sure the Tag is set to RELEASE, or whatever your organization uses.
2. loop through all the folders under the cmd directory for the build cmd and pull the application name from the folder name.
3. Run the build command with the version information.