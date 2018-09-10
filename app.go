package main

import (
	"fmt"
	"time"
)

type App struct {
	//main token to pass around
	stale     int  //number of ms we concider a messge to be stale and not process it.
	collision bool //have we just hit something?  Processing and movements will be slower if this has happened
	lastPrint time.Time
	/*
		brain will store whatever interface we are currently using.
		Most of the time this would be static, whatever type of navigation I am working on, but I also
		think it could be dynamic, such as using lineFollower when on a road to stay in the bike lane
		or bumpLogic in a dark environment
	*/
	brain           navLogic
	primaryDecision *motorPair
	serialConnected bool
	veh             vehicle
}

func (self *App) setup() {
	self.stale = 50 //ms
	self.veh = NewFakeVehical()
	self.collision = false
	self.lastPrint = time.Now()
	self.brain = new(bumpLogic)
	self.serialConnected = false

}

func (self *App) mix() {
	/*
		Not sure yet how we want to decide between each decision, maybe each of the logics should have a percentage on how
		confident it is?
	*/

}

func (self *App) printApp() {
	fmt.Println("************System Status************")
	fmt.Printf("Stale setting(ms): %d\n", self.stale)
	fmt.Printf("In Collision State: %t\n", self.collision)
	fmt.Println("Current Motor Settings:")
	fmt.Printf("Nav Type: %s\n", self.brain.showType())
	fmt.Println("Primary Decision:")
	self.primaryDecision.printMotor()
}

func (self *App) run() {

}
