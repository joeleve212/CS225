package main

import "fmt"
import "math/rand"

func sliceMaker(length, randMult int) {
	slice1 := make([]float32, length)
	var sum float32
	for i:=0;i<length;i++ {
		slice1[i] = float32(rand.Intn(100)*randMult)
		sum += slice1[i]
		fmt.Println("Slice[",i,"] = ",slice1[i])
	}
	fmt.Println("Full Slice: ",slice1)
	fmt.Println("Sum: ", sum)
}

func main() {
	sliceMaker(3, 2)
	sliceMaker(5, 16)
	sliceMaker(8, 4)
}
