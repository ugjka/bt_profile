//go:generate go run gen/main.go

package main

import (
	"flag"
	"os/exec"

	"github.com/getlantern/systray"
)

var sink string
var codec string
var showQuit bool

func main() {
	s := flag.String("sink", "1", "headset's pulseaudio sink")
	q := flag.Bool("quit", false, "show the quit item")
	c := flag.String("codec", "", "custom codec e.g. aac, ldac")
	flag.Parse()
	sink = *s
	showQuit = *q
	if *c != "" {
		codec = "_" + *c
	}
	systray.Run(onready, nil)
}

const hifi = "A2DP"
const headset = "HSP/HFP"

func onready() {
	systray.SetIcon(icon)
	a2dp := systray.AddMenuItem(hifi+" 🎧", "Switch to A2DP mode")
	hsphfp := systray.AddMenuItem(headset, "Switch to HSP/HFP mode")
	quit := systray.AddMenuItem("Quit", "Quit the app")
	if !showQuit {
		quit.Hide()
	}
	for {
		select {
		case <-quit.ClickedCh:
			systray.Quit()
		case <-a2dp.ClickedCh:
			exec.Command("pactl", "set-card-profile", sink, "a2dp_sink"+codec).Run()
			a2dp.SetTitle(hifi + " 🎧")
			hsphfp.SetTitle(headset)
		case <-hsphfp.ClickedCh:
			exec.Command("pactl", "set-card-profile", sink, "headset_head_unit").Run()
			hsphfp.SetTitle(headset + " 🎧")
			a2dp.SetTitle(hifi)
		}
	}
}
