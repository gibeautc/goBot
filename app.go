package main

import (
	"fmt"
	"time"
)

type App struct{
	//main token to pass around
	curMotor *motorPair
	stale int		//number of ms we concider a messge to be stale and not process it.
	collision bool  //have we just hit something?  Processing and movements will be slower if this has happened
	curSensor *sensorData
	lastPrint time.Time
	/*
	brain will store whatever interface we are currently using.
	Most of the time this would be static, whatever type of navigation I am working on, but I also
	think it could be dynamic, such as using lineFollower when on a road to stay in the bike lane
	or bumpLogic in a dark environment

	may want to consider have a secondary brain, or subBrain, and compare decisions.
	 */
 	primaryBrain   		navLogic
	secondaryBrain		navLogic
	tertiaryBrain  		navLogic
	primaryDecision 	*motorPair
	secondaryDecision 	*motorPair
	tertiaryDecision 	*motorPair
	primaryAssert bool //used if the primary logic wants complete control over decision

}

func(self *App) setup(){
	self.stale=50 //ms
	self.curMotor=new(motorPair)
	self.curSensor=new(sensorData)
	self.collision=false
	self.lastPrint=time.Now()

	self.primaryBrain=new(bumpLogic)
	self.secondaryBrain=new(lineFollower)
	self.tertiaryBrain=nil
	self.primaryAssert=false

}

func(self *App) mix(){
	/*
	Not sure yet how we want to decide between each decision, maybe each of the logics should have a percentage on how
	confident it is?
	 */

}

func(self *App) printApp(){
	fmt.Println("************System Status************")
	fmt.Printf("Stale setting(ms): %d\n",self.stale)
	fmt.Printf("In Collision State: %t\n",self.collision)
	fmt.Println("Current Motor Settings:")
	self.curMotor.printMotor()
	fmt.Printf("Bump1: %t\n",self.curSensor.bump1)
	fmt.Printf("Bump2: %t\n",self.curSensor.bump2)
	fmt.Printf("Bump3: %t\n",self.curSensor.bump3)
	fmt.Printf("Bump4: %t\n",self.curSensor.bump4)
	fmt.Printf("Sonar1(cm): %d\n",self.curSensor.sonar1)
	fmt.Printf("Sonar2(cm): %d\n\n",self.curSensor.sonar2)
	fmt.Printf("Primary Nav Type: %s\n",self.primaryBrain.showType())
	if self.secondaryBrain!=nil{
		fmt.Printf("Secondary Nav Type: %s\n",self.secondaryBrain.showType())
	}else{
		fmt.Printf("Secondary Nav Type: None\n")
	}

	if self.tertiaryBrain!=nil{
		fmt.Printf("Tertiary Nav Type: %s\n",self.tertiaryBrain.showType())
	}else{
		fmt.Printf("Tertiary Nav Type: None\n")
	}
	fmt.Println("Primary Decision:")
	self.primaryDecision.printMotor()
	fmt.Println("Secondary Decision:")
	self.secondaryDecision.printMotor()
	fmt.Println("Tertiary Decision:")
	self.tertiaryDecision.printMotor()
}
