package turns

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardOne := ParseCard(card1)
	cardTwo := ParseCard(card2)
	dealer := ParseCard(dealerCard)
	sum := cardOne + cardTwo

	switch {
	case sum == 22:
		return "P"
	case sum == 21:
		if dealer >= 10 {
			return "S"
		}
		return "W"

	case sum >= 17 && sum <= 20:
		return "S"
	case sum >= 12 && sum <= 16:
		if dealer < 7 {
			return "S"
		}
		return "H"

	default:
		return "H"
	}
}
