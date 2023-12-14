package day2

import(
	"github.com/IoanaAdrian/AOC-2023/reader"
	"fmt"
	"strings"
	"strconv"
	"errors"
)

type CubeColorAmount struct{
	red int 
	green int
	blue int 
}

func getGameId(line string) (int, error){
	indexFirstSpace := strings.Index(line, " ")
	indexColon := strings.Index(line, ":")
	gameId, err := strconv.ParseUint(string(line[indexFirstSpace+1:indexColon]), 10, 64)
	if err != nil{
		return 0, errors.New("Unable to process gameId")
	}
	return int(gameId), nil
}

func getGameSets(line string) []string{
	// Line form: Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	result := []string{}
	indexColon := strings.Index(line, ":")
	sets := line[indexColon+2:len(line)]
	indexSemiColon := strings.Index(sets,";")
	for indexSemiColon != -1{
		result = append(result, sets[:indexSemiColon])
		sets = sets[indexSemiColon+2:]
		indexSemiColon = strings.Index(sets,";")
	}
	result = append(result, sets)
	return result
}

func countCubeColors(colorLine string, c *CubeColorAmount) error{
	spaceIndex := strings.Index(colorLine, " ")
	cubeAmount, err := strconv.ParseUint(string(colorLine[:spaceIndex]), 10, 64)
	if err != nil{
		fmt.Println(err)
		return errors.New("Failed to parse cube amount")
	}
	if colorLine[spaceIndex+1:] == "red"{
		c.red += int(cubeAmount)
	}else if colorLine[spaceIndex+1:] == "green"{
		c.green += int(cubeAmount)
	}else if colorLine[spaceIndex+1:] == "blue"{
		c.blue += int(cubeAmount)
	}
	return nil
}

func getCubeColorAmounts(sets []string)[]CubeColorAmount{
	// Set form: 8 green, 6 blue, 20 red
	result := []CubeColorAmount{}
	for i := 0 ; i < len(sets); i++{
		result = append(result, CubeColorAmount{red: 0, blue: 0, green: 0})
		indexComma := strings.Index(sets[i],",")
		for indexComma != -1{
			colorLine := sets[i][:indexComma] // colorLine form: 1 blue
			err := countCubeColors(colorLine, &result[i])
			if err != nil{
				return []CubeColorAmount{}
			}

			sets[i] = sets[i][indexComma+2:]
			indexComma = strings.Index(sets[i],",")
		}
		colorLine := sets[i]
		err := countCubeColors(colorLine, &result[i])
		if err != nil{
			return []CubeColorAmount{}
		}
	}
	return result
}

func SolveA(){
	lines := reader.GetLines("./day2/input/full_input.txt")
	sumGameIds := 0
	for _, line := range lines {
		allSetsOk := true
		colorAmounts := getCubeColorAmounts(getGameSets(line))
		for i := 0; i < len(colorAmounts); i++{
			if colorAmounts[i].red > 12 || colorAmounts[i].green > 13 || colorAmounts[i].blue > 14{
				allSetsOk = false
			}
		}
		if allSetsOk{
			gameId, err := getGameId(line)
			if err != nil{
				return
			}
			sumGameIds += gameId
		}
	}
	fmt.Println(sumGameIds)
}

