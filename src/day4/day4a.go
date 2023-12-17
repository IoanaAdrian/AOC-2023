package day4

import (
	"fmt"
	"github.com/IoanaAdrian/AOC-2023/reader"
	"github.com/IoanaAdrian/AOC-2023/utils"
	"strings"
)

func getWinningNumbers(line string)[]int{
	colonIndex := strings.Index(line, ":")
	separationBarIndex := strings.Index(line, "|")
	if colonIndex == -1 || separationBarIndex == -1{
		return []int{}
	}
	winningNumbersString := line[colonIndex+1 : separationBarIndex-1]
	
	return utils.GetNumbersFromString(winningNumbersString)
}

func getScratchedNumbers(line string)[]int{
	separationBarIndex := strings.Index(line, "|")
	if separationBarIndex == -1{
		return []int{}
	}
	scratchedNumbersString := line[separationBarIndex+2:]

	return utils.GetNumbersFromString(scratchedNumbersString)
}

func SolveA(){
	sum := 0
	lines := reader.GetLines("./day4/input/full_input.txt")
	for _, line := range lines {
		cardWorth := 0
		winningNumbers := getWinningNumbers(line)
		scratchedNumbers := getScratchedNumbers(line)
		for i := 0; i < len(winningNumbers); i++{
			if utils.NumberIndexInArray(winningNumbers[i], scratchedNumbers) != -1{
				if cardWorth == 0{
					cardWorth = 1
				}else{
					cardWorth *= 2
				}
			}
		}
		sum += cardWorth
	}
	fmt.Println(sum)
}