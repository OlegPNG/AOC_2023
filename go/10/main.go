package main

import (
	"bufio"
	"log"
	"math"
	"os"
)

type Coord struct {
    x int
    y int
}

type Maze struct {
    MazeMap     [][]rune
    AnimalPos   Coord
}

func parseFile(fileName string) Maze {
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)

    var position Coord
    var maze [][]rune
    yPos := 0
    for scanner.Scan() {
        line := scanner.Text()
        for xPos, val := range line {
            if val == 'S' {
                position = Coord{ x: xPos, y: yPos }
            }
        }
        maze = append(maze, []rune(line))

        yPos++
    }

    return Maze{
        MazeMap: maze,
        AnimalPos: position,
    }
}

func findNeighboringPipe(position Coord, maze [][]rune) Coord {
    length := 0
    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            y := position.y + i
            x := position.x + j
            if i == 0  {
                if x >= 0 && x <len(maze[0]) {
                    switch j {
                    case -1:
                        if maze[y][x] == '-' {

                        }
                    }
                }
            }
        }
    }
}

func main() {

}
