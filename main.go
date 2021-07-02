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
	if strings.Contains(projectName, "<") || strings.HasPrefix(projectName, ".") || strings.Contains(projectName, ">") || strings.Contains(projectName, ":") || strings.Contains(projectName, `"`) || strings.Contains(projectName, "/") || strings.Contains(projectName, `\`) || strings.Contains(projectName, "|") || strings.Contains(projectName, "?") || strings.Contains(projectName, "*") || projectName == "." {
		log.Fatalf("Error: %s isn't a legitimate project name.\n", projectName)
	}
	// Check if a repository already exists
	if folderExists(projectName) || fileExists(projectName) {
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
	newContents := strings.Replace(gomod, ("main"), (projectName), -1)
	err = os.WriteFile("go.mod", []byte(newContents), 0)
	handleErrors(err)
	// Create go.sum file
	gosum := ""
	err = os.WriteFile("go.sum", []byte(gosum), 0644)
	handleErrors(err)
	// Create README.md file
	readme := `# [Project_Name]

## Overview

This is a basic Go application project layout. It isn't an official standard set by the Go development team.

### "/assets"

Any data or other component of the environment that enables information-related activities is referred to as an asset.

### "/docs"

Here you'll find design, use, and other project documents.

### "/examples"

Here you'll find examples of your project and library.

### "/scripts"

Scripts for different activities such as building, installing, analyzing, and so on.

### "/test"

This is where your application's data and testing will be stored.

### "/vendor"

The "go mod vendor" command will create the "/vendor" directory for you, which will include all of your application dependencies.`
	// Change from " to `
	changeString := strings.Replace(readme, (`"`), ("`"), -1)
	// Change the project name
	newContents = strings.Replace(changeString, (`[Project_Name]`), (projectName), -1)
	// write the file
	err = os.WriteFile("README.md", []byte(newContents), 0)
	handleErrors(err)
}

// Check if a folder exists
func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// Check if a file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
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
