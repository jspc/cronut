package main

import (
	"gopkg.in/yaml.v2"
)

// Job provides a deserialized job document
type Job struct {
	// Name is the name of the job
	Name string
	// Env is a list of environment vars to apply
	Env []string
	// Args is a list containing arguments to the command
	Args []string
	// Cron is a string representing a cron time string thing
	Cron string
	// Container is the container which is run
	Container string
}

func Parse(f []byte) (j Job, err error) {
	err = yaml.Unmarshal([]byte(f), &j)
	return
}
