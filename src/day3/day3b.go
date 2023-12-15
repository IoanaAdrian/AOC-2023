package day3

import (
	"fmt"
	"github.com/IoanaAdrian/AOC-2023/reader"
	"github.com/IoanaAdrian/AOC-2023/utils"
)

type numberAdjacentSymbol struct{
	number int
	symbolsCoordinate []Coordinate
}

func index(coordinate Coordinate, symbolsCoordinate []Coordinate)int{
	for i, x := range symbolsCoordinate{
		if x == coordinate{
			return i
		}
	}
	return -1
}

// TODO: Use this to maybe improve the sum algorithm
func removeElement(ind int, numbersAdjacentSymbol []numberAdjacentSymbol)[]numberAdjacentSymbol{
	return append(numbersAdjacentSymbol[:ind], numbersAdjacentSymbol[ind+1:]...)
}

func SolveB(){
	var mat []string
	var numbersAdjacentSymbol []numberAdjacentSymbol
	lines := reader.GetLines("./day3/input/full_input.txt")
	for _, line := range lines {
		mat = append(mat, line)
	}
	for lineIndex, line := range lines{
		lineNumbers := getLineNumbers(line)
		for i := 0; i < len(lineNumbers); i++{
			symbolsCoordinate := []Coordinate{}
			for j := 0; j < len(lineNumbers[i].numberIndexes); j++{
				for k := 0; k < len(directionArray); k++{
					x := lineIndex
					y := lineNumbers[i].numberIndexes[j]
					x += directionArray[k].x
					y += directionArray[k].y
					if isInMat(Coordinate{x, y}, mat){
						_, err := utils.IsDigit(mat[x][y])
						if err != nil && string(mat[x][y]) != "."{
							if index(Coordinate{x,y}, symbolsCoordinate) == -1{
								symbolsCoordinate = append(symbolsCoordinate, Coordinate{x,y})
							}
						}
					}
				}
			}
			if len(symbolsCoordinate) != 0{
				numbersAdjacentSymbol = append(numbersAdjacentSymbol, numberAdjacentSymbol{number: lineNumbers[i].number, symbolsCoordinate: symbolsCoordinate})
			}
		}
	}
	sum := 0
	for i := 0; i < len(numbersAdjacentSymbol); i++{
		for j := i+1; j < len(numbersAdjacentSymbol); j++{
			for k := 0; k < len(numbersAdjacentSymbol[i].symbolsCoordinate); k++{
				if index(numbersAdjacentSymbol[i].symbolsCoordinate[k], numbersAdjacentSymbol[j].symbolsCoordinate) != -1{
					sum = sum + numbersAdjacentSymbol[i].number * numbersAdjacentSymbol[j].number
				}
			}
		}
	}
	fmt.Println(sum)
}