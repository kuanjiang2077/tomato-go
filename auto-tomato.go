package main

import (
	"fmt"
	"os/exec"
	"time"
)

var workM = 25
var n = 0

var restM = 5

//var long_rest_m = 30

func main() {
	beep(1)
	for {
		n++
		fmt.Println("Now is Round:", (n+1)/2)
		round()
	}
}

func beep(times int) {
	for ; times > 0; times-- {
		cmd := exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/complete.oga")
		cmd.Run()
	}
}

func round() {

	var timer <-chan time.Time
	if n%2 == 1 {
		setTimer(&timer, "working", workM)
	} else {
		setTimer(&timer, "resting", restM)

	}
	<-timer
	beep(3)
}

func setTimer(t *<-chan time.Time, description string, minutes int) {
	fmt.Println(description, minutes, "m")
	*t = time.After(time.Duration(minutes) * time.Minute)

}
