package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)



func main(){
	var err error
	app:=new(App)
	app.setup()
	fmt.Println("Starting up")
	toUc:=make(chan motorPair,100)
	fromUc:=make(chan sensorData,100)
	go SerialThread(toUc,fromUc)

	for true{

		if time.Now().Unix()-app.lastPrint.Unix()>1{
			app.printApp()
			app.lastPrint=time.Now()
		}
		//check that channels are not being filled up
		if len(toUc)>20{
			fmt.Println("Channel toUc Length to high: ",len(toUc))
		}
		if len(fromUc)>20{
			fmt.Println("Channel fromUc Length to high: ",len(toUc))
		}
		time.Sleep(5*time.Millisecond)
		app.curMotor.motor2=50
		app.curMotor.motor1=50
		app.curMotor.time=time.Now()
		toUc<-*app.curMotor
		select {
		case msg := <-fromUc:
			//fmt.Println("received message", msg)
			delta:=time.Now().UnixNano()-msg.time.UnixNano()
			delta=delta/1000/1000
			if int(delta)>app.stale{
				fmt.Println("Stale data,ignoring")
				return
			}
			if msg.bump1 || msg.bump2 || msg.bump3 || msg.bump4{
				app.collision=true
			}
			app.curSensor.bump1=msg.bump1
			app.curSensor.bump2=msg.bump2
			app.curSensor.bump3=msg.bump3
			app.curSensor.bump4=msg.bump4
			app.curSensor.sonar1=msg.sonar1
			app.curSensor.sonar2=msg.sonar2
		default:
			fmt.Println("no message received")
		}
		app.primaryDecision,err=app.primaryBrain.decide(app)
		if err!=nil{
			fmt.Println("Primary Error")
			fmt.Println(err.Error())
		}
		if app.secondaryBrain!=nil{
			app.secondaryDecision,err=app.secondaryBrain.decide(app)
			if err!=nil{
				fmt.Println("Secondary Error")
				fmt.Println(err.Error())
			}
		}

		if app.tertiaryBrain!=nil{
			app.tertiaryDecision,err=app.tertiaryBrain.decide(app)
			if err!=nil{
				fmt.Println("Tertiary Error")
				fmt.Println(err.Error())
			}
		}
		if app.primaryAssert{
			app.curMotor=app.primaryDecision
		}else{
			app.mix()
		}
	}


}


func SerialThread(msgQueue chan motorPair,sensorQueue chan sensorData){
	//Launched as go routine. msgQueue is where main sends motor values to be sent to uC
	//sensorQueue is where we put sensor values received from uC to go back to main

	//connect to serial


	for true{
		//need to make sure these are non-blocking
		msg := <- msgQueue
		//out:=strconv.Itoa(msg.motor1)+","+strconv.Itoa(msg.motor2)
		_=strconv.Itoa(msg.motor1)+","+strconv.Itoa(msg.motor2)
		//serial.send out
		//read serial till newline or timeout (1 sec?)
		var resp sensorData
		var err error
		//actually receive this data from uC
		rec:="F,F,F,F,10,15"
		elem:=strings.Split(rec,",")
		if len(elem)!=6{
			fmt.Println("Error in data")
			fmt.Println(elem)
			continue
		}
		resp.time=time.Now()
		resp.bump1=elem[0]=="T"
		resp.bump2=elem[1]=="T"
		resp.bump3=elem[2]=="T"
		resp.bump4=elem[3]=="T"
		resp.sonar1,err=strconv.Atoi(elem[4])
		if err!=nil{
			fmt.Println("Error in string -> int")
			fmt.Println(err.Error())
			resp.sonar1=-1
		}
		resp.sonar2,err=strconv.Atoi(elem[5])
		if err!=nil{
			fmt.Println("Error in string -> int")
			fmt.Println(err.Error())
			resp.sonar2=-1
		}
		sensorQueue<-resp
	}

}