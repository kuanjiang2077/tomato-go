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

var workM = 25
var n = 0

//var rest_m = 5
//var long_rest_m = 30

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

	number := workM
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		text := scanner.Text()
		arr := strings.Split(text, " ")
		switch arr[0] {
		case "w":
			{
				fmt.Printf("working ")
			}
		case "r":
			{
				fmt.Printf("resting ")
			}
		}
		number, _ = strconv.Atoi(arr[1])
		fmt.Println(number, "minutes")
	}
	timer := time.After(time.Duration(number) * time.Minute)
	<-timer
	beep(3)
	fmt.Println("finish!")
}
