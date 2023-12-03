package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func potentialGear(engine [][]rune, y, x int) int {
    nums := []int{}

    for yShift := -1; yShift <= 1; yShift ++ {
        yIndex := y + yShift
        if yIndex >= 0 && yIndex <= len(engine) -1 {
            for xShift := -1; xShift <= 1; xShift ++ {
                xIndex := x + xShift
                if xIndex >= 0 && xIndex <= len(engine[yIndex]) -1 {
                    if unicode.IsNumber(engine[yIndex][xIndex]) {
                        for xIndex > 0 && unicode.IsNumber(engine[yIndex][xIndex - 1]) {
                            xIndex -= 1
                        }
                        numString := []rune{}
                        for xIndex <= len(engine[yIndex]) - 1 && unicode.IsNumber(engine[yIndex][xIndex]) {
                            numString = append(numString, engine[yIndex][xIndex])
                            xIndex += 1
                        }
                        num, err := strconv.Atoi(string(numString))
                        if err != nil {
                            log.Fatalf("Failed to parse number %v\n", string(numString))
                        }
                        nums = append(nums, num)
                        xShift = xIndex - x
                    }
                }
            }
        }
    }
    if len(nums) == 2 {
        return nums[0] * nums[1]
    } else {
        return 0
    }
}
func findNums(engine [][]rune, y, x int) int {
    subtotal := 0
    for yShift := -1; yShift <= 1; yShift ++ {
        yIndex := y + yShift
        if yIndex >= 0 && yIndex <= len(engine) -1 {
            for xShift := -1; xShift <= 1; xShift ++ {
                xIndex := x + xShift
                if xIndex >= 0 && xIndex <= len(engine[yIndex]) -1 {
                    if unicode.IsNumber(engine[yIndex][xIndex]) {
                        for xIndex > 0 && unicode.IsNumber(engine[yIndex][xIndex - 1]) {
                            xIndex -= 1
                        }
                        numString := []rune{}
                        for xIndex <= len(engine[yIndex]) - 1 && unicode.IsNumber(engine[yIndex][xIndex]) {
                            numString = append(numString, engine[yIndex][xIndex])
                            xIndex += 1
                        }
                        num, err := strconv.Atoi(string(numString))
                        if err != nil {
                            log.Fatalf("Failed to parse number %v\n", string(numString))
                        }
                        subtotal += num
                        xShift = xIndex - x
                    }
                }
            }
        }
    }
    return subtotal
}

func main() {
    //symbols := "@#$%&*-=+/"
    engine := [][]rune{}
    
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        engine = append(engine, []rune(scanner.Text()))
    }

    total := 0
    for y, line := range engine {
        for x, char := range line {
            if char == '*' {
                //total += findNums(engine, y, x)  
                total += potentialGear(engine, y, x)
            }
        }
    }

    log.Printf("Total: %v\n", total)
}
