package blackjack

const Stand string = "S"
const Hit string = "H"
const Split string = "P"
const Win string = "W"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	CardVal1 := ParseCard(card1)
	CardVal2 := ParseCard(card2)
	CardSum := CardVal1 + CardVal2
	DealerCardVal := ParseCard(dealerCard)
	switch {
	case CardVal1 == 11 && CardVal2 == 11:
		return Split
	case CardSum <= 20 && CardSum >= 17:
		return Stand
	case CardSum == 21 && (DealerCardVal == 11 || DealerCardVal == 10):
		return Stand
	case CardSum == 21 && (DealerCardVal != 11 && DealerCardVal != 10):
		return Win
	case CardSum <= 17 && CardSum >= 12 && DealerCardVal >= 7:
		return Hit
	case CardSum <= 17 && CardSum >= 12 && DealerCardVal < 7:
		return Stand
	case CardSum <= 11:
		return Hit
	default:
		return "Unknown combindation"
	}
}
