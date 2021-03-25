package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"time"
	"strconv"
	"strings"
)

const (
	PREFIX = "/sys/class/power_supply/BAT0"
	MIN = 60
	MAX = 80
)

func readString(fn string) (string, error) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func readInt(fn string) (int, error) {
	data, err := readString(fn)
	if err != nil {
		return 0, err
	}
	num, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(num), nil
}

func levelToThresh(level int, min int, max int) int {
	if level < min {
		return max
	}
	return min
}

func main() {
	status, err := readString(PREFIX + "/status")
	if err != nil {
		os.Exit(1)
	}
	level, err := readInt(PREFIX + "/capacity")
	if err != nil {
		os.Exit(1)
	}
	thresh, err := readInt(PREFIX + "/charge_control_end_threshold")
	if err != nil {
		os.Exit(1)
	}
	newThresh := 0
	if (thresh == 10) || (status == "Discharging") {
		newThresh = levelToThresh(level, MIN, MAX)
	}
	fmt.Println(status, level, thresh, newThresh)
	time.Sleep(2 * time.Second)
}
