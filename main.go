package main

import (
	"fmt"
)

func main() {
	app := new(App)
	app.setup()
	fmt.Println("Starting up")
	app.run()
}
