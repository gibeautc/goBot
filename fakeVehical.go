package main

import (
	"math"
	"time"
)

type fakeVehical struct {
	lat            float64
	lon            float64
	heading        int
	speed          int
	leftMotor      int
	rightMotor     int
	lastUpdateTime time.Time
	maxSpeed       int //mph
	maxTurnRate    int //degrees per sec
}

func NewFakeVehical() *fakeVehical {
	var n *fakeVehical
	n.lat = 44.123
	n.lon = 123.123
	n.speed = 0
	n.heading = 0
	n.lastUpdateTime = time.Now()
	return n
}

func (self *fakeVehical) move(drive int, turn int) {
	//does not have any ramping being done, speed, turning will happen instantly.

	//speed in mph, we need inches moved
	//5280*12/60/60    inches/sec
	disDelta := int64(float64(self.speed) * 17.6 * float64(time.Now().UnixNano()-self.lastUpdateTime.UnixNano()) / 1000 / 1000 / 1000)
	self.lat, self.lon = getCords(self.lat, self.lon, disDelta, self.heading)
	self.speed = int(float64(self.maxSpeed) * (float64(drive) / 100))
	if turn != 0 {
		//turning too
		//speed needs to be decreased by percent of turn as well, and heading needs to change
		//
		self.speed = int(float64(self.speed) * (1 - (math.Abs(float64(turn)))/100))
	}

}

func (self *fakeVehical) getScan() []ScanData {
	lst := make([]ScanData, 0)
	return lst
}

func (self *fakeVehical) sendJson(cmd string) {
	/*
		This is where we would send other data to uC other then motor commands, Lights, turn on sensors, whatever
	*/

}

func (self *fakeVehical) getJson() string {
	/*
		this is where we could get other data from the uC other then scan data, other sensors what what not
	*/
	return ""
}
