package main

import "fmt"

type Liters float64
type Gallons float64

func main() {

	var carFuel Gallons = Gallons(10.5)
	busFuel := Liters(240.4) /* short variable declaration */

	fmt.Println(carFuel, busFuel) //10.5 240.4

	//=== conversion ===============
	//carFuel = Liters(1.5) /* directly not convert to underlying value*/
	//carFuel = Gallons(Liters(1.5)) /* invalid calculation */
	carFuel = Gallons(Liters(1.5) * 0.264)
	fmt.Println(carFuel) //0.396
}
