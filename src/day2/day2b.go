package day2

import(
	"github.com/IoanaAdrian/AOC-2023/reader"
	"fmt"
)

func SolveB(){
	lines := reader.GetLines("./day2/input/full_input.txt")
	powerSum := 0
	for _, line := range lines {
		colorAmounts := getCubeColorAmounts(getGameSets(line))
		minRedAmount := -1 // Min amount for the game to be possible, so max amount from all values
		minBlueAmount := -1
		minGreenAmount := -1
		for i := 0; i < len(colorAmounts); i++{
			if minRedAmount == -1 || minRedAmount < colorAmounts[i].red{
				minRedAmount = colorAmounts[i].red
			}

			if minBlueAmount == -1 || minBlueAmount < colorAmounts[i].blue{
				minBlueAmount = colorAmounts[i].blue
			}
			if minGreenAmount == -1 || minGreenAmount < colorAmounts[i].green{
				minGreenAmount = colorAmounts[i].green
			}
		}
		powerSum += (minBlueAmount*minGreenAmount*minRedAmount)
	}
	fmt.Println(powerSum)
}
