package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	red   int
	green int
	blue  int
	sum   int
}

func main() {
	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	q, err := os.ReadFile(getwd + "/day2/question")
	if err != nil {
		log.Println(err)
		return
	}
	row := strings.Split(string(q), "\n")
	//solveD2P1(row)
	solveD2P2(row)
}

func solveD2P1(row []string) {
	reRGBCount := regexp.MustCompile(`(\d+ (blue|red|green))`)
	reGame := regexp.MustCompile(`Game \d+`)
	finalGame := &Game{}
	for _, col := range row {
		game := reGame.FindAllString(col, -1)
		gameSegment := strings.Split(game[0], " ")
		gameRound := gameSegment[1]
		validSet := true
		for _, rgbSet := range strings.Split(col, ";") {
			g := &Game{}
			rgb := reRGBCount.FindAllString(rgbSet, -1)
			log.Println(rgb)
			for _, s := range rgb {
				rgbSegment := strings.Split(s, " ")
				rgbCount := rgbSegment[0]
				rgbColor := rgbSegment[1]
				switch rgbColor {
				case "green":
					g.green += StrToInt(rgbCount)
				case "red":
					g.red += StrToInt(rgbCount)
				case "blue":
					g.blue += StrToInt(rgbCount)
				}
			}
			if g.red > 12 || g.green > 13 || g.blue > 14 {
				validSet = false
				break
			}
		}
		if validSet {
			finalGame.sum += StrToInt(gameRound)
		}
	}
	log.Println(finalGame.sum)
}
func solveD2P2(row []string) {
	reRGBCount := regexp.MustCompile(`(\d+ (blue|red|green))`)
	sum := 0
	for _, col := range row {
		g := &Game{}
		for _, rgbSet := range strings.Split(col, ";") {
			rgb := reRGBCount.FindAllString(rgbSet, -1)
			for _, s := range rgb {
				rgbSegment := strings.Split(s, " ")
				rgbCount := rgbSegment[0]
				rgbColor := rgbSegment[1]
				switch rgbColor {
				case "green":
					if g.green < StrToInt(rgbCount) {
						g.green = StrToInt(rgbCount)
					}
				case "red":
					if g.red < StrToInt(rgbCount) {
						g.red = StrToInt(rgbCount)
					}
				case "blue":
					if g.blue < StrToInt(rgbCount) {
						g.blue = StrToInt(rgbCount)
					}
				}
			}
		}
		sum += g.green * g.blue * g.red
	}
	log.Println(sum)
}

func StrToInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}
