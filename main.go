package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	// forgo error checking for simplicity
	fmt.Println("LED Flash Start")
	const LEDBrightness string = "/sys/class/leds/beaglebone:green:usr0/brightness"
	isLEDOn := 1
	for i := 0; i < 10; i++ {
		cmd := exec.Command("echo", strconv.Itoa(isLEDOn))
		cmd2 := exec.Command("tee", LEDBrightness)
		cmd2.Stdin, _ = cmd.StdoutPipe()
		cmd2.Stdout = os.Stdout
		_ = cmd2.Start()
		_ = cmd.Run()
		_ = cmd2.Wait()
		time.Sleep(1000000 * time.Microsecond)
		isLEDOn ^= 1
	}
	fmt.Println("LED Flash End")
}
