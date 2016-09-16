//Author: Peter Nagy - https://github.com/pete314
//Since: 2016.09.15.
//Description: Create permutation from string input

package main

import (
	"fmt"
)
const STR_LENGTH = 4
var permCounter = 0

func exchangeChars(charSequence []rune, fromInx int, toIndex int) []rune{
	tmpChar := charSequence[fromInx]
	charSequence[fromInx] = charSequence[toIndex]
	charSequence[toIndex] = tmpChar

	return charSequence
}

func selfPermutation(charSequence []rune, holdIndex int, iterations int) []rune{
	permutationStr := string(charSequence)
	for iCnt := 0; iCnt < iterations; iCnt++ {
		permCounter++
		fmt.Printf("\nPermutation %v: %v", permCounter, permutationStr)

		if iCnt % 2 == 1 && iCnt % 3 != 0{
			permutationStr = string(exchangeChars(charSequence, iCnt, iCnt * 2))
		}else if iCnt != STR_LENGTH - 1 {
			permutationStr = string(exchangeChars(charSequence, iCnt, iCnt + 1))
		}else{
			permutationStr = string(exchangeChars(charSequence, iCnt, 0))
		}
	}

	if holdIndex != STR_LENGTH - 1{
		return []rune(permutationStr)//exchangeChars(charSequence, holdIndex, holdIndex+1)
	}else{
		return exchangeChars(charSequence, holdIndex, 0)
	}
}

func main(){
	var inputStr string
	fmt.Printf("Please give %v char string:", STR_LENGTH)
	fmt.Scan(&inputStr)

	properLenStr := inputStr[0:STR_LENGTH]
	inputChars := []rune(properLenStr)

	for iCnt := 0; iCnt < 4; iCnt++{
		inputChars = selfPermutation(inputChars, iCnt, STR_LENGTH);
	}
}
