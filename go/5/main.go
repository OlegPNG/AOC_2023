package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type MapGroup [][]int

type SeedRange struct {
    Start int
    Length int
}

type Interval struct {
    Start int
    End int
}

type Data struct {
    Seeds []int
    SeedRanges []Interval
    Groups []MapGroup
}

func parseFile(fileName string) Data {
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)

    scanner.Scan()
    line := scanner.Text()
    seedStrings := strings.Split(line[7:], " ")
    var seeds []int
    var seedRanges []Interval
    var currentRange = Interval{}
    for i, v := range seedStrings {
        num, err := strconv.Atoi(v)
        if err != nil {
            log.Fatal(err)
        }
        if i % 2 == 0 {
            currentRange.Start = num
        } else {
            currentRange.End = currentRange.Start + num - 1
            seedRanges = append(seedRanges, currentRange)
            currentRange = Interval{}
        }
        seeds = append(seeds, num)
    }
    scanner.Scan()
    scanner.Scan()
    
    data := Data{
        Seeds: seeds,
        SeedRanges: seedRanges,
    }
    currentGroup := MapGroup{}
    for scanner.Scan() {
        currentMapping := []int{}
        line := scanner.Text()
        if line == "" {
            data.Groups = append(data.Groups, currentGroup)
            currentGroup = MapGroup{}
            scanner.Scan()
            scanner.Scan()
            line = scanner.Text()
        }
        slice := strings.Split(line, " ")
        for _, v := range slice {
            num, err := strconv.Atoi(v)
            if err != nil {
                log.Fatal(err)
            }
            currentMapping = append(currentMapping, num)
        }
        currentGroup = append(currentGroup, currentMapping)
    }
    return data
}

// returns mapping and isMapped
func checkForMapping(value int, mapping []int) (int, bool) {
    dStart := mapping[0]
    sStart := mapping[1]
    rangeLen := mapping[2]

    if value >= sStart && value <= sStart + rangeLen - 1 {
        diff := value - sStart
        return dStart + diff, true
    }
    return 0, false
}

//func checkMappingForRange(seeds SeedRange, mapping []int) ()

func getMappedValue(value int, group MapGroup) int {
    for _, mapping := range group {
        result, isMapped := checkForMapping(value, mapping)
        if isMapped {
            return result
        }
    }
    return value
}

func partOne(data Data) {
    var smallest *int = nil
    current := 0
    for _, seed := range data.Seeds {
        current = seed
        for _, group := range data.Groups {
            current = getMappedValue(current, group)
        }
        if smallest == nil {
            smallest = new(int)
            *smallest = current
        } else if current < *smallest {
            *smallest = current
        }
        fmt.Printf("Seed %d: %d\n", seed, current)
    }
    log.Printf("Final: %d\n", *smallest)
}

func partTwo(data Data) {
    var smallest *int = nil
    var current int
    for _, seedRange := range data.SeedRanges {
        for seed := seedRange.Start; seed < seedRange.End; seed++ {
            fmt.Printf("Counter: %d\n", seed)
            //seed := seedRange.Start + diff
            current = seed
            for _, group := range data.Groups {
                current = getMappedValue(current, group)
            }
            if smallest == nil {
                smallest = new(int)
                *smallest = current
            }
            if current < *smallest {
                //fmt.Printf("New Smallest: %d\n", current)
                *smallest = current
            }
            //fmt.Printf("Seed %d: %d\n", seed, current)
        }
    }
    log.Printf("Final: %d\n", *smallest)
}

func main() {
    data := parseFile("sample.txt")
    //partOne(data)
    partTwo(data)
}
