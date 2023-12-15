package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var CardValues = map[byte]int{
    'J': 0,
    '2': 1,
    '3': 2,
    '4': 3,
    '5': 4,
    '6': 5,
    '7': 6,
    '8': 7,
    '9': 8,
    'T': 9,
    'Q': 10,
    'K': 11,
    'A': 12,
}

type Hand struct {
    handType    HandType
    cards       string
    bid         int
}

// Returns true if current hand > next hand
func(hand Hand) compareHands(other Hand) bool {
    if hand.handType > other.handType {
        return true
    }
    if hand.handType < other.handType {
        return false
    }
    for i := 0; i < len(hand.cards); i++ {
        if hand.cards[i] == other.cards[i] {
            continue
        }
        if CardValues[hand.cards[i]] > CardValues[other.cards[i]] {
            return true
        }
        return false
    }
    return false
}

type HandType int

const (
    HighCard HandType = iota
    OnePair 
    TwoPair
    ThreeKind
    FullHouse // ThreeKind + OnePair
    FourKind
    FiveKind
)

func(h HandType) toString() string {
    switch h {
    case HighCard:
        return "HighCard"
    case OnePair:
        return "OnePair"
    case TwoPair:
        return "TwoPair"
    case ThreeKind:
        return "ThreeKind"
    case FullHouse:
        return "FullHouse"
    case FourKind:
        return "FourKind"
    case FiveKind:
        return "FiveKind"
    default:
        return "None"
    }
}

func bubbleSort(hands []Hand) []Hand {
    for i := len(hands) - 1; i > 0; i-- {
        for j := 0; j < i; j++ {
            leftHand := hands[j]
            rightHand := hands[j + 1]
            if leftHand.compareHands(rightHand) {
                hands[j], hands[j+1] = hands[j+1], hands[j]
            }
        }
    }
    return hands
}

func splitLine(s string) (string, string) {
    x := strings.Split(s, " ")
    return x[0], x[1]
}

func getHandType(hand string) HandType {
    jokerCount := 0
    cardMap := map[rune]int{}
    for _, card := range(hand) {
        if card == 'J' {
            jokerCount += 1
        }
        cardMap[card] += 1
    }
    threeKindPresent := false
    onePairPresent := false
    handType := HighCard
    //log.Println("Hand:")
    for key, value := range(cardMap) {
        //log.Printf("Card: %v | Total: %d", string(key), value)
        var cardTotal int
        if key != 'J' {
            cardTotal = value + jokerCount
        } else {
            cardTotal = value
        }
        switch cardTotal {
        case 2:
            if onePairPresent {
                //handType = TwoPair
                //return handType
                if TwoPair > handType {
                    handType = TwoPair
                    continue
                }
            }
            if threeKindPresent {
                //handType = FullHouse
                //return handType
                if FullHouse > handType {
                    handType = FullHouse
                    continue
                }
            }
            onePairPresent = true
            if OnePair > handType {
                handType = OnePair
            }

        case 3:
            if onePairPresent {
                //handType = FullHouse
                //return handType
                if FullHouse > handType {
                    handType = FullHouse
                    continue
                }
            }
            threeKindPresent = true
            if ThreeKind > handType {
                handType = ThreeKind
            }
        case 4:
            //handType = FourKind
            //return handType
            if FourKind > handType {
                handType = FourKind
            }
        case 5:
            handType = FiveKind
            return handType
        }
    }
    return handType
}

func partOne(scanner *bufio.Scanner) {
    hands := []Hand{}
    counter := 0
    for scanner.Scan() {
        counter++
        line := scanner.Text()
        cards, bidString := splitLine(line)
        bid, err := strconv.Atoi(strings.ReplaceAll(bidString, " ", ""))
        if err != nil {
            log.Fatal(err)
        }
        hand := Hand{
            handType: getHandType(cards),
            cards: cards,
            bid: bid,
        }
        //log.Printf("Hand %d Type: %v\n", counter, hand.handType.toString())
        hands = append(hands, hand)
    }
    hands = bubbleSort(hands)
    total := 0
    for index, hand := range(hands) {
        winnings := hand.bid * (index + 1)
        log.Printf("Hand %d: %v Bid: %d Winnings: %v\n", index + 1, hand.cards, hand.bid, winnings)
        total += winnings
    }
    //log.Println("Hand 500: ", hands[876].cards)
    log.Printf("Total: %d\n", total)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)
    partOne(scanner)
}
