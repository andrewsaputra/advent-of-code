package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	result := solve("inputs.txt")
	fmt.Println("Result:", result)

	duration := time.Now().UnixMilli() - startTime
	fmt.Printf("Duration: %vms\n", duration)
}

func solve(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		hands = append(hands, parseInput(line))
	}

	sort.Slice(hands, func(i, j int) bool {
		return comparePower(hands[i], hands[j])
	})

	total := 0
	for rank := 1; rank <= len(hands); rank++ {
		total += rank * hands[rank-1].Bid
	}

	return total
}

type Combo int

const (
	FIVE_OF_A_KIND  Combo = 7
	FOUR_OF_A_KIND  Combo = 6
	FULL_HOUSE      Combo = 5
	THREE_OF_A_KIND Combo = 4
	TWO_PAIR        Combo = 3
	ONE_PAIR        Combo = 2
	HIGH_CARD       Combo = 1
)

type Hand struct {
	Card  string
	Bid   int
	Combo Combo
}

func parseInput(line string) Hand {
	strs := strings.Split(line, " ")
	card := strs[0]
	bid, err := strconv.Atoi(strs[1])
	if err != nil {
		log.Panic(err)
	}

	m := map[rune]int{}
	for _, v := range card {
		m[v]++
	}

	first, firstAmount := getLargest(m, 0)
	_, secondAmount := getLargest(m, first)

	combo := HIGH_CARD
	switch firstAmount {
	case 5:
		combo = FIVE_OF_A_KIND
	case 4:
		combo = FOUR_OF_A_KIND
	case 3:
		combo = THREE_OF_A_KIND
		if secondAmount == 2 {
			combo = FULL_HOUSE
		}
	case 2:
		combo = ONE_PAIR
		if secondAmount == 2 {
			combo = TWO_PAIR
		}
	}

	return Hand{
		Card:  card,
		Bid:   bid,
		Combo: combo,
	}
}

func getLargest(m map[rune]int, exclude rune) (rune, int) {
	maxAmount := 0
	var largest rune

	for k, v := range m {
		if k == exclude {
			continue
		}

		if v > maxAmount {
			maxAmount = v
			largest = k
		}
	}

	return largest, maxAmount
}

func comparePower(hand1 Hand, hand2 Hand) bool {
	findPos := func(c rune) int {
		for i, v := range "23456789TJQKA" {
			if c == v {
				return i
			}
		}

		return -1
	}

	if hand1.Combo == hand2.Combo {
		for i, v1 := range hand1.Card {
			v2 := rune(hand2.Card[i])
			if v1 == v2 {
				continue
			}

			return findPos(v1) < findPos(v2)
		}

		return true
	} else {
		return hand1.Combo < hand2.Combo
	}
}
