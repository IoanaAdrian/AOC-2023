package day1

import(
	"fmt"
	"github.com/IoanaAdrian/AOC-2023/reader"
	"github.com/IoanaAdrian/AOC-2023/utils"
)

func SolveA(){
	lines := reader.GetLines("./day1/input/full_input.txt")
	result := 0
	for _, line := range lines {
		firstDigit, err := utils.GetFirstDigit(line)
		if err == nil{
			lastDigit, err := utils.GetLastDigit(line)
			if err == nil{
				result = result + (int(firstDigit)*10+int(lastDigit))
			}
		}
	}
	fmt.Println(result)
}