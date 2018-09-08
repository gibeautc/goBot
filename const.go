package main

import (
	"fmt"
	"time"
)

type location struct {
	id  int     //for OSM style mapping
	lat float32 //global cords
	lon float32
	x   int //local cords
	y   int
}

type navLogic interface {
	decide(app *App) (*motorPair, error)
	showType() string
}

type motorPair struct {
	motor1 int
	motor2 int
	time   time.Time
}

func (self *motorPair) printMotor() {
	if self != nil {
		fmt.Printf("Motor1: %d\n", self.motor1)
		fmt.Printf("Motor2: %d\n", self.motor2)
	} else {
		fmt.Println("Motor Pair is Null")
	}
}

type sensorData struct {
	raw  string
	time time.Time
}

func stop() *motorPair {
	return new(motorPair)
}
