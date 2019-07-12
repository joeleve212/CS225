package main

import "fmt"
// import "math"
//import s "strings"

func main(){
	wordIn := ""
	repeat := true
	rep := ""
	//fmt.Println(s.Contains("tester", "e"), " poop", test)
	for repeat{
		fmt.Println("Input a word to Hawaiian-ify: ")
		fmt.Scan(&wordIn)
		//TODO: check for each letter combo, printing 
		for i := 0; i<len(wordIn);i++{
			//example usage: fmt.Print(string(wordIn[i])," ")
			if string(wordIn[i]) == "w"{
				//check w cases
				continue
			}	else if string(wordIn[i]) == "p" || string(wordIn[i])=="k"||string(wordIn[i])=="h"||
					string(wordIn[i]) == "l" || string(wordIn[i]) == "m"||string(wordIn[i])=="n"||
					string(wordIn[i]) == "'"||string(wordIn[i]) == " "{
				fmt.Print(string(wordIn[i]))
				continue
			}	else if string(wordIn[i]) == "a" {
				if i == len(wordIn)-1{
					fmt.Print("ah")
				} else if string(wordIn[i+1]) == "i"||string(wordIn[i+1]) == "e"{
					fmt.Print("eye")
					i++
				} else if string(wordIn[i+1]) == "o"||string(wordIn[i+1]) == "u"{
					fmt.Print("ow")
					i++
				} else{
					fmt.Print("ah")
				}
			}	else if string(wordIn[i]) == "e"{
				//check e cases
				if i == len(wordIn)-1{
					fmt.Print("eh")
				} else if string(wordIn[i+1]) == "i"{
					fmt.Print("ay")
					i++
				} else if string(wordIn[i+1]) == "u"{
					fmt.Print("eh-oo")
					i++
				} else{
					fmt.Print("eh")
				}
			}	else if string(wordIn[i]) == "i"{
				//check i case
				if i == len(wordIn)-1{
					fmt.Print("ee")
				} else if string(wordIn[i+1]) == "u"{
					fmt.Print("ew")
					i++
				} else{
					fmt.Print("ee")
				}
			}	else if string(wordIn[i]) == "o"{
				//check o cases
				if i == len(wordIn)-1{
					fmt.Print("oh")
				} else if string(wordIn[i+1]) == "i"{
					fmt.Print("oy")
					i++
				} else if string(wordIn[i+1]) == "u"{
					fmt.Print("ow")
					i++
				} else {
					fmt.Print("oh")
				}
			}	else if string(wordIn[i]) == "u"{
				//check u case
				if i == len(wordIn)-1{
					fmt.Print("oo")
				} else if string(wordIn[i+1]) == "i"{
					fmt.Print("ooey")
					i++
				} else{
					fmt.Print("oo")
				}
			}	else {
				fmt.Println("\n\nInvalid character present!")
				break
			}
			if i < len(wordIn)-1 && string(wordIn[i+1])!=" " {
				fmt.Print("-")
			}
		}
		//TODO: check for each illegit letter
		fmt.Println("\nWould you like to try another word? (Y/N)")
		fmt.Scan(&rep)
		if rep == "N"||rep=="n"||rep=="no" {repeat = false;}
	}
	
}