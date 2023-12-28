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

func main() {

	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	q, err := os.ReadFile(getwd + "/day5/question")
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
	sourceFound := &sync.Map{}
	mappedSeed := &sync.Map{}
	minLocation := math.MaxInt
	visited := 0
	for _, col := range row {
		if col == "" {
			continue
		}
		mapType := regexMap.FindAllString(col, -1)
		if len(mapType) != 0 {
			currentMapType = mapType[0]
			currentMapType = strings.Split(mapType[0], " ")[0]
			log.Println("processing: ", currentMapType)
			sourceFound = &sync.Map{}
			mappedSeed = &sync.Map{}
			visited = 0
			continue
		}

		switch currentMapType {
		case "seed-to-soil":
			visited = populateSourceFound(visited, seeds, sourceFound, mappedSeed)
			soils = moveSourceToDest(col, seeds, soils, sourceFound, mappedSeed)
		case "soil-to-fertilizer":
			visited = populateSourceFound(visited, soils, sourceFound, mappedSeed)
			fertilizers = moveSourceToDest(col, soils, fertilizers, sourceFound, mappedSeed)
		case "fertilizer-to-water":
			visited = populateSourceFound(visited, fertilizers, sourceFound, mappedSeed)
			waters = moveSourceToDest(col, fertilizers, waters, sourceFound, mappedSeed)
		case "water-to-light":
			visited = populateSourceFound(visited, waters, sourceFound, mappedSeed)
			lights = moveSourceToDest(col, waters, lights, sourceFound, mappedSeed)
		case "light-to-temperature":
			visited = populateSourceFound(visited, lights, sourceFound, mappedSeed)
			temperatures = moveSourceToDest(col, lights, temperatures, sourceFound, mappedSeed)
		case "temperature-to-humidity":
			visited = populateSourceFound(visited, temperatures, sourceFound, mappedSeed)
			humidities = moveSourceToDest(col, temperatures, humidities, sourceFound, mappedSeed)
		case "humidity-to-location":
			visited = populateSourceFound(visited, humidities, sourceFound, mappedSeed)
			locations = moveSourceToDest(col, humidities, locations, sourceFound, mappedSeed)
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

func populateSourceFound(visited int, source []string, sourceFound, mappedSeed *sync.Map) int {
	if visited == 0 {
		for i, sourcesSource := range source {
			sourceFound.Store(fmt.Sprintf("%v|%v", sourcesSource, i), strToInt(sourcesSource))
			mappedSeed.Store(fmt.Sprintf("%v|%v", sourcesSource, i), false)
		}
		return 1
	}
	return visited
}

func moveSourceToDest(col string, sources, destinations []string, sourceFound, mappedSeed *sync.Map) []string {
	splitCol := strings.Split(col, " ")
	dest := strToInt(splitCol[0])
	source := strToInt(splitCol[1])
	length := strToInt(splitCol[2])
	wg := &sync.WaitGroup{}
	for i := 0; i < length; i++ {
		t := length - i - 1
		if i > t {
			break
		}
		if t != i {
			wg.Add(1)
			go mappingSourceFound(sources, mappedSeed, source, i, sourceFound, dest, wg)
		}
		wg.Add(1)
		go mappingSourceFound(sources, mappedSeed, source, t, sourceFound, dest, wg)
	}
	wg.Wait()
	destinations = mapping(sourceFound, len(sources))
	return destinations
}

func mappingSourceFound(sources []string, mappedSeed *sync.Map, source int, i int, sourceFound *sync.Map, dest int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg2 := &sync.WaitGroup{}
	for j, sourcesSource := range sources {
		wg2.Add(1)
		sourcesSource := sourcesSource
		j := j
		go func() {
			sourceKey := fmt.Sprintf("%v|%v", sourcesSource, j)
			isMapped, _ := mappedSeed.Load(sourceKey)
			if strToInt(sourcesSource) == source+i && !isMapped.(bool) {
				sourceFound.Store(sourceKey, dest+i)
				mappedSeed.Store(sourceKey, true)
			}
			defer wg2.Done()
		}()
	}
	wg2.Wait()
}

func mapping(sourceFound *sync.Map, sourceLen int) []string {
	destinations := make([]string, sourceLen)
	sourceFound.Range(func(k, v interface{}) bool {
		segmentK := strings.Split(k.(string), "|")
		destinations[strToInt(segmentK[1])] = fmt.Sprintf("%v", v)
		return true
	})
	return destinations
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
