package main

import (
	"fmt"
	"testing"
)

func TestGetCords(t *testing.T) {
	lat1 := 44.588935
	lon1 := -123.239885

	lat2 := 44.588996
	lon2 := -123.237367

	dist, bearing := distanceBetween(lat1, lon1, lat2, lon2)
	fmt.Println("Distance And Brearing")
	fmt.Println(dist)
	fmt.Println(bearing)
	lat3, lon3 := getCords(lat1, lon2, dist, bearing)
	fmt.Println("Cords Again")
	fmt.Println(lat3)
	fmt.Println(lon3)

}
