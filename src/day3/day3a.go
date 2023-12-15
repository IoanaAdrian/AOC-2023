package day3

import (
	"fmt"
	"github.com/IoanaAdrian/AOC-2023/reader"
	"github.com/IoanaAdrian/AOC-2023/utils"
)

type LineNumber struct{
	number int
	numberIndexes []int
}

type Coordinate struct{
	x int
	y int
}

var directionArray = []Coordinate{
	{x: 0, y: 1},
	{x: 0, y: -1},
	{x: 1, y: 0},
	{x: -1, y: 0},
	
	{x: -1, y: -1},
	{x: -1, y: 1},
	{x: 1, y: 1},
	{x: 1, y: -1},
}

func isInMat(coord Coordinate, mat []string) bool{
	return coord.x >= 0 && coord.x < len(mat[0]) && coord.y >= 0 && coord.y < len(mat)
}

func getLineNumbers(line string) []LineNumber{
	num := 0
	var indexes []int
	var result []LineNumber
	for i := 0; i < len(line); i++{
		n, err := utils.IsDigit(line[i])
		if err == nil{
			num = num*10 + int(n)
			indexes = append(indexes, i)
		}else if num != 0{
			result = append(result, LineNumber{number: num, numberIndexes: indexes})
			indexes = []int{}
			num = 0
		}
	}
	if num != 0{
		result = append(result, LineNumber{number: num, numberIndexes: indexes})
	}
	return result
}

func SolveA(){
	var mat []string
	sum := 0
	lines := reader.GetLines("./day3/input/full_input.txt")
	for _, line := range lines {
		mat = append(mat, line)
	}
	for lineIndex, line := range lines{
		lineNumbers := getLineNumbers(line)
		for i := 0; i < len(lineNumbers); i++{
			hasAdjacentSymbol := false
			for j := 0; j < len(lineNumbers[i].numberIndexes); j++{
				for k := 0; k < len(directionArray); k++{
					x := lineIndex
					y := lineNumbers[i].numberIndexes[j]
					x += directionArray[k].x
					y += directionArray[k].y
					if isInMat(Coordinate{x, y}, mat){
						_, err := utils.IsDigit(mat[x][y])
						if err != nil && string(mat[x][y]) != "."{
							hasAdjacentSymbol = true
						}
					}
				}
			}
			if hasAdjacentSymbol{
				sum += lineNumbers[i].number
			}
		}
	}
	fmt.Println(sum)
}