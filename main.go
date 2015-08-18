package main

import (
	"fmt"
	"io"

	"os/exec"
	"runtime"

	"github.com/thefryscorer/gocontroller"
)

var cmdPipe io.WriteCloser

func keypress(key string) {
	cmdPipe.Write([]byte(fmt.Sprintf("keydown %v\n", key)))
}

func release(key string) {
	cmdPipe.Write([]byte(fmt.Sprintf("keyup %v\n", key)))
}

func main() {
	xte := exec.Command("xte")
	var err error
	cmdPipe, err = xte.StdinPipe()
	if err != nil {
		panic(err)
	}
	defer cmdPipe.Close()

	if err := xte.Start(); err != nil {
		panic(err)
	}

	runtime.GOMAXPROCS(8)
	layout := gocontroller.Layout{Style: gocontroller.DefaultCSS, Buttons: []gocontroller.Button{
		{Left: 20, Top: 20, Key: "Up"},
		{Left: 20, Top: 60, Key: "Down"},
		{Left: 10, Top: 40, Key: "Left"},
		{Left: 30, Top: 40, Key: "Right"},
		{Left: 70, Top: 20, Key: "w", Color: "#f92672"},
		{Left: 60, Top: 40, Key: "a", Color: "#82b414"},
		{Left: 80, Top: 40, Key: "d", Color: "#56c2d6"},
		{Left: 70, Top: 60, Key: "s", Color: "#8c54fe"},
		{Left: 0, Top: 0, Key: "l", Label: "L"},
		{Left: 95, Top: 0, Key: "r", Label: "R"},
		{Left: 35, Top: 10, Key: "Escape"},
		{Left: 55, Top: 10, Key: "Return"},
	}}
	server := gocontroller.NewServer(layout, gocontroller.DefaultPort)
	server.Start()
	fmt.Println("Server started on port: " + gocontroller.DefaultPort)
	inAgg := server.NewInputAggregator()
	for {
		inAgg.Collect()
		for _, in := range inAgg.Inputs {
			if in.Event == gocontroller.PRESS {
				keypress(in.Key)
			} else if in.Event == gocontroller.RELEASE {
				release(in.Key)
			}
		}

		//Clear inputs
		inAgg.Clear()

	}
}
