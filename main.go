package main

import (
	"os"
	"io/ioutil"
)

func createDirectory() {
	os.Mkdir("api", 0755)
	os.Mkdir("assets", 0755)
	os.Mkdir("build", 0755)
	os.Mkdir("cmd", 0755)
	os.Mkdir("build", 0755)
	os.Mkdir("config", 0755)
	os.Mkdir("deployments", 0755)
	os.Mkdir("docs", 0755)
	os.Mkdir("examples", 0755)
	os.Mkdir("githooks", 0755)
	os.Mkdir("init", 0755)
	os.Mkdir("internal", 0755)
	os.Mkdir("pkg", 0755)
	os.Mkdir("scripts", 0755)
	os.Mkdir("test", 0755)
	os.Mkdir("third_party", 0755)
	os.Mkdir("tools", 0755)
	os.Mkdir("vendor", 0755)
	os.Mkdir("web", 0755)
	os.Mkdir("website", 0755)
}

func createFiles() {
	gitignore := `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/`
	ioutil.WriteFile("./gitignore", []byte(gitignore), 0755)
	ioutil.WriteFile("./api/README.md", []byte("# `/api`"), 0755)
}

func main() {
	createDirectory()
}
