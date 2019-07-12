//asdf
package main

import "fmt"
import "math/rand"
import "math"

func main(){
	//Read in which digit (MSD is 1st, then 2nd down to LSD) to look at
	dig, length, j:= 0,0, 0
	fmt.Println("Which digit should be counted?(1-4)")
	fmt.Scan(&dig)
	//Produce random data set to evaluate
	fmt.Println("How many numbers should be used?")
	fmt.Scan(&length)
	nums := make([]int, length)
	var digs [10] int 
	digs[0],digs[1],digs[2],digs[3],digs[4],digs[5],digs[6],digs[7],digs[8],digs[9]= 0,0,0,0,0,0,0,0,0,0
	for i:=0;i<length;i++{
		nums[i] = rand.Intn(1000)
//Testing:		fmt.Println("New Num: ",nums[i])
		temp:=nums[i]
		if float64(nums[i]) > math.Pow10(dig-1) {
			for j=1;temp>int(math.Pow10(dig));j++ {
				temp = temp/10
			}
			digs[temp % 10]++
		}
	}
	for i:=0;i<10;i++{
		fmt.Println("Total ",i,"'s: ", digs[i])
	}
}