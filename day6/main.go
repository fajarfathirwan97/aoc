package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type race struct {
	times    []int
	distance []int
	result   int
}

func main() {
	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	q, err := os.ReadFile(getwd + "/day6/example")
	if err != nil {
		log.Println(err)
		return
	}

	row := strings.Split(string(q), "\n")
	time := strings.ReplaceAll(row[0], "Time:", "")
	distance := strings.ReplaceAll(row[1], "Distance:", "")
	r := &race{}
	r.times = parseFileP2(time)
	r.distance = parseFileP2(distance)
	r.result = 1
	//solveP1(r)
	solveP2(r)
	log.Println(r.result)

}

func solveP1(r *race) {
	for i := 0; i < len(r.times); i++ {
		time := r.times[i]
		distance := r.distance[i]
		totalWin := 0
		for speed := 1; speed < time; speed++ {
			timeRemaining := time - speed
			finalDistance := speed * timeRemaining
			if finalDistance > distance {
				totalWin++
			}
		}
		r.result *= totalWin
	}
}
func solveP2(r *race) {
	for i := 0; i < len(r.times); i++ {
		time := r.times[i]
		distance := r.distance[i]
		totalWin := 0
		for speed := 1; speed < time; speed++ {
			timeRemaining := time - speed
			finalDistance := speed * timeRemaining
			if finalDistance > distance {
				totalWin++
			}
		}
		r.result *= totalWin
	}
}

func parseFileP2(s string) []int {
	s = strings.Replace(s, " ", "", -1)
	i := strToInt(s)
	return []int{i}
}
func parseFile(s string) []int {
	timeInt := 0
	strFormat := ""
	result := []int{}
	for i, t := range s {
		if t == 32 {
			strFormat = ""
			if timeInt > 0 {
				result = append(result, timeInt)
			}
			timeInt = 0
		} else {
			strFormat += string(t)
			timeInt = strToInt(strFormat)
			if i == len(s)-1 {
				result = append(result, timeInt)
			}
		}
	}
	return result
}
func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
