package main

import (
	"doudizhu/card"
	"doudizhu/game"
)

func main() {
	// 创建一副扑克牌
	cards := card.NewCards()
	// 洗牌
	card.Shuffle(cards)
	// 发牌
	players := game.FaPai(cards)

	game.StartGame(players)
}
