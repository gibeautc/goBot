package main

const(
	followTypeCenter=0x01	//we want to be right over the line
	followTypeLeft=0x02		//we want to be left of the line
	followTypeRight=0x03	//we want to be right of the line
	/*
	followTypeRight would be the default for staying in a bike lane on the correct side of the road
	  */
	  followTypeBetween=0x04	//we want to be inbetween two lines

)



type lineFollower struct{

}

//required part of interface
func(self *lineFollower) decide(app *App) (*motorPair,error){

	return stop(),nil
}

//required part of interface
func(self *lineFollower) showType() string{
	return "LineFollower"
}