package blackjack

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
		return 0 // other card
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	dealerCardWinBreaker := dealerCard == "ace" || ParseCard(dealerCard) == 10
	cardSum := ParseCard(card1) + ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case cardSum == 21 && !dealerCardWinBreaker:
		return "W"
	case cardSum == 21 && dealerCardWinBreaker:
		return "S"
	case cardSum >= 17 && cardSum <= 20:
		return "S"
	case cardSum >= 12 && cardSum <= 16 && !(dealerCardValue >= 7):
		return "S"
	case cardSum >= 12 && cardSum <= 16 && (dealerCardValue >= 7):
		return "H"
	case cardSum <= 11:
		return "H"
	default:
		return "" // Unknown action
	}
}
