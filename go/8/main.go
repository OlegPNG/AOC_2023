package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type Instruction int
const (
    Left Instruction = iota
    Right
)

type Data struct {
    Instructions    []Instruction 
    StartNodes      []Node
    Network         Network
}

type Node struct {
    Index   string
    Left    string
    Right   string
}

type Network map[string]Node

func parseFile(fileName string) Data {
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    line := scanner.Text()

    instructions := []Instruction{}
    for _, c := range(line) {
        //println(char)
        if c == 'L' {
            instructions = append(instructions, Left)
        } else {
            instructions = append(instructions, Right)
        }
    }
    scanner.Scan()

    network := Network{}
    startNodes := []Node{}
    for scanner.Scan() {
        line := scanner.Text()
        equalIndex := strings.Index(line, "=")
        
        nodeIndex := line[:equalIndex - 1]

        lParen := strings.Index(line, "(") + 1
        rParen := strings.Index(line, ")")
        //foo := line[lParen:rParen]
        nodeStrings := strings.Split(line[lParen:rParen], ", ")
        //fmt.Printf("(%v)", foo)

        node := Node{
            Index: nodeIndex,
            Left: nodeStrings[0],
            Right: nodeStrings[1],
        }
        network[nodeIndex] = node
        if strings.Contains(nodeIndex, "A") {
            startNodes = append(startNodes, node)
        }
    }

    return Data{
        Instructions: instructions,
        Network: network,
        StartNodes: startNodes,
    }

}

func partOne(data Data) {
    node := data.Network["AAA"]
    count := 0
    for node.Index != "ZZZ" {
        instruction := data.Instructions[count % len(data.Instructions)]
        if instruction == Left {
            node = data.Network[node.Left]
        } else {
            node = data.Network[node.Right]
        }
        count++
    }
    fmt.Printf("Total: %d\n", count)
}

func partTwo(data Data) {
    nodes := data.StartNodes

    resultPointers := []*int{}

    var wg sync.WaitGroup
    for _, node := range(nodes) {
        wg.Add(1)
        var result int
        go thread(&result, node, data, &wg)
        resultPointers = append(resultPointers, &result)
    }
    wg.Wait()
    results := []int{}
    for _, num := range(resultPointers) {
        results = append(results, *num)
    }

    a := results[0]
    b := results[1]
    results = results[2:]
    total := LCM(a, b, results...)
    fmt.Printf("Result: %d\n", total)
}

func thread(result *int, start Node, data Data, wg *sync.WaitGroup) {
    defer wg.Done()
    count := 0
    node := start
    for !checkFinished(node) {
        instruction := data.Instructions[count % len(data.Instructions)]
        if instruction == Left {
            node = data.Network[node.Left]
        } else {
            node = data.Network[node.Right]
        }
        count ++
    }
    *result = count
}

func checkFinished(node Node) bool {
    index := node.Index
    return index[len(index) - 1] == 'Z' 
}
func checkAllFinished(nodes []Node) bool {
    for _, node := range(nodes) {
        if index := node.Index; index[len(index) - 1] != 'Z' {
            return false
        }
    }
    return true
}

func main() {
    data := parseFile("input.txt")
    //partOne(data)
    partTwo(data)
}

func GCD(a, b int) int {
    for b!= 0 {
        t := b
        b = a % b
        a = t
    }
    return a
}

func LCM(a, b int, integers ...int) int {
    result := a * b / GCD(a, b)
    for i := 0; i < len(integers); i ++ {
        result = LCM(result, integers[i])
    }

    return result
}
