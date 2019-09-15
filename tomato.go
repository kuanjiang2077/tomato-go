package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var n = 0

func main() {
	fmt.Println("type in 'w 25' for work for 25m and 'r 5' for rest 5m")
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
	var s string
	var i int
	fmt.Scanf("%s %d", &s, %i)
	
	switch s {
	case "w": {
		fmt.Printf("working %d minutes\n", i)	
	}
	case "r": {
		fmt.Printf("resting %d minutes\n", i)
	}
	}
	
	
	timer := time.After(time.Duration(i) * time.Second)
	<-timer
	beep(3)
	fmt.Println("finish!")
}
