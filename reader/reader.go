package reader
import (
	"bufio"
	"fmt"
	"os"
)

// This function uses .append()
// Check its performance over allocating in chunks with make[] 
func GetLines(filename string) []string{
	readFile, err := os.Open(filename)
	arr := []string{}
 
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
 
	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}
 
	readFile.Close()

	return arr
}