package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var multiplyMap = map[string][]int{}

func main() {
	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	q, err := os.ReadFile(getwd + "/day3/question")
	if err != nil {
		log.Println(err)
		return
	}
	row := strings.Split(string(q), "\n")
	log.Println(solveP2(row))
}

func solveP2(row []string) int {
	sum := 0
	for i, col := range row {
		upperCol := ""
		lowerCol := ""
		if i-1 > 0 {
			upperCol = row[i-1]
		}
		if i+1 < len(row) {
			lowerCol = row[i+1]
		}
		re := regexp.MustCompile(`\d+`)
		strIndices := re.FindAllStringIndex(col, -1)
		for _, index := range strIndices {
			if idx := checkMultiply(index, col); idx != -1 {
				mapKey := fmt.Sprintf("%v:%v", i, idx)
				multiplyMap[mapKey] = append(
					multiplyMap[mapKey],
					StrToInt(col[index[0]:index[1]]),
				)
				continue
			}
			if idx := checkMultiply(index, lowerCol); idx != -1 {
				mapKey := fmt.Sprintf("%v:%v", i+1, idx)
				multiplyMap[mapKey] = append(
					multiplyMap[mapKey],
					StrToInt(col[index[0]:index[1]]),
				)
				continue
			}
			if idx := checkMultiply(index, upperCol); idx != -1 {
				mapKey := fmt.Sprintf("%v:%v", i-1, idx)
				multiplyMap[mapKey] = append(
					multiplyMap[mapKey],
					StrToInt(col[index[0]:index[1]]),
				)
				continue
			}
		}
	}
	log.Println(multiplyMap)
	for _, ints := range multiplyMap {
		multResult := 1
		if len(ints) != 1 {
			for _, v := range ints {
				multResult *= v
			}
		}
		if multResult != 1 {
			sum += multResult
		}
	}
	return sum
}

func solveP1(row []string) int {
	sum := 0
	for i, col := range row {
		upperCol := ""
		lowerCol := ""
		if i-1 > 0 {
			upperCol = row[i-1]
		}
		if i+1 < len(row) {
			lowerCol = row[i+1]
		}
		re := regexp.MustCompile(`\d+`)
		strIndices := re.FindAllStringIndex(col, -1)
		for _, index := range strIndices {
			if checkSymbol(index, col) {
				sum += StrToInt(col[index[0]:index[1]])
				log.Println(col[index[0]:index[1]])
				continue
			}
			if checkSymbol(index, lowerCol) {
				sum += StrToInt(col[index[0]:index[1]])
				log.Println(col[index[0]:index[1]])
				continue
			}
			if checkSymbol(index, upperCol) {
				sum += StrToInt(col[index[0]:index[1]])
				log.Println(col[index[0]:index[1]])
				continue
			}
		}
	}
	return sum
}

func checkSymbol(index []int, col string) bool {
	if col == "" {
		return false
	}
	for j := index[0]; j < index[1]; j++ {
		char := string(col[j])
		if isValidSymbol(char) {
			return true
		}
		if j == index[0] && j != 0 {
			if isValidSymbol(string(col[j-1])) {
				return true
			}
		}
		if j == index[1]-1 && j != len(col)-1 {
			if isValidSymbol(string(col[j+1])) {
				return true
			}
		}
	}
	return false
}

func checkMultiply(index []int, col string) int {
	if col == "" {
		return -1
	}
	for j := index[0]; j < index[1]; j++ {
		char := string(col[j])
		if isMultiply(char) {
			return j
		}
		if j == index[0] && j != 0 {
			if isMultiply(string(col[j-1])) {
				return j - 1
			}
		}
		if j == index[1]-1 && j != len(col)-1 {
			if isMultiply(string(col[j+1])) {
				return j + 1
			}
		}
	}
	return -1
}

func isValidSymbol(char string) bool {
	if char == "$" ||
		char == "#" ||
		char == "+" ||
		char == "-" ||
		char == "/" ||
		char == "%" ||
		char == "(" ||
		char == ")" ||
		char == "=" ||
		char == "@" ||
		char == "&" ||
		char == "*" {
		return true
	}
	return false
}
func isMultiply(char string) bool {
	return char == "*"
}

func StrToInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}
