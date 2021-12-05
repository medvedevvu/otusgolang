package main

import (
	"log"
	"os"
)

func main() {
	// Place your code here.
	args := os.Args
	if len(args) < 3 {
		log.Fatal("you must define at least 2 arguments, path and execute command")
	}

	env, err := ReadDir(args[1])
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(RunCmd(args[2:], env))
}
