package main

type fakeVehical struct {
	lat     float64
	lon     float64
	heading int
	speed   int
}

func (self *fakeVehical) move(drive int, turn int) {
	/*
		tell the vehical to move
		drive is forward and backword, turn is degrees left or right + or -
	*/
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
