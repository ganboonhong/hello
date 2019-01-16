package reference

import (
    "fmt"
)

type ClothWasher struct {
    target string
}

type DishWasher struct {
	target string
}

type Washer interface {
	GetCleaner() string
	GetWaterVolume() int
}

func StartWashing(w Washer) string {
	cleaner := w.GetCleaner()  
    progress := fmt.Sprintf("start washing with %s" , cleaner)
	return progress
}

func StopWashing(w Washer) string {
    vol := w.GetWaterVolume()  
    progress := fmt.Sprintf("rinsing with %d liters of water" , vol)
    return progress
}

func (d DishWasher) GetCleaner() string {
	return "dishwashing liquid"
}

func (d DishWasher) GetWaterVolume() int {
	return 2
}

func (c ClothWasher) GetCleaner() string {
	return "detergent"
}

func (c ClothWasher) GetWaterVolume() int {
	return 10
}

func main() {
    c := ClothWasher{}
	d := DishWasher{}
    fmt.Println(StartWashing(c)) // start washing with detergent
    fmt.Println(StopWashing(c)) // rinsing with 10 liters of water

    fmt.Println(StartWashing(d)) // start washing with dishwashing liquid
	fmt.Println(StopWashing(d)) // rinsing with 2 liters of water
}
