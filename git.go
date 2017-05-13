package main

import (
	"io/ioutil"

	"gopkg.in/libgit2/git2go.v25"
)

func Clone() string {
	dir, err := ioutil.TempDir("/tmp", "cronut")
	if err != nil {
		panic(err)
	}

	L.Printf("Cloning %q into %q", *GitRepo, dir)

	_, err = git.Clone(*GitRepo, dir, &git.CloneOptions{
		CheckoutBranch: "master",
	})

	if err != nil {
		panic(err)
	}

	return dir
}
