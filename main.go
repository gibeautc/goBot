package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var err error
	app := new(App)
	app.setup()
	fmt.Println("Starting up")
	toUc := make(chan motorPair, 100)
	fromUc := make(chan sensorData, 100)
	go SerialThread(toUc, fromUc)
	for true {

		if time.Now().Unix()-app.lastPrint.Unix() > 1 {
			app.printApp()
			app.lastPrint = time.Now()
		}
		//check that channels are not being filled up
		if len(toUc) > 20 {
			fmt.Println("Channel toUc Length to high: ", len(toUc))
		}
		if len(fromUc) > 20 {
			fmt.Println("Channel fromUc Length to high: ", len(toUc))
		}
		time.Sleep(5 * time.Millisecond)
		app.curMotor.motor2 = 50
		app.curMotor.motor1 = 50
		app.curMotor.time = time.Now()
		toUc <- *app.curMotor
		select {
		case msg := <-fromUc:
			if msg.raw == "connected" {
				fmt.Println("Serial Connected!")
				app.serialConnected = true
				continue
			} else if msg.raw == "disconnected" {
				app.serialConnected = false
				continue
			}
			if !app.serialConnected {
				fmt.Println("Serial Disconnected!!!!!")
				time.Sleep(1 * time.Second)
				continue
			}
			delta := time.Now().UnixNano() - msg.time.UnixNano()
			delta = delta / 1000 / 1000
			if int(delta) > app.stale {
				fmt.Println("Stale data,ignoring")
				continue
			}
			app.curSensor = &msg
		}
		app.primaryDecision, err = app.brain.decide(app)
		if err != nil {
			fmt.Println("Brain Error")
			fmt.Println(err.Error())
		}
	}
}

func SerialThread(msgQueue chan motorPair, sensorQueue chan sensorData) {
	//Launched as go routine. msgQueue is where main sends motor values to be sent to uC
	//sensorQueue is where we put sensor values received from uC to go back to main

	//connect to serial

	//once connected, send msg back to main
	conResp := new(sensorData)
	conResp.raw = "connected"
	sensorQueue <- *conResp
	for true {
		//need to make sure these are non-blocking
		msg := <-msgQueue
		//out:=strconv.Itoa(msg.motor1)+","+strconv.Itoa(msg.motor2)
		_ = strconv.Itoa(msg.motor1) + "," + strconv.Itoa(msg.motor2)
		//serial.send out
		//read serial till newline or timeout (1 sec?)
		//actually receive this data from uC
		rec := "{\"bump1\":false,\"bump2\":false,\"bump3\":false,\"bump4\":false,\"sonar1\":10,\"sonar2\":15}"
		resp := new(sensorData)
		resp.raw = rec
		resp.time = time.Now()
		sensorQueue <- *resp
	}
}
