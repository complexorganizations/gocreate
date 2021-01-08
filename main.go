package main

import (
	"io/ioutil"
	"os"
)

// https://github.com/golang-standards/project-layout

func createProject() {
	// take user argument
	argument := os.Args[1]
	// start creating project
	os.Mkdir(argument, 0755)
	// go to the folder
	os.Chdir(argument)
	// create the rest of the folders, files
	os.Mkdir("api", 0755)
	ioutil.WriteFile("./api/README.md", []byte("# `/api`"), 0755)
	os.Mkdir("assets", 0755)
	ioutil.WriteFile("./assets/README.md", []byte("# `/assets`"), 0755)
	os.Mkdir("build", 0755)
	ioutil.WriteFile("./build/README.md", []byte("# `/build`"), 0755)
	os.Mkdir("build/ci", 0755)
	ioutil.WriteFile("./build/ci/README.md", []byte("# `/build/ci`"), 0755)
	os.Mkdir("build/package", 0755)
	ioutil.WriteFile("./build/package/README.md", []byte("# `/build/package`"), 0755)
	os.Mkdir("cmd", 0755)
	ioutil.WriteFile("./cmd/README.md", []byte("# `/cmd`"), 0755)
	os.Mkdir("config", 0755)
	ioutil.WriteFile("./config/README.md", []byte("# `/config`"), 0755)
	os.Mkdir("deployments", 0755)
	ioutil.WriteFile("./deployments/README.md", []byte("# `/deployments`"), 0755)
	os.Mkdir("docs", 0755)
	ioutil.WriteFile("./docs/README.md", []byte("# `/docs`"), 0755)
	os.Mkdir("examples", 0755)
	ioutil.WriteFile("./examples/README.md", []byte("# `/examples`"), 0755)
	os.Mkdir("githooks", 0755)
	ioutil.WriteFile("./githooks/README.md", []byte("# `/githooks`"), 0755)
	os.Mkdir("init", 0755)
	ioutil.WriteFile("./init/README.md", []byte("# `/init`"), 0755)
	os.Mkdir("internal", 0755)
	ioutil.WriteFile("./internal/README.md", []byte("# `/internal`"), 0755)
	os.Mkdir("pkg", 0755)
	ioutil.WriteFile("./pkg/README.md", []byte("# `/pkg`"), 0755)
	os.Mkdir("scripts", 0755)
	ioutil.WriteFile("./scripts/README.md", []byte("# `/scripts`"), 0755)
	os.Mkdir("test", 0755)
	ioutil.WriteFile("./test/README.md", []byte("# `/test`"), 0755)
	os.Mkdir("third_party", 0755)
	ioutil.WriteFile("./third_party/README.md", []byte("# `/third_party`"), 0755)
	os.Mkdir("tools", 0755)
	ioutil.WriteFile("./tools/README.md", []byte("# `/tools`"), 0755)
	os.Mkdir("vendor", 0755)
	ioutil.WriteFile("./vendor/README.md", []byte("# `/vendor`"), 0755)
	os.Mkdir("web", 0755)
	ioutil.WriteFile("./web/README.md", []byte("# `/web`"), 0755)
	os.Mkdir("web/app", 0755)
	ioutil.WriteFile("./web/app/README.md", []byte("# `/web/app`"), 0755)
	os.Mkdir("web/static", 0755)
	ioutil.WriteFile("./web/static/README.md", []byte("# `/web/static`"), 0755)
	os.Mkdir("web/template", 0755)
	ioutil.WriteFile("./web/template/README.md", []byte("# `/web/template`"), 0755)
	os.Mkdir("website", 0755)
	ioutil.WriteFile("./website/README.md", []byte("# `/website`"), 0755)
}

func main() {
	createProject()
	gitignore()
	readmemd()
}
