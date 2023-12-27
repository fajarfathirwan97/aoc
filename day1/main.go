package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var digitMap = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	q, err := os.ReadFile(getwd + "/day1/question")
	if err != nil {
		log.Println(err)
		return
	}
	row := strings.Split(string(q), "\n")
	//log.Println(solveP1(row))
	log.Println(solveP2(row))
}

func solveP2(row []string) int {
	sum := 0
	for _, col := range row {
		re := regexp.MustCompile(`([0-9]|one|two|three|four|five|six|seven|eight|nine)`)
		startIdx := 9999999999
		endIdx := -1
		var start, end int
		for key, _ := range digitMap {
			pos := strings.Index(col, key)
			lastPos := strings.LastIndex(col, key)
			//log.Println(pos, key)
			if pos != -1 && pos <= startIdx {
				startIdx = pos
				start = digitMap[col[pos:pos+len(key)]]
			}
			if lastPos != -1 && lastPos >= endIdx {
				endIdx = lastPos
				end = digitMap[col[lastPos:lastPos+len(key)]]
			}
		}
		allString := re.FindAllString(col, -1)
		val, _ := strconv.Atoi(fmt.Sprintf("%v%v", start, end))
		val2, _ := strconv.Atoi(fmt.Sprintf("%v%v", digitMap[allString[0]], digitMap[allString[len(allString)-1]]))
		if val != val2 {
			log.Println(val, val2, allString, col)
		}
		sum += val
	}
	return sum
}
func solveP1(row []string) int {
	sum := 0
	re := regexp.MustCompile(`([0-9])`)
	for _, col := range row {
		start := -1
		end := -1
		strArr := strings.Split(col, "")
		for i, _ := range strArr {
			j := len(strArr) - 1 - i
			if re.MatchString(strArr[i]) && start == -1 {
				if v, ok := digitMap[fmt.Sprintf("%v", strArr[i])]; ok {
					start = v
				}
			}
			if re.MatchString(strArr[j]) && end == -1 {
				if v, ok := digitMap[fmt.Sprintf("%v", strArr[j])]; ok {
					end = v
				}
			}
			if start != -1 && end != -1 {
				val, _ := strconv.Atoi(fmt.Sprintf("%v%v", start, end))
				sum += val
				break
			}
		}
	}
	return sum
}
