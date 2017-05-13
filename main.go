package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/moby/moby/client"
	"github.com/zeebox/go-http-middleware"
	"github.com/zeebox/goose4"
)

var (
	GitRepo = flag.String("repo", "", "Git repo which stores job configuration")
	L       = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	Docker *client.Client
)

func main() {
	flag.Parse()
	L.Print("cronut")
	L.Printf("Initialized with repo: %q", *GitRepo)

	c := goose4.Config{
		ArtifactID:      "cronut",
		BuildNumber:     "n/a",
		BuildMachine:    "",
		BuiltBy:         "",
		BuiltWhen:       time.Now(),
		CompilerVersion: "",
		GitSha:          "",
		RunbookURI:      "",
		Version:         "v0.0.1",
	}

	se4, err := goose4.NewGoose4(c)
	if err != nil {
		panic(err)
	}

	Docker, err = client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	go func() {
		checkoutDir := Clone()
		//defer os.RemoveAll(checkoutDir)

		gitDir, err := os.Open(checkoutDir)
		if err != nil {
			panic(err)
		}

		dirContents, err := gitDir.Readdir(0)
		if err != nil {
			panic(err)
		}

		for _, f := range dirContents {
			if strings.HasSuffix(f.Name(), ".conf") {
				L.Printf("Loading: %q", f.Name())

				go Cron(path.Join(checkoutDir, f.Name()))
			}
		}
	}()

	http.Handle("/service/", middleware.NewMiddleware(se4))
	panic(http.ListenAndServe(":8080", nil))
}
