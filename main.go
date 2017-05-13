package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/zeebox/go-http-middleware"
	"github.com/zeebox/goose4"
)

var (
	GitRepo = flag.String("repo", "", "Git repo which stores job configuration")
	L       = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
)

func main() {
	flag.Parse()

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

	http.Handle("/service/", middleware.NewMiddleware(se4))
	panic(http.ListenAndServe(":8080", nil))
}
