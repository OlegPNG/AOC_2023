package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const sample_race_duration = 7
const sample_distance = 9

func raceLoop(distance, remaining_duration, speed int) bool {
    distTraveled := remaining_duration * speed
    //fmt.Printf("Distance Traveled: %v\n", distTraveled)
    //fmt.Printf("Returning: %v\n", distTraveled >= distance)
    return distTraveled > distance
}
func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    line1 := scanner.Text()
    scanner.Scan()
    line2 := scanner.Text()

    colonIndex := strings.Index(line1, ":")
    line1 = line1[colonIndex + 1:]
    println("Line 1 Before: ", line1)
    line1 = strings.ReplaceAll(line1, " ", "")
    println("Line 1 After: ", line1)
    time, err := strconv.Atoi(line1)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Time: %v\n", time)
    colonIndex = strings.Index(line2, ":")
    line2 = strings.ReplaceAll(line2[colonIndex + 1:], " ", "")
    distance, err := strconv.Atoi(line2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Distance: %v\n", distance)

    total := 0
    speed := 0
    for j := time; j > 0; j-- {
        if raceLoop(distance, j, speed) {
            total += 1
        }
        speed += 1
    }
    
    fmt.Printf("Total: %v\n", total)
}
