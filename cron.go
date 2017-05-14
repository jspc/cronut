package main

import (
	"time"

	"github.com/gorhill/cronexpr"
	"github.com/satori/go.uuid"
)

func Cron(j Job) {
	nextTime := cronexpr.MustParse(j.Cron).Next(time.Now())
	ticker := time.NewTicker(time.Second * 30)

	for _ = range ticker.C {
		if nextTime.Sub(time.Now()) <= 0 {
			jid := uuid.NewV4().String()

			L.Printf("%s -> running: %v", jid, j)
			Start(jid, j)
			return
		}
	}
}
