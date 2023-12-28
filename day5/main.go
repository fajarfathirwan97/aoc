package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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
	soils := []string{}
	fertilizers := []string{}
	waters := []string{}
	lights := []string{}
	temperatures := []string{}
	locations := []string{}
	humidities := []string{}
	currentMapType := ""
	sourceFound := map[string]int{}
	minLocation := math.MaxInt
	for _, col := range row {
		visited := 0
		if col == "" {
			continue
		}
		mapType := regexMap.FindAllString(col, -1)
		if len(mapType) != 0 {
			currentMapType = mapType[0]
			currentMapType = strings.Split(mapType[0], " ")[0]
			sourceFound = map[string]int{}
			visited = 0
			continue
		}

		switch currentMapType {
		case "seed-to-soil":
			populateSourceFound(visited, seeds, sourceFound)
			soils = moveSourceToDest(col, seeds, soils, sourceFound)
			log.Println("seeds", seeds, "soils", soils, "sourceFound", sourceFound)
		case "soil-to-fertilizer":
			populateSourceFound(visited, soils, sourceFound)
			fertilizers = moveSourceToDest(col, soils, fertilizers, sourceFound)
		case "fertilizer-to-water":
			populateSourceFound(visited, fertilizers, sourceFound)
			waters = moveSourceToDest(col, fertilizers, waters, sourceFound)
		case "water-to-light":
			populateSourceFound(visited, waters, sourceFound)
			lights = moveSourceToDest(col, waters, lights, sourceFound)
		case "light-to-temperature":
			populateSourceFound(visited, lights, sourceFound)
			temperatures = moveSourceToDest(col, lights, temperatures, sourceFound)
		case "temperature-to-humidity":
			populateSourceFound(visited, temperatures, sourceFound)
			humidities = moveSourceToDest(col, temperatures, humidities, sourceFound)
		case "humidity-to-location":
			populateSourceFound(visited, humidities, sourceFound)
			locations = moveSourceToDest(col, humidities, locations, sourceFound)
		}
	}

	for i, _ := range seeds {
		log.Println("seed",
			seeds[i],
			"soil",
			soils[i],
			"fertilizer",
			fertilizers[i],
			"water",
			waters[i],
			"light",
			lights[i],
			"temperature",
			temperatures[i],
			"humidity",
			humidities[i],
			"location",
			locations[i],
		)
		location := strToInt(locations[i])
		if minLocation > location {
			minLocation = location
		}
	}

	log.Println(minLocation)
}

func populateSourceFound(visited int, source []string, sourceFound map[string]int) {
	if visited == 0 {
		visited = 1
		for i, sourcesSource := range source {
			sourceFound[fmt.Sprintf("%v|%v", sourcesSource, i)] = strToInt(sourcesSource)
		}
	}
}

func moveSourceToDest(col string, sources, destinations []string, sourceFound map[string]int) []string {
	splitCol := strings.Split(col, " ")
	dest := strToInt(splitCol[0])
	source := strToInt(splitCol[1])
	length := strToInt(splitCol[2])
	for i := 0; i < length; i++ {
		for j, sourcesSource := range sources {
			sourceKey := fmt.Sprintf("%v|%v", sourcesSource, j)
			if strToInt(sourcesSource) == source+i {
				sourceFound[sourceKey] = dest + i
			}
		}
	}
	destinations = mapping(sourceFound, len(sources))
	return destinations
}

func mapping(sourceFound map[string]int, sourceLen int) []string {
	destinations := make([]string, sourceLen)
	for k, dest := range sourceFound {
		segmentK := strings.Split(k, "|")
		destinations[strToInt(segmentK[1])] = fmt.Sprintf("%v", dest)
	}
	return destinations
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
