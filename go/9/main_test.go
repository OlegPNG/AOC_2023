package main

import (
	"testing"
)

func TestFileParsing(t *testing.T) {
    var tests = [][]int{
	{0, 3, 6, 9, 12, 15},
	{1, 3, 6, 10, 15, 21},
	{10, 13, 16, 21, 30, 45},
    }
    histories := parseFile("sample.txt")

    for i := 0; i < 3; i++ {
	for j := 0; j < len(tests[0]); j++ {
	    want := tests[i][j]
	    got := histories[i][j]
	    if want != got {
		t.Errorf("got %d, want %d", got, want)
	    }
	}
    }
}

func TestNextValue(t *testing.T) {
    var tests = []struct {
	input []int
	want int
    }{
	{
	    []int{0, 3, 6, 9, 12, 15},
	    18,
	},
	{
	    []int{1, 3, 6, 10, 15, 21},
	    28,
	},
	{
	    []int{10, 13, 16, 21, 30, 45},
	    68,
	},
    }
    for _, tt := range tests {
	result := getNextValue(tt.input)
	if result != tt.want {
	    t.Errorf("got %d, want %d", result, tt.want)
	}
    }
}

func TestBackValue(t *testing.T) {
    var tests = []struct {
	input []int
	want int
    }{
	{
	    []int{0, 3, 6, 9, 12, 15},
	    -3,
	},
	{
	    []int{1, 3, 6, 10, 15, 21},
	    0,
	},
	{
	    []int{10, 13, 16, 21, 30, 45},
	    5,
	},
    }
    for _, tt := range tests {
	result := getBackValue(tt.input)
	if result != tt.want {
	    t.Errorf("got %d, want %d", result, tt.want)
	}
    }
}
