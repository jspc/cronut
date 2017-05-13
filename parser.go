package main

import (
	"gopkg.in/yaml.v2"
)

// Job provides a deserialized joburation document
type Job struct {
	// Name is the name of the job
	Name string
	// Env is a list of environment vars to apply
	Env []string
	// Args is a string containing arguments to the command
	Args string
}

func Parse(f []byte) (c Job, err error) {
	err = yaml.Unmarshal([]byte(f), &c)
	return
}
