package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type instruction int

const (
    Left instruction = iota
    Right
)

type Element struct {
    Index   string
    Left    string
    Right   string
}

func(e *Element) toString() string {
    return fmt.Sprintf("%v: (%v, %v)", e.Index, e.Left, e.Right)
}

type Document struct {
    Instructions    []instruction 
    Elements        map[string]*Element
    CurrentElement  *Element
    CurrentElementList []*Element
}

func parseFile(scanner *bufio.Scanner) Document {
    instructions := []instruction{}
    elements := map[string]*Element{}
    startingElements := []*Element{}
    var current *Element
    scanner.Scan()
    line := scanner.Text()
    for _, char := range(line) {
        if char == 'L' {
            instructions = append(instructions, Left)
        } else if char == 'R' {
            instructions = append(instructions, Right)
        }
    }
    scanner.Scan()

    //Get first element
    //scanner.Scan()
    //line = scanner.Text()
    //key, current := parseElement(line)
    //elements[key] = &current
    //log.Printf("(%v) (%v, %v)\n\n", key, current.Left, current.Right)

    for scanner.Scan() {
        line := scanner.Text()
        key, element := parseElement(line)
        elements[key] = &element
        if key[2] == 'A' {
            startingElements = append(startingElements, &element)
        }
        //log.Printf("Adding element: %v\n", element.toString())
        //log.Printf("(%v) (%v, %v)\n\n", key, element.Left, element.Right)
    }

    current = elements["AAA"]

    return Document{
        Instructions: instructions,
        Elements: elements,
        CurrentElement: current,
        CurrentElementList: startingElements,
    }
}

// Returns key and element
func parseElement(line string) (string, Element) {
    line = strings.ReplaceAll(line, " ", "")
    index := line[:strings.Index(line, "=")]

    left := line[strings.Index(line, "(") + 1 : strings.Index(line, ",")]

    right := line[strings.Index(line, ",") + 1 : strings.Index(line, ")")]

    element := Element{
        Index: index,
        Left: left,
        Right: right,
    }
    return index, element
}

func partOne(document Document) {
    counter := 0
    for document.CurrentElement.Index != "ZZZ" {
        instruction := document.Instructions[counter % len(document.Instructions)]
        switch instruction {
        case Left:
            println("L ")
            newElement, ok := document.Elements[document.CurrentElement.Left]
            fmt.Printf("New Element: %v\n", newElement.toString())
            if !ok {
                log.Fatal("Element not found")
            }
            document.CurrentElement = newElement
        case Right:
            println("R ")
            newElement, ok := document.Elements[document.CurrentElement.Right]
            fmt.Printf("New Element: %v\n", newElement.toString())
            if !ok {
                log.Fatal("Element not found")
            }
            document.CurrentElement = newElement
        }
        counter++
    }
    log.Printf("Steps: %d\n", counter)
}

func checkAllRoutesComplete(elements []*Element) bool {
    //for _, element := range(elements) {
    //    fmt.Printf("%v,", element.Index)
    //}
    //println()
    for _, element := range(elements) {
        if element.Index[2] != 'Z' {
            return false
        }
        //println()
    }
    return true
}

func partTwo(document Document) {
    counter := 0
    println(len(document.CurrentElementList))
    for !checkAllRoutesComplete(document.CurrentElementList) {
        fmt.Printf("Step %d: \n", counter)
        instruction := document.Instructions[counter % len(document.Instructions)]
        
        switch instruction {
        case Left:
            for i, element := range(document.CurrentElementList) {
                newElement := document.Elements[element.Left]
                document.CurrentElementList[i] = newElement
                //fmt.Printf("New Element: %v\n", newElement.toString())
            }
        case Right:
            for i, element := range(document.CurrentElementList) {
                newElement := document.Elements[element.Right]
                document.CurrentElementList[i] = newElement
                //fmt.Printf("New Element: %v\n", newElement.toString())
            }
        }
        counter++
    }

    log.Println("Steps: ", counter)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    document := parseFile(scanner)
    //partOne(document)
    partTwo(document)
}
