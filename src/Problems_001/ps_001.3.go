//Author: Peter Nagy - https://github.com/pete314
//Since: 2016.09.15.
//Description: Reverse user input string

package main

import "fmt"

func reverseString(input string) string{
	inputChars, reversedChars := []rune(input) , []rune{}

	for iCnt := len(inputChars) - 1 ; iCnt >= 0; iCnt--{
		reversedChars = append(reversedChars, inputChars[iCnt])
	}

	return string(reversedChars)
}

func main(){
	var inputString string
	fmt.Print("Enter the string: ")
	fmt.Scanln(&inputString)

	reversedInput := reverseString(inputString)
	fmt.Println("The reversed input: ", reversedInput)
}
