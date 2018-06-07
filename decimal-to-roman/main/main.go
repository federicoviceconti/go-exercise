package main

import (
	"fmt"
	"strconv"
	"math"
)

type RomanLiteral struct {
	letter string
	value  uint16
}

var FOUR = "4"
var NINE = "9"

func main() {
	var romanDecimalStart uint16 = 3999
	transformRomanToDecimal(romanDecimalStart)
}

func transformRomanToDecimal(romanToTransform uint16) {
	var romanLiterals = createRomanArray()
	var currentValue = romanToTransform

	fmt.Println(getResult(currentValue, romanLiterals))
}
func getResult(counterRomanValue uint16, romanLiterals []RomanLiteral) (string) {
	var result = ""
	for index, currentLiteral := range romanLiterals {
		for isDecimalMajorThanRomanLiteral(counterRomanValue, currentLiteral)  {
			if isASpecialCase(counterRomanValue) {
				result += getSpecialCaseRoman(getFirstCharacterOfANumber(int(counterRomanValue)), romanLiterals, index)
				counterRomanValue -= subtractNumberForFirstCharacter(int(counterRomanValue))
			} else {
				result += currentLiteral.letter
				counterRomanValue -= currentLiteral.value
			}
		}
	}

	return result
}
func getSpecialCaseRoman(firstCharacter string, romanLiterals []RomanLiteral, index int) string {
	if firstCharacter == FOUR {
		return romanLiterals[index].letter + romanLiterals[index - 1].letter
	} else {
		return romanLiterals[index + 1].letter + romanLiterals[index - 1].letter
	}
}

func subtractNumberForFirstCharacter(num int) (uint16){
	var firstCharacter = getFirstCharacterOfANumber(num)
	var firstCharacterNum, err = strconv.Atoi(firstCharacter)
	if err != nil {
		return 0
	}

	var decimalPositions = float64(len(strconv.Itoa(num)) - 1)
	return uint16(firstCharacterNum * int(math.Pow(10, decimalPositions)))
}

func getFirstCharacterOfANumber(num int) (string) {
	var stringNumber = strconv.Itoa(num)
	return string( stringNumber[0])
}

func isASpecialCase(num uint16) (bool) {
	var firstCharacter = getFirstCharacterOfANumber(int(num))
	return firstCharacter == FOUR || firstCharacter == NINE
}

func isDecimalMajorThanRomanLiteral(romanValue uint16, currentLiteral RomanLiteral) (bool) {
	return romanValue >= currentLiteral.value
}

func createRomanArray() ([]RomanLiteral) {
	var romans = []RomanLiteral {
		{"M", 1000},
		{"D", 500},
		{"C", 100},
		{"L", 50},
		{"X", 10},
		{"V", 5},
		{"I", 1},
	}

	return romans
}
