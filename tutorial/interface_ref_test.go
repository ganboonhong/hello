package washer

import (
    "testing"
)

func TestStartWashing(t *testing.T){
    c := ClothWasher{}
    result := StartWashing(c)
    expected := "start washing with detergent"

    if result != expected {
        t.Errorf("expected: %s, result: %s", expected, result)
    }
}

func TestStopWashing(t *testing.T){
    c := ClothWasher{}
    result := StopWashing(c)
    expected := "rinsing with 10 liters of water"
    
    if result != expected {
        t.Errorf("expected: %s, result: %s", expected, result)
    }
}