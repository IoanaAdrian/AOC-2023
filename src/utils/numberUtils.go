package utils

import (
	"errors"
	"strconv"
)

func IsDigit(b byte) (int8, error) {
    r, err := strconv.ParseUint(string(b), 10, 4) // one digit can be 0..9, so max 4 bits (9 = 1001)
    if err != nil{
		return 0, errors.New("character is not a number")
	}
	return int8(r), nil
}

func GetFirstDigit(line string) (int8, error){
	for i := 0; i < len(line); i++{
		n, err := IsDigit(line[i])
		if err == nil{
			return n, nil
		}
	}
	return 0, errors.New("input string does not contain any number")
}

func GetLastDigit(line string) (int8, error){
	for i := len(line)-1; i >= 0; i--{
		n, err := IsDigit(line[i])
		if err == nil{
			return n, nil
		}
	}
	return 0, errors.New("input string does not contain any number")
}

func GetNumbersFromString(input string)[]int{ // "12 13 14 5" => [12,13,14,5]
	input += " " // to avoid duplicating the logic
	var result []int
	eachNumber := 0
	for i := 0; i < len(input); i++{
		d, err := IsDigit(input[i])
		if err == nil{
			eachNumber = eachNumber * 10 + int(d)
		}else if eachNumber != 0{
			result = append(result, eachNumber)
			eachNumber = 0
		}
	}
	return result
}

func NumberIndexInArray(n int, arr []int)int{
	for i := 0; i < len(arr); i++{
		if n == arr[i]{
			return i
		}
	}
	return -1
}