package main

import (
	"errors"
	"fmt"
	"github.com/vcraescu/go-xrandr"
	"github.com/weakish/goaround"
	"os/exec"
)

func getPrimaryMonitor(screen xrandr.Screen) (xrandr.Monitor, error) {
	for _, monitor := range screen.Monitors {
		if monitor.Primary && monitor.Connected {
			return monitor, nil
		}
	}
	return xrandr.Monitor{}, errors.New("no primary monitor connected")
}

func main() {
	screens, err := xrandr.GetScreens()
	goaround.FatalIf(err)
	screen := screens[0]

	primaryMonitor, err := getPrimaryMonitor(screen)
	monitorID := primaryMonitor.ID
	widthPX := int(primaryMonitor.Resolution.Width)
	heightPX := int(primaryMonitor.Resolution.Height)
	widthMM := int(primaryMonitor.Size.Width)
	heightMM := int(primaryMonitor.Size.Height)

	halfWidthPX := widthPX / 2
	halfWidthMM := widthMM / 2
	leftHalf := fmt.Sprintf("xrandr --setmonitor left %d/%dx%d/%d+0+0 %s",
		halfWidthPX, halfWidthMM, heightPX, heightMM, monitorID)
	rightHalf := fmt.Sprintf("xrandr --setmonitor right %d/%dx%d/%d+%d+0 none",
		halfWidthPX, halfWidthMM, heightPX, heightMM, halfWidthPX)
	err = exec.Command(leftHalf).Run()
	goaround.FatalIf(err)
	err = exec.Command(rightHalf).Run()
	goaround.FatalIf(err)
}
