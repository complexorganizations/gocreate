package main

import (
	"os"
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
}

func main() {
	createDirectory()
}
