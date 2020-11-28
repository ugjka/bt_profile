//go:generate go run gen/main.go

package main

import (
	"flag"
	"os/exec"

	"github.com/getlantern/systray"
)

var sink string
var showQuit bool

func main() {
	s := flag.String("sink", "1", "headset's pulseaudio sink")
	q := flag.Bool("quit", false, "show the quit item")
	flag.Parse()
	sink = *s
	showQuit = *q
	systray.Run(onready, nil)
}

func onready() {
	systray.SetIcon(icon)
	a2dp := systray.AddMenuItem("A2DP", "Switch to A2DP mode")
	hsphfp := systray.AddMenuItem("HSP/HFP", "Switch to HSP/HFP mode")
	quit := systray.AddMenuItem("Quit", "Quit the app")
	if !showQuit {
		quit.Hide()
	}
	for {
		select {
		case <-quit.ClickedCh:
			systray.Quit()
		case <-a2dp.ClickedCh:
			exec.Command("pactl", "set-card-profile", sink, "a2dp_sink").Run()
		case <-hsphfp.ClickedCh:
			exec.Command("pactl", "set-card-profile", sink, "headset_head_unit").Run()
		}
	}
}
