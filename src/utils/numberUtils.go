package utils

import(
	"strconv"
	"errors"
)

func IsDigit(b byte) (int8, error) {
    r, err := strconv.ParseUint(string(b), 10, 4) // one digit can be 0..9, so max 4 bits (9 = 1001)
    if err != nil{
		return 0, errors.New("Character is not a number")
	}
	return int8(r), nil
}

func GetFirstDigit(line string) (int8, error){
	for i := 0; i < len(line); i++{
		n, err := IsDigit(line[i])
		if (err == nil){
			return n, nil
		}
	}
	return 0, errors.New("Input string does not contain any number")
}

func GetLastDigit(line string) (int8, error){
	for i := len(line)-1; i >= 0; i--{
		n, err := IsDigit(line[i])
		if (err == nil){
			return n, nil
		}
	}
	return 0, errors.New("Input string does not contain any number")
}
