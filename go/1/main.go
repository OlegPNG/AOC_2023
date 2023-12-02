package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"time"

	//"slices"
	"unicode"
)

func indexOf(slice []string, value string) (int, bool) {
    for i, v := range slice {
        if v == value {
            return i, true
        }
    }
    return -1, false
}

func findWords(chars[]byte) []int {
    nums := make([]int, 0)
    lowestIndex := 100

    for charIndex := 0; charIndex < (len(chars)); charIndex++ {
        for i := 3; i <= 5; i++ {
            if (charIndex + i) <= len(chars) {
                potentialWord := chars[charIndex:(charIndex + i)]
                num, ok := indexOf(numString(), string(potentialWord))
                if ok {
                    if(charIndex < lowestIndex) {
                        nums = slices.Insert[[]int](nums, 0, num)
                        lowestIndex = charIndex
                    } else {
                        nums = append(nums, num)
                    }
                }
            }
        }
    }
    return nums
}

func numString() []string {
    return []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

}
func main() {
    start := time.Now()
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
        return
    }
    defer file.Close()
    reader := bufio.NewReaderSize(file, 1)
    scanner := bufio.NewScanner(reader)

    total := 0
    for scanner.Scan() {
        line := scanner.Text()
        nums := make([]int, 0)
        currentWord := make([]byte, 0)
        for i, rune := range line {
            if (i == len(line) - 1) {
                currentWord = append(currentWord, byte(rune))
                nums = append(nums, findWords(currentWord)...)
            }
            if unicode.IsDigit(rune) {
                nums = append(nums, findWords(currentWord)...)
                currentWord = make([]byte, 0)
                nums = append(nums, int(rune - '0'))
            } else {
                currentWord = append(currentWord, byte(rune))
            }
        }
        foo := (nums[0] * 10) + nums[len(nums)-1]
        //fmt.Printf("%v: nums %v\n", foo, nums)
        total += foo
    }
    fmt.Printf("Total: %v\n", total)
    
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    duration := time.Since(start)
    fmt.Printf("In: %v\n", duration)
}
