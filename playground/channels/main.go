package main

import (
	"runtime"
	"time"
	"fmt"
)

type ball struct {
	speed float64
	swingAngle int
}

func (b ball) String() string {
	return fmt.Sprintf("Speed:%f, swingingAngle: %d", b.speed, b.swingAngle)
}

type player struct {
	name string

}

func (p player) bowl(pitch chan ball) {
	pitch <- ball {60, 10}
}

func (p player) bat(pitch chan ball) {
	b := <-pitch 
	fmt.Println(p.name, "playing", b)
}

func main() {
	sachin := player {"sachin"}
	warne := player {"warne"}
	centerPitch := make(chan ball)

	printSystemInfo()
	
	go sachin.bat(centerPitch)
	go warne.bowl(centerPitch)

	// how to keep main waiting without waitGroup?
	time.Sleep(time.Second * 2)
}

func printSystemInfo() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("GO OS:", runtime.GOOS)
	fmt.Println("GO ARCH:", runtime.GOARCH)
}