//Author: Peter Nagy - https://github.com/pete314
//Since: 2016.09.15.
//Description: Find repeating sequence based by applying the following to user input
//		if number is even divide it by 2 if odd multiply by 3 and add 1

package main

import (
	"fmt"
	"strconv"
)

const (
	LOWER_BOUND = 0
	UPPER_BOUND = 1000000 //defines the maximum sequences to run
)

func getStartingNumber() int64{
	var strNum string
	for{
		fmt.Print("Enter a positive integer to start with: ")
		fmt.Scan(&strNum)
		num, err := strconv.ParseInt(strNum, 10, 64)

		if err != nil {
			fmt.Println("Could not parse given number, try again!")
		}else if num <= 0{
			fmt.Println("Number can not be 0 or negative, try again!")
		}else{
			return num
		}
	}
}

func main(){
	currentSequence := make(map[int64]int64)
	seqCnt, number := 0, getStartingNumber()
	currentSequence[number] = number

	fmt.Print("The current sequence: ")
	for{
		fmt.Printf("%v ", number)
		if number % 2 == 0 {
			number /= 2
		}else{
			number = (number * 3) + 1
		}

		_ , hasValue := currentSequence[number]

		if hasValue {
			fmt.Println("\n\nSequence would now repeate, current number: ", number)
			break
		}else{
			currentSequence[number] = number
		}

		seqCnt++
		if seqCnt == UPPER_BOUND || number == LOWER_BOUND {
			fmt.Println("\n\nLower or upper bound, current number: ", number)
			break
		}
	}
}