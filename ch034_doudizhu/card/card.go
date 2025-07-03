package card

import (
	"math/rand"
	"time"
)

// Card 扑克牌
type Card struct {
	Color string // 花色，如"♠", "♥", "♣", "♦"
	Val   string // 牌面值，如"A", "2", "3"等
}

var (
	bigJoker    = Card{Color: "", Val: "bigJoker"}
	littleJoker = Card{Color: "", Val: "littleJoker"}
)

// NewCards 生成一副完整的扑克牌
func NewCards() []Card {
	colors := []string{"♠", "♥", "♣", "♦"}
	vals := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	cards := make([]Card, 0, 54)

	for _, val := range vals {
		for _, color := range colors {
			cards = append(cards, Card{Color: color, Val: val})
		}
	}
	cards = append(cards, bigJoker, littleJoker)
	return cards
}

// Shuffle 洗牌函数，打乱牌组的顺序
func Shuffle(cards []Card) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}
