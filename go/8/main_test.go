package main

import "testing"

func TestFileParsing(t *testing.T) {
    var test = []Instruction{Right, Left}

    var testSN = []string{"11A", "22A"}
        
    parsedFile := parseFile("sample_1.txt")
    result := parsedFile.Instructions


    for i := 0; i < len(result); i++ {
        if result[i] != test[i] {
            t.Errorf("%d: got %d, want %d", i, result[i], test[i])
        }
    }

    newParsedFile := parseFile("sample_2.txt")
    resultSN := newParsedFile.StartNodes
    for i := 0; i < len(resultSN); i++ {
        if sn := resultSN[i].Index; sn != testSN[i] {
            t.Errorf("StartNode: got %v, want %v", sn, testSN[i])
        }
    }
}
