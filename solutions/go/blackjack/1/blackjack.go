package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace": return 11
        case "two": return 2
        case "three": return 3
        case "four": return 4
        case "five": return 5
        case "six": return 6
        case "seven": return 7
        case "eight": return 8
        case "nine": return 9
        case "ten", "jack", "queen", "king": return 10
		default: return 0
    }
}

const STAND string = "S"
const HIT string = "H"
const SPLIT string = "P"
const WIN string = "W"

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    playerValue := ParseCard(card1) + ParseCard(card2)
    dealerValue := ParseCard(dealerCard)
	switch {
        case card1 == "ace" && card2 == "ace":
        	return SPLIT
        case playerValue == 21 && dealerValue != 10 && dealerCard != "ace":
        	return WIN
    	case 17 <= playerValue && playerValue <= 21:
        	return STAND
        case 12 <= playerValue && playerValue <= 16 && dealerValue >= 7:
        	return HIT
        case 12 <= playerValue && playerValue <= 16:
        	return STAND
        default:
        	return HIT
    }
}
