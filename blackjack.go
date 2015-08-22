package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"bufio"
	"os"
)

type card struct {
	suit string
	value int
	name string
}

func initSuit(suit string) []card {
	var suitArr []card
	for i := 1; i <= 13; i++ {
		var card card
		card.suit = suit
		card.value = i

		if i == 1 {
			card.name = "Ace"
			card.value = 11
		} else if i == 11 {
			card.name = "Jack"
			card.value = 10
		} else if i == 12 {
			card.name = "Queen"
			card.value = 10
		} else if i == 13 {
			card.name = "King"
			card.value = 10
		} else {
			card.name = strconv.Itoa(card.value)
		}

		card.name = card.name + " of " + strings.Title(card.suit)

		suitArr = append(suitArr, card)
	}
	return suitArr
}

func initDeck() []card {
	var deck []card

	clubs := initSuit("clubs")
	hearts := initSuit("hearts")
	diamonds := initSuit("diamonds")
	spades := initSuit("spades")

	deck = append(deck, clubs...)
	deck = append(deck, hearts...)
	deck = append(deck, diamonds...)
	deck = append(deck, spades...)

	shuffle(deck)

	return deck
}

func shuffle(deck []card) {
	rand.Seed( time.Now().UTC().UnixNano())
	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
}

func drawCard(deck []card) card {
	var c card = deck[0]
	copy(deck[0:], deck[1:])

	var emptyCard card
	deck[len(deck) - 1] = emptyCard

	return c
}

func initHand(deck []card) []card {
	var hand []card
	for i := 0; i < 2; i++ {
		hand = append(hand, drawCard(deck))
	}
	return hand
}

func getTotal(hand []card) int {
	var aces int
	var total int
	for i := 0; i < len(hand); i++ {
		total = total + hand[i].value

		if hand[i].value == 11 {
			aces++
		}
	}

	if total > 21 && aces > 0 {
		for i := 0; i < aces; i++ {
			total = total - 10
			if total < 21 {
				break
			}
		}
	}

	return total
}

func printHand(msg string, hand []card, hidden bool) {
	if len(msg) != 0 {
		fmt.Println(msg)
	}

	for i := 0; i < len(hand); i++ {
		if i == 0 && hidden {
			fmt.Println("Hidden Card")
		} else {
			fmt.Println(hand[i].name)
		}
	}

	if !hidden {
		fmt.Println("Total: " + strconv.Itoa(getTotal(hand)))
	}
}

func getInput(msg string, reader *bufio.Reader) string {
	fmt.Print(msg)
	input, _ := reader.ReadString('\n')
	return strings.ToLower(strings.Fields(input)[0])
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input := getInput("Would you like to play blackjack (Y/N)? ", reader)

	replay := true

	for replay {
		if input == "y" {

			deck := initDeck()
			dealersHand := initHand(deck)
			usersHand := initHand(deck)

			printHand("\nDealer's Hand", dealersHand, true)
			printHand("\nYour Hand", usersHand, false)

			gameOver := false

			for !gameOver {
				input = getInput("\nWould you like to (H)it or (S)tand? ", reader)
				validInput := input == "h" || input == "s"

				for(!validInput) {
					input = getInput("Invalid input. Would you like to (H)it or (S)tay? ", reader)
					validInput = input == "h" || input == "s"
				}

				if input == "h" {
					usersHand = append(usersHand, drawCard(deck))

					if getTotal(dealersHand) < 17 {
						fmt.Println("\nDealer Hits")
						dealersHand = append(dealersHand, drawCard(deck))
					}

					printHand("\nDealer's Hand", dealersHand, true)
					printHand("\nYour Hand", usersHand, false)

					if getTotal(usersHand) > 21 {
						fmt.Println("\nBust. You lose.")
						gameOver = true
					}
				} else if input == "s" {
					for getTotal(dealersHand) < 17 {
						fmt.Println("\nDealer Hits")
						dealersHand = append(dealersHand, drawCard(deck))
						printHand("\nDealer's Hand", dealersHand, true)
					}

					fmt.Println("\nDealer Flips Hidden Card")

					printHand("\nDealer's Hand", dealersHand, false)
					printHand("\nYour Hand", usersHand, false)

					if getTotal(usersHand) > 21 && getTotal(dealersHand) > 21 {
						fmt.Println("\nPush. Both players bust.")
					} else if(getTotal(dealersHand) > 21) {
						fmt.Println("\nYou win. Dealer busts.")
					} else if (getTotal(usersHand) > 21) {
						fmt.Println("\nBust. You lose.")
					} else if(getTotal(usersHand) > getTotal(dealersHand)) {
						fmt.Println("\nYou win!")
					} else if(getTotal(usersHand) == getTotal(dealersHand)) {
						fmt.Println("\nPush. Both hands are equal.")
					} else {
						fmt.Println("\nYou lose. Dealer wins.")
					}

					gameOver = true
				}
			}
		}

		input = getInput("\nWould you like to replay (Y/N)? ", reader)

		if input != "y" {
			replay = false
		}

	}
}