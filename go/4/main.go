package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Card struct {
    index   int
    winning []string
    mine    []string
}

func findNumMatches(winning, mine []string) int {
    total := 0
    for _, v := range mine {
        if v != "" {
            if slices.Contains(winning, v) {
                total += 1
            } 
        }
    }
    return total
}
func findMatches(winning, mine []string) int {
    total := 0
    matches := []string{}
    for _, v := range mine {
        if v != "" {
            if slices.Contains(winning, v) {
                matches = append(matches, v)
                if total == 0 {
                    total = 1
                } else {
                    total *= 2
                }
            } 
        }
    }
    print("Matches: ")
    for _, v := range matches {
        fmt.Printf("%v ", v)
    }
    println()
    fmt.Printf("Num of Matches: %v\n", len(matches))
    return total
}

func findMatchesRecursive(card Card, cards []Card, memo map[int]int) int {
    memVal, ok := memo[card.index]
    if ok {

    fmt.Printf("Memoized Card %v: %v matches\n", card.index, memVal)
        return memVal
    }
    matches := findNumMatches(card.winning, card.mine)
    fmt.Printf("Card %v: %v matches\n", card.index, matches)
    if matches != 0 {
        total := 1
        for i := 0; i < matches; i++ {
            nextCard := cards[card.index + i]
            total += findMatchesRecursive(nextCard, cards, memo)
        }
        memo[card.index] = total
        fmt.Printf("Total Cards added for Card %v: %v\n", card.index, total)
        return total
    }
    return 1
}

func main() {
    file, err := os.Open("input.txt") 
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    cardIndex := 0
    total := 0
    cards := []Card{}
    //additionalCards := []Card{}

    // Maps card Index with number of cards added
    memo := map[int]int{}

    for scanner.Scan() {
        cardIndex += 1
        line := scanner.Text()
        colonIndex := strings.Index(line, ":")

        card := line[colonIndex:]
        fmt.Printf("Card %v: %v\n", cardIndex, card)
        
        sepLists := strings.Split(card, " | ")
        winning := strings.Split(sepLists[0], " ")
        myNums := strings.Split(sepLists[1], " ")
        cards = append(cards, Card{
            index: cardIndex,
            winning: winning,
            mine: myNums,
        })
        
        //total += findMatches(winning, myNums)
        //log.Printf("Total: %v\n", total)
    }

    for _, card := range cards {
        fmt.Printf("****Loop for Card %v****\n", card.index)
        total += findMatchesRecursive(card, cards, memo)
    }
    //fmt.Printf("# of Original Cards: %v\n", len(cards))
    //total += len(cards)

    /*
    for _, card := range cards {
        cardsToAdd := findNumMatches(card.winning, card.mine)
        memo[card.index] = cardsToAdd
        for i := 1; i < cardsToAdd && i < len(cards); i++ {
            additionalCards = append(additionalCards, cards[card.index + i])
        }
    }
    for _, card := range additionalCards {
        cardsToAdd := memo[card.index]
        for i := 1; i < cardsToAdd && i < len(cards); i++ {
            additionalCards = append(additionalCards, cards[card.index + i])
        }
    }
    for k, v := range memo {
        fmt.Printf("Card #%v adds %v cards\n", k, v)
    }
    total += len(additionalCards)
    */

    log.Printf("Total: %v\n", total)

}
