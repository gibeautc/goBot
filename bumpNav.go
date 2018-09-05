package main


type bumpLogic struct{

}

//required part of interface
func(self *bumpLogic) decide(app *App) (*motorPair,error){
	return stop(),nil
}

//required part of interface
func(self *bumpLogic) showType() string{
	return "BumpLogic"
}



