package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment)

	for _, file := range files {
		envName := strings.ReplaceAll(file.Name(), "=", "")

		if file.Size() == 0 {
			env[envName] = EnvValue{NeedRemove: true}
			continue
		}

		value, err := readFile(path.Join(dir, file.Name()))
		if err != nil || value == "" {
			env[envName] = EnvValue{NeedRemove: true}
			continue
		}

		env[envName] = EnvValue{Value: value, NeedRemove: false}
	}
	return env, nil
}

func readFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	r := bufio.NewReader(file)

	line, _, err := r.ReadLine()
	if err != nil {
		return "", err
	}

	line = bytes.ReplaceAll(line, []byte{0x00}, []byte("\n"))
	line = bytes.TrimRight(line, " \t")

	return string(line), nil
}
