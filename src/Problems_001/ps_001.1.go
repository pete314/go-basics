//Author: Peter Nagy
//Since: 15/09/2016
//Description: Calculate the sum of number divisible by 3 or 5 between 0-1000

package main

import "fmt"

func main(){
	var totalSum int

	for iCnt := 0; iCnt < 1000; iCnt++ {
		if iCnt % 3 == 0 || iCnt % 5 == 0 {
			totalSum += iCnt
		}
	}

	fmt.Println("The sum of numbers divisible by 3 or 5 between 0 and 1000 is: ", totalSum)
}