package main

import (
	"log"
	"os"
	"regexp"
	"strings"
)

// p1:30 min
var processedMap = map[int]int{}
var processedMapInt = map[int][]int{}
var numberOfStamp = 0

func main() {
	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	q, err := os.ReadFile(getwd + "/day4/question")
	if err != nil {
		log.Println(err)
		return
	}
	//sum := P1(q)
	P2(q)
}

func P2(q []byte) int {
	row := strings.Split(string(q), "\n")
	sum := 0
	for i := 0; i < len(row); i++ {
		//for i := 0; i < 1; i++ {
		processedMapInt[i] = append(processedMapInt[i], i)
		numberOfStamp++
		processP2(row, i)
	}
	log.Println(numberOfStamp)
	return sum
}

func processP2(row []string, i int) {
	matchedNumber, ok := processedMap[i]
	if !ok {
		_, matchedNumber = cardNoANdMatchedNumber(row, i)
	}
	processedMap[i] = matchedNumber
	for j := 0; j < matchedNumber; j++ {
		numberOfStamp++
		processP2(row, j+i+1)
	}
}

func cardNoANdMatchedNumber(row []string, i int) (string, int) {
	segment := strings.Split(row[i], " | ")
	leftSegment := strings.Split(segment[0], ": ")
	reCard := regexp.MustCompile(`\d+`)
	cardNo := reCard.FindAllString(leftSegment[0], -1)
	numbers := strings.Split(leftSegment[1], " ")
	winners := strings.Split(segment[1], " ")
	matchedNumber := calcMatchedNumber(numbers, winners)
	return cardNo[0], matchedNumber
}

func calcMatchedNumber(numbers []string, winners []string) int {
	matchedNumber := 0
	for _, n := range numbers {
		for _, winner := range winners {
			if n == winner && n != "" {
				matchedNumber += 1
				break
			}
		}
	}
	return matchedNumber
}
func P1(q []byte) int {
	row := strings.Split(string(q), "\n")
	sum := 0
	for _, col := range row {
		segment := strings.Split(col, " | ")
		leftSegment := strings.Split(segment[0], ": ")
		numbers := strings.Split(leftSegment[1], " ")
		winners := strings.Split(segment[1], " ")
		matchedNumber := 0
		for _, n := range numbers {
			for _, winner := range winners {
				if n == winner && n != "" {
					matchedNumber += 1
					break
				}
			}
		}
		if matchedNumber > 0 {
			sum += 1 << (matchedNumber - 1)
		}

	}
	return sum
}
