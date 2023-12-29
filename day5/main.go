package main

import (
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type almanacMap struct {
	min   int
	max   int
	shift int
}

type almanac struct {
	seeds              []string
	seedToSoil         []almanacMap
	soilToFertilizer   []almanacMap
	fertilizerToWater  []almanacMap
	waterToLight       []almanacMap
	lightToTemp        []almanacMap
	tempToHumidity     []almanacMap
	humidityToLocation []almanacMap
	minLocation        int
}

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
	alm := almanac{}
	alm.seeds = seeds
	alm.minLocation = math.MaxInt
	regexMap := regexp.MustCompile(`(\w+-to-\w+)`)
	currentMapType := ""
	for _, col := range row {
		mapType := regexMap.FindAllString(col, -1)
		if col == "" {
			continue
		}
		if len(mapType) > 0 {
			currentMapType = mapType[0]
			continue
		}

		splitCol := strings.Split(col, " ")
		dest := strToInt(splitCol[0])
		source := strToInt(splitCol[1])
		length := strToInt(splitCol[2])
		switch currentMapType {
		case "seed-to-soil":
			alm.seedToSoil = append(alm.seedToSoil, almanacMap{
				min:   source,
				max:   source + length - 1,
				shift: dest - source,
			})
		case "soil-to-fertilizer":
			alm.soilToFertilizer = append(alm.soilToFertilizer, almanacMap{
				min:   source,
				max:   source + length - 1,
				shift: dest - source,
			})
		case "fertilizer-to-water":
			alm.fertilizerToWater = append(alm.fertilizerToWater, almanacMap{
				min:   source,
				max:   source + length - 1,
				shift: dest - source,
			})
		case "water-to-light":
			alm.waterToLight = append(alm.waterToLight, almanacMap{
				min:   source,
				max:   source + length - 1,
				shift: dest - source,
			})
		case "light-to-temperature":
			alm.lightToTemp = append(alm.lightToTemp, almanacMap{
				min:   source,
				max:   source + length - 1,
				shift: dest - source,
			})
		case "temperature-to-humidity":
			alm.tempToHumidity = append(alm.tempToHumidity, almanacMap{
				min:   source,
				max:   source + length - 1,
				shift: dest - source,
			})
		case "humidity-to-location":
			alm.humidityToLocation = append(alm.humidityToLocation, almanacMap{
				min:   source,
				max:   source + length - 1,
				shift: dest - source,
			})
		}

	}

	P2(alm)
}
func P2(alm almanac) int {
	for i, seed := range alm.seeds {
		seed := strToInt(seed)
		if i%2 == 0 {
			wg := &sync.WaitGroup{}
			loop := strToInt(alm.seeds[i+1])
			for j := 0; j < loop; j++ {
				wg.Add(1)
				j := j
				go func() {
					defer wg.Done()
					newSeed := seed + j
					soil := populateSourceToDest(newSeed, alm.seedToSoil)
					fertilizers := populateSourceToDest(soil, alm.soilToFertilizer)
					water := populateSourceToDest(fertilizers, alm.fertilizerToWater)
					light := populateSourceToDest(water, alm.waterToLight)
					temp := populateSourceToDest(light, alm.lightToTemp)
					humidty := populateSourceToDest(temp, alm.tempToHumidity)
					location := populateSourceToDest(humidty, alm.humidityToLocation)
					if location < alm.minLocation {
						alm.minLocation = location
					}
				}()
			}
			wg.Wait()
		}
	}
	log.Println(alm.minLocation)
	return alm.minLocation
}

func populateSourceToDest(source int, sourceToDest []almanacMap) int {
	for _, almMap := range sourceToDest {
		if source < almMap.min && source > almMap.max {
			return source
		}
		if source >= almMap.min && source <= almMap.max {
			return source + almMap.shift
		}
	}
	return source
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
