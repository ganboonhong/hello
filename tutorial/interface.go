package main

import (
    "fmt"
)

type clothWasher struct {
    target string
}

type dishWasher struct {
	target string
}

type Washer interface {
	getCleaner() string
	getWaterVolume() string
}

func startWashing(w Washer) {
	cleaner := w.getCleaner()  
	fmt.Println("start washing with " + cleaner)
}

func stopWashing(w Washer) {
    vol := w.getWaterVolume()  
    fmt.Println("rinsing with " + vol + " liters of water")
}

func (d dishWasher) getCleaner() string {
	return "dishwashing liquid"
}

func (d dishWasher) getWaterVolume() string {
	return "2"
}

func (c clothWasher) getCleaner() string {
	return "detergent"
}

func (c clothWasher) getWaterVolume() string {
	return "10"
}

func main() {
    c := clothWasher{}
	d := dishWasher{}
    startWashing(c) // start washing with detergent
    stopWashing(c) // rinsing with 10 liters of water

    startWashing(d) // start washing with dishwashing liquid
	stopWashing(d) // start washing with dishwashing liquid
}
