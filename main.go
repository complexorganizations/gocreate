package main

import (
	"fmt"
	"log"
	"os"
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
	if len(projectName) < 1 {
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
	// Create assets folder
	createFolder(projectName + "/assets")
	writeToFile(fmt.Sprint(projectName+"/assets/README.md"), "### `/assets`")
	// Create assets folder
	createFolder(projectName + "/cmd")
	writeToFile(fmt.Sprint(projectName+"/cmd/README.md"), "### `/cmd`")
	// Create build folder
	createFolder(projectName + "/build")
	writeToFile(fmt.Sprint(projectName+"/build/README.md"), "### `/build`")
	// Create pkg folder
	createFolder(projectName + "/pkg")
	writeToFile(fmt.Sprint(projectName+"/pkg/README.md"), "### `/pkg`")
	// Create internal folder
	createFolder(projectName + "/internal")
	writeToFile(fmt.Sprint(projectName+"/internal/README.md"), "### `/internal`")
	// Create scripts folder
	createFolder(projectName + "/scripts")
	writeToFile(fmt.Sprint(projectName+"/scripts/README.md"), "### `/scripts`")
	// Create vendor folder
	createFolder(projectName + "/vendor")
	writeToFile(fmt.Sprint(projectName+"/vendor/README.md"), "### `/vendor`")
	// Create main.go file
	main := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}`
	writeToFile(fmt.Sprint(projectName+"/main.go"), main)
	// Let's make a go.mod file and name it after the project.
	gomod := `module [Project_Name]

go 1.17`
	newContents := strings.Replace(gomod, ("[Project_Name]"), (projectName), -1)
	writeToFile(fmt.Sprint(projectName+"/go.mod"), newContents)
	// Create the go.sum file, but keep it blank because we don't have any dependencies.
	writeToFile(fmt.Sprint(projectName+"/go.sum"), "")
	// The README.md file's contents
	readme := `# [Project_Name]`
	// Let's change the string to the name of the project.
	newContents = strings.Replace(readme, (`[Project_Name]`), (projectName), -1)
	// Let's create a readme file for the entire repository.
	writeToFile(fmt.Sprint(projectName+"/README.md"), newContents)
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

// Append and write to file
func writeToFile(pathInSystem string, content string) {
	filePath, err := os.OpenFile(pathInSystem, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	_, err = filePath.WriteString(content + "\n")
	if err != nil {
		log.Println(err)
	}
	err = filePath.Close()
	if err != nil {
		log.Println(err)
	}
}

// Create a folder in the system.
func createFolder(foldername string) {
	err = os.Mkdir(foldername, 0755)
	if err != nil {
		log.Printf("Error: Failed to create %s directory.\n", foldername)
	}
}
