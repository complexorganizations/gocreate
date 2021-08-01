package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	projectName string
	err         error
)

func init() {
	// Ascertain that the user has provided a project name.
	if len(os.Args) < 1 {
		log.Fatal("Error: The project name was not been given.")
	} else {
		projectName = os.Args[1]
	}
	// If no name is specified or if the name is the default, exit.
	if len(projectName) < 1 || projectName == "/user/example/folder" {
		log.Fatal("Error: The name of the project has not been given.")
	}
	// Make sure the project name doesn't contain any characters that aren't allowed.
	if strings.Contains(projectName, "<") || strings.HasPrefix(projectName, ".") || strings.Contains(projectName, ">") || strings.Contains(projectName, ":") || strings.Contains(projectName, `"`) || strings.Contains(projectName, "/") || strings.Contains(projectName, `\`) || strings.Contains(projectName, "|") || strings.Contains(projectName, "?") || strings.Contains(projectName, "*") || projectName == "." {
		log.Fatalf("Error: %s isn't a legitimate project name.\n", projectName)
	}
	// Check to see if the folder or file you're looking for already exists.
	if folderExists(projectName) || fileExists(projectName) {
		log.Fatalf("Error: Failed to create %s directory.\n", projectName)
	}
	// Make sure we have git installed in the system.
	commandExists("git")
}

func main() {
	createProjectStructure()
}

// Create Project Structure
func createProjectStructure() {
	// Create project folder
	err = os.Mkdir(projectName, 0755)
	if err != nil {
		log.Fatalf("Error: Failed to create %s directory.\n", projectName)
	}
	os.Chdir(projectName)
	// Create assets folder
	os.Mkdir("assets", 0755)
	os.WriteFile("assets/README.md", []byte("### `/assets`"), 0644)
	// Create docs folder
	os.Mkdir("docs", 0755)
	os.WriteFile("docs/README.md", []byte("### `/docs`"), 0644)
	// Create examples folder
	os.Mkdir("examples", 0755)
	os.WriteFile("examples/README.md", []byte("### `/examples`"), 0644)
	// Create scripts folder
	os.Mkdir("scripts", 0755)
	os.WriteFile("scripts/README.md", []byte("### `/scripts`"), 0644)
	// Create test folder
	os.Mkdir("test", 0755)
	os.WriteFile("test/README.md", []byte("### `/test`"), 0644)
	// Create vendor folder
	os.Mkdir("vendor", 0755)
	os.WriteFile("vendor/README.md", []byte("### `/vendor`"), 0644)
	// Create main.go file
	main := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}`
	os.WriteFile("main.go", []byte(main), 0644)
	// Let's make a go.mod file and name it after the project.
	gomod := `module [Project_Name]

go 1.17`
	newContents := strings.Replace(gomod, ("[Project_Name]"), (projectName), -1)
	os.WriteFile("go.mod", []byte(newContents), 0)
	// Create the go.sum file, but keep it blank because we don't have any dependencies.
	os.WriteFile("go.sum", []byte(""), 0644)
	// The README.md file's contents
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
	changeString := strings.Replace(readme, (`"`), ("`"), -1)
	// Let's change the string to the name of the project.
	newContents = strings.Replace(changeString, (`[Project_Name]`), (projectName), -1)
	// Let's create a readme file for the entire repository.
	os.WriteFile("README.md", []byte(newContents), 0)
	// Init the repo and add the gitignore file.
	exec.Command("git", "init").Run()
	gitignore := `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with go test -c
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/`
	os.WriteFile(".gitignore", []byte(gitignore), 0644)
}

// Check to see if a folder already exists.
func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// Check to see whether a file already exists.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// Check to see if the program has been installed and added to the path.
func commandExists(cmd string) {
	cmd, err := exec.LookPath(cmd)
	if err != nil {
		log.Fatalf("Error: The application %s was not found in the system.\n", cmd)
	}
}
