package day2

import(
	"fmt"
	"github.com/IoanaAdrian/AOC-2023/reader"
	"github.com/IoanaAdrian/AOC-2023/utils"
)


func Solve(){
	lines := reader.GetLines("./day1/full_input.txt")
	s := 0
	for _, line := range lines {
		HRline := utils.TranformDigitsToHR(line)
		s = s + int(utils.GetFirstHRDigit(HRline)) * 10 + int(utils.GetLastHRDigit(HRline))
	}
	fmt.Println(s)
}