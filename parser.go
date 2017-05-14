package main

import (
	"io/ioutil"

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

func Parse(path string) (j Job, err error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(dat, &j)
	return
}
