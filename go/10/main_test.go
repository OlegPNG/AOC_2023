package main

import "testing"

func TestParseFile(t *testing.T) {
    var test = [][]rune{
        {'.','.','.','.','.',},
        {'.','S','-','7','.',},
        {'.','|','.','|','.',},
        {'.','L','-','J','.',},
        {'.','.','.','.','.',},
    }

    maze := parseFile("sample1.txt")

    for i := 0; i < len(test); i++ {
        for j := 0; j < len(test[i]); j++ {
            expected := test[i][j]
            got := maze.MazeMap[i][j]
            if got != expected {
                t.Errorf("got %v, want %v", got, expected)
            }
        }
    }
}
