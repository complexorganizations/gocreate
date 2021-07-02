package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

var (
	projectName string
	err         error
)

func init() {
	if len(os.Args) > 1 {
		// Supported Flags
		tempProjectName := flag.String("name", "/user/example/folder", "The name of the go project.")
		flag.Parse()
		projectName = *tempProjectName
	} else {
		log.Fatal("Error: The project name was not been given.")
	}
	// Project name empty
	if len(projectName) < 1 || projectName == "/user/example/folder" {
		log.Fatal("Error: The name of the project has not been given.")
	}
	// Invalid stuff
	if strings.Contains(projectName, "<") || strings.Contains(projectName, ">") || strings.Contains(projectName, ":") || strings.Contains(projectName, `"`) || strings.Contains(projectName, "/") || strings.Contains(projectName, `\`) || strings.Contains(projectName, "|") || strings.Contains(projectName, "?") || strings.Contains(projectName, "*") {
		log.Fatalf("Error: %s isn't a legitimate project name.\n", projectName)
	}
	// Check if a repository already exists
	if folderExists(projectName) {
		log.Fatalf("Error: Failed to create %s directory.\n", projectName)
	} else if fileExists(projectName) {
		log.Fatalf("Error: Failed to create %s directory.\n", projectName)
	}
}

func main() {
	createProjectStructure()
	createProjectFiles()
}

// Create Project Structure
func createProjectStructure() {
	// Create project folder
	err = os.Mkdir(projectName, 0755)
	handleErrors(err)
	err = os.Chdir(projectName)
	handleErrors(err)
	// Create assets folder
	err = os.Mkdir("assets", 0755)
	handleErrors(err)
	err = os.WriteFile("assets/README.md", []byte("### `/assets`"), 0644)
	handleErrors(err)
	// Create docs folder
	err = os.Mkdir("docs", 0755)
	handleErrors(err)
	err = os.WriteFile("docs/README.md", []byte("### `/docs`"), 0644)
	handleErrors(err)
	// Create examples folder
	err = os.Mkdir("examples", 0755)
	handleErrors(err)
	err = os.WriteFile("examples/README.md", []byte("### `/examples`"), 0644)
	handleErrors(err)
	// Create scripts folder
	err = os.Mkdir("scripts", 0755)
	handleErrors(err)
	err = os.WriteFile("scripts/README.md", []byte("### `/scripts`"), 0644)
	handleErrors(err)
	// Create test folder
	err = os.Mkdir("test", 0755)
	handleErrors(err)
	err = os.WriteFile("test/README.md", []byte("### `/test`"), 0644)
	handleErrors(err)
	// Create vendor folder
	err = os.Mkdir("vendor", 0755)
	handleErrors(err)
	err = os.WriteFile("vendor/README.md", []byte("### `/vendor`"), 0644)
	handleErrors(err)
}

// Create Project Files
func createProjectFiles() {
	// Create main.go file
	main := `package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}`
	err = os.WriteFile("main.go", []byte(main), 0644)
	handleErrors(err)
	// Create go.mod file
	gomod := `module main

go 1.17`
	err = os.WriteFile("go.mod", []byte(gomod), 0644)
	handleErrors(err)
	read, err := os.ReadFile("go.mod")
	handleErrors(err)
	newContents := strings.Replace(string(read), ("main"), (projectName), -1)
	err = os.WriteFile("go.mod", []byte(newContents), 0)
	handleErrors(err)
	// Create go.sum file
	gosum := ""
	err = os.WriteFile("go.sum", []byte(gosum), 0644)
	handleErrors(err)
	// Create README.md file
	readme := `# Standard Go Project Layout

## Overview

This is a basic layout for Go application projects. It's not an official standard defined by the core Go dev team.

### "/assets"

Other assets to go along with your repository (images, logos, etc).

### "/docs"

Design and user documents (in addition to your godoc generated documentation).

See the ["/docs"](docs/README.md) directory for examples.

### "/examples"

Examples for your applications and/or public libraries.

See the ["/examples"](examples/README.md) directory for examples.

### "/scripts"

Scripts to perform various build, install, analysis, etc operations.

These scripts keep the root level Makefile small and simple (e.g., ["https://github.com/hashicorp/terraform/blob/main/Makefile"](https://github.com/hashicorp/terraform/blob/main/Makefile)).

See the ["/scripts"](scripts/README.md) directory for examples.

### "/test"

Additional external test apps and test data. Feel free to structure the "/test" directory anyway you want. For bigger projects it makes sense to have a data subdirectory. For example, you can have "/test/data" or "/test/testdata" if you need Go to ignore what's in that directory. Note that Go will also ignore directories or files that begin with "." or "_", so you have more flexibility in terms of how you name your test data directory.

See the ["/test"](test/README.md) directory for examples.

### "/vendor"

Application dependencies (managed manually or by your favorite dependency management tool like the new built-in ["Go Modules"](https://github.com/golang/go/wiki/Modules) feature). The "go mod vendor" command will create the "/vendor" directory for you. Note that you might need to add the "-mod=vendor" flag to your "go build" command if you are not using Go 1.14 where it's on by default.

Don't commit your application dependencies if you are building a library.

Note that since ["1.13"](https://golang.org/doc/go1.13#modules) Go also enabled the module proxy feature (using ["https://proxy.golang.org"](https://proxy.golang.org) as their module proxy server by default). Read more about it ["here"](https://blog.golang.org/module-mirror-launch) to see if it fits all of your requirements and constraints. If it does, then you won't need the "vendor" directory at all.`
	err = os.WriteFile("README.md", []byte(readme), 0644)
	handleErrors(err)
	read, err = os.ReadFile("README.md")
	handleErrors(err)
	newContents = strings.Replace(string(read), (`"`), ("`"), -1)
	err = os.WriteFile("README.md", []byte(newContents), 0)
	handleErrors(err)
}

// Check if a folder exists
func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// Check if a file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Log errors
func handleErrors(err error) {
	if err != nil {
		log.Println(err)
	}
}
