package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Mapping struct {
    destStart   int
    sourceStart int
    rangeLength int
}

func(m *Mapping) checkForMapping(value int) *int {
    if value >= m.sourceStart && value <= m.sourceStart + m.rangeLength - 1 {
        offset := value - m.sourceStart
        result := m.destStart + offset
        log.Printf("New Mapping: %v\n", result)
        return &result
    }
    //log.Println("No mapping")
    return nil
}

type MappingGroup []Mapping

func(mg MappingGroup) getMapping(value int) int {
    for _, m := range mg {
        num := m.checkForMapping(value)
        if num != nil {
            return *num
        }
    }
    return value
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    seedLine := scanner.Text()
    colonIndex := strings.Index(seedLine, ":") + 2
    seedLine = seedLine[colonIndex:]

    seedStrings := strings.Split(seedLine, " ")
    seeds := []int{}

    for _, v := range seedStrings {
        num, err := strconv.Atoi(v)
        if err != nil {
            log.Fatal(err)
        }
        seeds = append(seeds, num)
    }

    print("Seeds: ")
    for _, v := range seeds {
        fmt.Printf("%v, ", v)
    }
    println()
    
    almanac := []MappingGroup{}
    currentGroup := MappingGroup{}
    scanner.Scan()
    scanner.Scan()

    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            fmt.Printf("Group len: %v\n", len(currentGroup))
            almanac = append(almanac, currentGroup)
            currentGroup = MappingGroup{}
            scanner.Scan()
            continue
        }
        currentMapping := Mapping{}
        numStrings := strings.Split(line, " ")

        for i, v := range numStrings {
            num, err := strconv.Atoi(v)
            if err != nil {
                log.Fatal(err)
            }
            switch i {
            case 0:
                currentMapping.destStart = num
            case 1:
                currentMapping.sourceStart = num
            case 2:
                currentMapping.rangeLength = num
            }
        }
        currentGroup = append(currentGroup, currentMapping)
    }
    almanac = append(almanac, currentGroup)

    var lowest *int

    for _, seed := range seeds {
        currentVal := seed
        for _, group := range almanac {
            currentVal = group.getMapping(currentVal)
        }
        fmt.Printf("Seed %v Location: %v\n", seed, currentVal)
        if lowest == nil {
            lowest = &currentVal
        } else if currentVal < *lowest {
            lowest = &currentVal
        }
    }

    log.Printf("Lowest Location: %v\n", *lowest)

}
