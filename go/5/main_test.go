package main

import "testing"

func TestCheckForMapping(t *testing.T) {
    var tests = []struct {
        value, want int
        mapping []int
    }{
        {79, 81, []int{52, 50, 48}},
        {98, 50, []int{50, 98, 2}},
    }

    for i, tt := range tests {
        result, isMapped := checkForMapping(tt.value, tt.mapping)
        if !isMapped {
            t.Errorf("test %d: val %d is not mapped", i, tt.value)
        }
        if result != tt.want {
            t.Errorf("got: %d, want: %d", result, tt.want)
        }
    }
}

func TestGetMappedValue(t *testing.T) {
    group := MapGroup{
        {50, 98, 2},
        {52, 50, 48},
    }
    var tests = []struct{
        value, want int
    }{
        {79, 81},
        {55, 57},
        {13, 13},
    }

    for _, tt := range tests {
        result := getMappedValue(tt.value, group)
        if result != tt.want {
            t.Errorf("got: %d, want: %d", result, tt.want)
        }
    }
}

