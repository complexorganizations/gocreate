package main

import "io/ioutil"

func gomod() {
	gomod := `module main
	go 1.15`

	ioutil.WriteFile("./go.mod", []byte(gomod), 0755)
}