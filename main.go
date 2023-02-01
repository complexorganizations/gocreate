package main

import (
	"log"
	"os"
	"strings"
)

var (
	projectName string
)

func init() {
	// Ascertain that the user has provided a project name.
	if len(os.Args) < 1 {
		log.Fatalln("Error: The project name was not been given.")
	} else {
		projectName = os.Args[1]
	}
	// Check to see if the folder or file you're looking for already exists.
	if folderExists(projectName) || fileExists(projectName) {
		log.Fatalf("Error: Failed to create %s directory.\n", projectName)
	}
}

func main() {
	createProjectStructure()
}

// Create Project Structure
func createProjectStructure() {
	// Create project folder
	createFolder(projectName)
	// Create assets folder
	createFolder(projectName + "/assets")
	writeToFile(projectName+"/assets/README.md", "### `/assets`")
	// Create assets folder
	createFolder(projectName + "/cmd")
	writeToFile(projectName+"/cmd/README.md", "### `/cmd`")
	// Create build folder
	createFolder(projectName + "/build")
	writeToFile(projectName+"/build/README.md", "### `/build`")
	// Create pkg folder
	createFolder(projectName + "/pkg")
	writeToFile(projectName+"/pkg/README.md", "### `/pkg`")
	// Create internal folder
	createFolder(projectName + "/internal")
	writeToFile(projectName+"/internal/README.md", "### `/internal`")
	// Create scripts folder
	createFolder(projectName + "/scripts")
	writeToFile(projectName+"/scripts/README.md", "### `/scripts`")
	// Create vendor folder
	createFolder(projectName + "/vendor")
	writeToFile(projectName+"/vendor/README.md", "### `/vendor`")
	// Create main.go file
	main := `package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}`
	writeToFile(projectName+"/main.go", main)
	// Let's make a go.mod file and name it after the project.
	gomod := `module [Project_Name]

go 1.17`
	newContents := strings.Replace(gomod, ("[Project_Name]"), (projectName), -1)
	writeToFile(projectName+"/go.mod", newContents)
	// Create the go.sum file, but keep it blank because we don't have any dependencies.
	writeToFile(projectName+"/go.sum", "")
	// The README.md file's contents
	readme := `# [Project_Name]`
	// Let's change the string to the name of the project.
	newContents = strings.Replace(readme, (`[Project_Name]`), (projectName), -1)
	// Let's create a readme file for the entire repository.
	writeToFile(projectName+"/README.md", newContents)
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

// Don't append and write to file
func writeToFile(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

// Create a folder in the system.
func createFolder(foldername string) {
	err := os.MkdirAll(foldername, 0755)
	if err != nil {
		log.Fatalf("Error: Failed to create %s directory.\n", foldername)
	}
}
