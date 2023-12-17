package day4

import (
	"fmt"
	"github.com/IoanaAdrian/AOC-2023/reader"
	"github.com/IoanaAdrian/AOC-2023/utils"
	"strings"
)

type cardElement struct{
	id int
	value int
}

func getCardId(line string) int{
	colonIndex := strings.Index(line, ":")
	numbers := utils.GetNumbersFromString(line[:colonIndex])
	if len(numbers) == 0{
		return 0
	}
	return numbers[0]
}

func SolveB(){
	cardValues := make(map[int]int)
	var queue []cardElement
	lines := reader.GetLines("./day4/input/full_input.txt")
	for _, line := range lines {
		cardId := getCardId(line)
		nOfMatches := 0
		winningNumbers := getWinningNumbers(line)
		scratchedNumbers := getScratchedNumbers(line)
		for i := 0; i < len(winningNumbers); i++{
			if utils.NumberIndexInArray(winningNumbers[i], scratchedNumbers) != -1{
				nOfMatches++
			}
		}
		cardValues[cardId] = nOfMatches
		queue = append(queue, cardElement{id: cardId, value: nOfMatches})
	}
	count := 0
	for len(queue) != 0{
		count++
		for i := 1; i <= queue[0].value; i++{
			queue = append(queue, cardElement{queue[0].id+i, cardValues[queue[0].id+i]})
		}
		queue = queue[1:]
	}
	
	fmt.Println(count)
}