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
	os.Mkdir(fmt.Sprint(projectName+"/assets"), 0755)
	os.WriteFile(fmt.Sprint(projectName+"/assets/README.md"), []byte("### `/assets`"), 0644)
	// Create assets folder
	os.Mkdir(fmt.Sprint(projectName+"/cmd"), 0755)
	os.WriteFile(fmt.Sprint(projectName+"/cmd/README.md"), []byte("### `/cmd`"), 0644)
	// Create build folder
	os.Mkdir(fmt.Sprint(projectName+"/build"), 0755)
	os.WriteFile(fmt.Sprint(projectName+"/build/README.md"), []byte("### `/build`"), 0644)
	// Create pkg folder
	os.Mkdir(fmt.Sprint(projectName+"/pkg"), 0755)
	os.WriteFile(fmt.Sprint(projectName+"/pkg/README.md"), []byte("### `/pkg`"), 0644)
	// Create internal folder
	os.Mkdir(fmt.Sprint(projectName+"/internal"), 0755)
	os.WriteFile(fmt.Sprint(projectName+"/internal/README.md"), []byte("### `/internal`"), 0644)
	// Create scripts folder
	os.Mkdir(fmt.Sprint(projectName+"/scripts"), 0755)
	os.WriteFile(fmt.Sprint(projectName+"/scripts/README.md"), []byte("### `/scripts`"), 0644)
	// Create vendor folder
	os.Mkdir(fmt.Sprint(projectName+"/vendor"), 0755)
	os.WriteFile(fmt.Sprint(projectName+"/vendor/README.md"), []byte("### `/vendor`"), 0644)
	// Create main.go file
	main := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}`
	os.WriteFile(fmt.Sprint(projectName+"/main.go"), []byte(main), 0644)
	// Let's make a go.mod file and name it after the project.
	gomod := `module [Project_Name]

go 1.17`
	newContents := strings.Replace(gomod, ("[Project_Name]"), (projectName), -1)
	os.WriteFile(fmt.Sprint(projectName+"/go.mod"), []byte(newContents), 0)
	// Create the go.sum file, but keep it blank because we don't have any dependencies.
	os.WriteFile(fmt.Sprint(projectName+"/go.sum"), []byte(""), 0644)
	// The README.md file's contents
	readme := `# [Project_Name]`
	// Let's change the string to the name of the project.
	newContents = strings.Replace(readme, (`[Project_Name]`), (projectName), -1)
	// Let's create a readme file for the entire repository.
	os.WriteFile(fmt.Sprint(projectName+"/README.md"), []byte(newContents), 0)
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
