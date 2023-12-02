package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Color int
const (
    Blue Color = iota
    Green
    Red
)

const (
    redTotal int = 12
    greenTotal int = 13
    blueTotal int = 14
)

func getPower(game string, id int) int {
    redMax := 0
    greenMax := 0
    blueMax := 0 

    sets := strings.Split(game, "; ")
    fmt.Printf("Game %v: %v\n", id, game)

    for _, v := range sets {
        groups := strings.Split(v, ", ")
        for _, g := range groups {
            spaceIndex := strings.Index(g, " ")
            count, err := strconv.Atoi(string(g[0: spaceIndex]))
            if err != nil {
                log.Fatalf("Could not convert count: %v to int: %v", g[0: spaceIndex], err)
            }
            color := g[spaceIndex + 1:]

            switch color {
            case "red":
                if count > redMax {
                    redMax = count
                }
            case "green":
                if count > greenMax {
                    greenMax = count
                }
            case "blue":
                if count > blueMax {
                    blueMax = count
                }
            }
        }
    }

    return redMax * greenMax * blueMax

}
// Returns true if possible given color totals
func checkGame(game string, id int) bool {

    redMax := 0
    greenMax := 0
    blueMax := 0 

    sets := strings.Split(game, "; ")
    fmt.Printf("Game %v: %v\n", id, game)
    for _, v := range sets {
        //fmt.Printf("Set: %v\n", v)

        groups := strings.Split(v, ", ")
        for _, g := range groups {
            spaceIndex := strings.Index(g, " ")
            count, err := strconv.Atoi(string(g[0: spaceIndex]))
            if err != nil {
                log.Fatalf("Could not convert count: %v to int: %v", g[0: spaceIndex], err)
            }
            color := g[spaceIndex + 1:]

            switch color {
            case "red":
                if count > redTotal {
                    log.Printf("Too many red: %v", count)
                    return false
                } else {
                    if count > redMax {
                        redMax = count
                    }
                }

            case "green":
                if count > greenTotal {
                    log.Printf("Too many green: %v", count)
                    return false
                } else {
                    if count > greenMax {
                        greenMax = count
                    }
                }
                
            case "blue":
                if count > blueTotal {
                    log.Printf("Too many blue: %v", count)
                    return false
                } else {
                    if count > blueMax {
                        blueMax = count
                    }
                }
            }
        }
    }

    /*
    fmt.Printf("Red Max: %v\n", redMax)
    fmt.Printf("Green Max: %v\n", greenMax)
    fmt.Printf("Blue Max: %v\n", blueMax)
    */

    return true
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    total := 0
    for scanner.Scan() {
        game := scanner.Text()
        colonIndex := strings.IndexAny(game, ":")
        id_string := game[5:colonIndex]
        id, err := strconv.Atoi(id_string)
        if err != nil {
            log.Fatalf("Could not get id as integer: %v\n", err)
        }
        game = strings.TrimPrefix(game, game[:colonIndex + 2])
        total += getPower(game, id)
        //if checkGame(game, id) {
        //    log.Printf("Success id: %v\n", id)
        //    total += id
        //} else {
        //    log.Printf("Fail id: %v\n", id)
        //}

    }

    log.Printf("Total: %v\n", total)
    
}
