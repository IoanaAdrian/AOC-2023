package utils

import(
	"strings"
)

var digitsToHRMap map[int8]string = map[int8]string{
	1 : "one",
	2 : "two",
	3 : "three",
	4 : "four",
	5 : "five",
	6 : "six",
	7 : "seven",
	8 : "eight",
	9 : "nine",
}

var HRToDigitsMap map[string]int8 = invertMap(digitsToHRMap)

func reverseString(input string) string{
	output := []byte{}
	for i := len(input)-1; i >= 0; i--{
		output = append(output, input[i])
	}
	return string(output)
}

func invertMap(m map[int8]string) map[string]int8 {
    inv := make(map[string]int8)
    for k, v := range m {
        inv[v] = k
    }
    return inv

}

func TranformDigitsToHR(line string) string{
	arr := []byte{}
	for i := 0; i < len(line); i++{
		n, err := IsDigit(line[i])
		if err == nil{
			digitHR := digitsToHRMap[n]
			for j := 0; j < len(digitHR); j++{
				arr = append(arr, digitHR[j])
			}
		}else{
			arr = append(arr, line[i])
		}
	}
	return string(arr)
}

func GetFirstHRDigit(line string) int8{
	digit := ""
	ind := len(line)
	for _,v := range digitsToHRMap{
		i := strings.Index(line, v)
		if i >= 0 && i < ind{
			digit = v
			ind = i
		}
	}
	return HRToDigitsMap[digit]
}

func GetLastHRDigit(line string) int8{
	
	reversedString := reverseString(line) 
	
	digit := ""
	ind := len(reversedString)
	for _,v := range digitsToHRMap{
		i := strings.Index(reversedString, reverseString(v))
		if i >= 0 && i < ind{
			digit = v
			ind = i
		}
	}
	return HRToDigitsMap[digit]
}