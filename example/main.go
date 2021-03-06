package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/vicanso/elton"

	tracker "github.com/vicanso/elton-tracker"
)

func main() {
	e := elton.New()

	loginTracker := tracker.New(tracker.Config{
		OnTrack: func(info *tracker.Info, _ *elton.Context) {
			buf, _ := json.Marshal(info)
			fmt.Println(string(buf))
		},
	})

	e.Use(func(c *elton.Context) error {
		c.RequestBody = []byte(`{
			"account": "tree.xie",
			"password": "123456"
		}`)
		return c.Next()
	})

	e.POST("/user/login", loginTracker, func(c *elton.Context) (err error) {
		c.SetHeader(elton.HeaderContentType, elton.MIMEApplicationJSON)
		c.BodyBuffer = bytes.NewBuffer(c.RequestBody)
		return
	})

	err := e.ListenAndServe(":3000")
	if err != nil {
		panic(err)
	}
}
