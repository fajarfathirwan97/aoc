package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var mux = sync.RWMutex{}

func main() {

	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	q, err := os.ReadFile(getwd + "/day5/example")
	if err != nil {
		log.Println(err)
		return
	}

	parseFile(string(q))
}

func parseFile(s string) {
	row := strings.Split(s, "\n")
	seedsStr := strings.Split(strings.ReplaceAll(row[0], "seeds: ", ""), ": ")[0]
	seeds := strings.Split(seedsStr, " ")
	regexMap := regexp.MustCompile(`([a-z\-]+) map:`)
	regexNumber := regexp.MustCompile(`[0-9]+`)
	soils := []string{}
	fertilizers := []string{}
	waters := []string{}
	lights := []string{}
	temperatures := []string{}
	locations := []string{}
	humidities := []string{}
	currentMapType := ""
	sourceToDest := map[int]int{}
	minLocation := math.MaxInt
	for _, col := range row {
		if col == "" {
			continue
		}
		mapType := regexMap.FindAllString(col, -1)
		if len(mapType) != 0 {
			currentMapType = mapType[0]
			currentMapType = strings.Split(mapType[0], " ")[0]
			sourceToDest = map[int]int{}
			continue
		}

		switch currentMapType {
		case "seed-to-soil":
			soils = moveSourceToDest(col, regexNumber, seeds, soils, sourceToDest)
		case "soil-to-fertilizer":
			fertilizers = moveSourceToDest(col, regexNumber, soils, fertilizers, sourceToDest)
		case "fertilizer-to-water":
			waters = moveSourceToDest(col, regexNumber, fertilizers, waters, sourceToDest)
		case "water-to-light":
			lights = moveSourceToDest(col, regexNumber, waters, lights, sourceToDest)
		case "light-to-temperature":
			temperatures = moveSourceToDest(col, regexNumber, lights, temperatures, sourceToDest)
		case "temperature-to-humidity":
			humidities = moveSourceToDest(col, regexNumber, temperatures, humidities, sourceToDest)
		case "humidity-to-location":
			locations = moveSourceToDest(col, regexNumber, humidities, locations, sourceToDest)
		}
	}

	for i, _ := range seeds {
		location := strToInt(locations[i])
		if minLocation > location {
			minLocation = location
		}
	}

	log.Println(minLocation)
}

func moveSourceToDest(col string, regexNumber *regexp.Regexp, sources, destinations []string, sourceToDest map[int]int) []string {
	splitCol := strings.Split(col, " ")
	numbers := regexNumber.FindAllString(splitCol[0], -1)
	if len(numbers) != 0 {
		dest := strToInt(splitCol[0])
		source := strToInt(splitCol[1])
		length := strToInt(splitCol[2])
		for i := 0; i < length; i++ {
			sourceToDest[source+i] = dest + i
		}
		destinations = mapping(sources, sourceToDest)
	}
	return destinations
}

func mapping(sources []string, sourceToDest map[int]int) []string {
	destinations := []string{}
	//destinations := make([]string, len(sources))
	for _, source := range sources {
		source := source
		//go func() {
		//	mux.Lock()
		//	mux.RLock()
		v, ok := sourceToDest[strToInt(source)]
		if !ok {
			v = strToInt(source)
		}
		destinations = append(destinations, fmt.Sprintf("%v", v))
		//	defer mux.RUnlock()
		//	defer mux.Unlock()
		//}()
	}
	return destinations
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
