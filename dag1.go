package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	var coodinates []string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter coordinates")

	for {
		scanner.Scan()
		coor := scanner.Text()
		if coor == "" {
			break
		}
		coodinates = append(coodinates, coor)
	}

	len := len(coodinates)
	acc := 0

	for i := 0; i < len; i++ {
		var str string = coodinates[i]
		fNum := ""
		lNum := ""
		for _, char := range str {
			if (unicode.IsDigit(char)) && fNum == "" && lNum == "" {
				fNum = string(char)
				lNum = string(char)
			} else if (unicode.IsDigit(char)) && fNum != "" {
				lNum = string(char)
			}
		}
		di, err := strconv.Atoi(fNum + lNum)
		if err == nil {
			acc += di
		}
	}
	fmt.Println("Total:", acc)
}
