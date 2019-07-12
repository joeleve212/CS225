package main

import "fmt"
import "math"

func calcCircle(rad float64) (float64, float64){
	circ := math.Pi * 2 * rad
	area := math.Pi * math.Pow(rad,2)
	return circ, area
}

func main(){
	radius := 7.5
	c, a := calcCircle(radius)
	fmt.Println("Radius: ", radius, " Circumference: ", c, " Area: ", a)
}