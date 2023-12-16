package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseFile(fileName string) [][]int {
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    histories := [][]int{}

    for scanner.Scan() {
        line := scanner.Text()
        histString := strings.Split(line, " ")
        history := []int{}
        for _, v := range(histString) {
            num, err := strconv.Atoi(v)
            if err != nil {
                log.Fatal(err)
            }
            history = append(history, num)
        }
        histories = append(histories, history)
    }

    return histories
}

func sliceIsAllZeroes(slice []int) bool {
    for _, v := range slice {
        if v != 0 { return false }
    }
    return true
}

func getNextValue(history []int) int {
    diffs := []int{}
    for i := 0; i < len(history) - 1; i++ {
        diff := history[i + 1] - history[i]
        diffs = append(diffs, diff)
    }
    nextValue := 0
    if !sliceIsAllZeroes(diffs) {
        nextValue = getNextValue(diffs)
    }
    return history[len(history)-1] + nextValue
}

func getBackValue(history []int) int {
    diffs := []int {}
    for i := 0; i < len(history) - 1; i++ {
        diff := history[i + 1] - history[i]
        diffs = append(diffs, diff)
    }
    backValue := 0
    if !sliceIsAllZeroes(diffs) {
        backValue = getBackValue(diffs)
    }

    return history[0] - backValue
}

func partOne(histories [][]int) {
    total := 0
    for _, v := range histories {
        result := getNextValue(v)
        total += result
    }
    log.Printf("Part One Total: %d\n", total)
}

func partTwo(histories [][]int) {
    total := 0
    for _, v := range histories {
        result := getBackValue(v)
        total += result
    }
    log.Printf("Part Two Total: %d\n", total)
}

func main() {
    histories := parseFile("input.txt")
    partOne(histories)
    partTwo(histories)
}
