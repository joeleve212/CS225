//Joe Leveille PS7 project 1 - Pig Dice Game
package main
import "fmt"
import "math/rand"
func main(){
	win,currPlay, sum, currRoll:=0,0,0,0
	var score [2] int 
	score[0],score[1]= 0, 0
	choice := ""
	for win == 0 {
		fmt.Println("Player ",currPlay+1,", it is your turn")
		if sum > 1 {
			fmt.Println("Would you like to roll again? (Y/N)")
			fmt.Scan(&choice)
			if choice == "N" || choice == "n" {
				score[currPlay] += sum
				sum = 0
				if score[currPlay] >= 100 {
					win = currPlay +1
					continue
				}
				currPlay = (currPlay + 1) % 2
				continue
			}
		}
		currRoll = rand.Intn(6) + 1
		fmt.Println("You roll: ",currRoll)
		if currRoll == 1 {
			fmt.Println("That ends your turn!")
			sum = 0
			currPlay = (currPlay + 1) % 2
			continue
		}	else{
			sum += currRoll
			fmt.Println("Your current sum for this turn: ", sum)
		}
	}
	fmt.Println("Player ", win, " won the game!")
}