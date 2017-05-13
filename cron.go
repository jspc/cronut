package main

import (
	"io/ioutil"
	"time"

	"github.com/gorhill/cronexpr"
	"github.com/satori/go.uuid"
)

func Cron(p string) {
	dat, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	j, err := Parse(dat)
	if err != nil {
		panic(err)
	}

	for {
		nextTime := cronexpr.MustParse(j.Cron).Next(time.Now())
		ticker := time.NewTicker(time.Second * 30)

		for _ = range ticker.C {
			if nextTime.Sub(time.Now()) <= 0 {
				jid := uuid.NewV4().String()

				L.Printf("%s -> running: %v", jid, j)
				Start(jid, j)
				break
			}
		}
	}
}
