package main

import (
	"fmt"
	"time"
)

type navLogic interface {
	decide(app *App) (*motorPair, error)
	showType() string
}

type vehicle interface {
	move(drive int, turn int)
	getScan() []ScanData
	sendJson(cmd string)
	getJson() string
}

type Node struct {
	id    int
	osmId int
	lat   float64
	lon   float64
	x     int //local cords
	y     int
}

type Way struct {
	nodes []Node
}

type ScanData struct {
	angle    int
	distance int
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

func stop() *motorPair {
	return new(motorPair)
}
