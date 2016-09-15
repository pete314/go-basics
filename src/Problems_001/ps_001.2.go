//Author: Peter Nagy - https://github.com/pete314
//Since: 2016.09.15.
//Description: Find a the n-th prime number

package main

import "fmt"

//The n-th prime to find
const PRIME_IN_ROW = 10001

//Check weather a number is prime; return bool
func checkIsPrime(num int) (bool){
	if num % 2 == 0 {
		return false;
	}

	//Check if num can be divided by odd numbers
	for iCnt := 3; iCnt * iCnt <= num; iCnt += 2{
		if num % iCnt == 0 {
			return false
		}
	}

	return true
}

func main(){
	primeCnt, number := 0, 0

	for {
		number++
		if checkIsPrime(number){
			primeCnt++
		}
		if primeCnt == PRIME_IN_ROW{
			break
		}
	}

	fmt.Printf("The %v prime number is %v", PRIME_IN_ROW, number)
}
