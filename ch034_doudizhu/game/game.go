package game

import (
	"doudizhu/card"
	"doudizhu/player"
	"fmt"
)

// FaPai 发牌函数，将一副牌分发给三个玩家
func FaPai(cards []card.Card) []*player.Player {
	players := []*player.Player{
		{Name: "Player 1"},
		{Name: "Player 2"},
		{Name: "Player 3"},
	}
	for i, card := range cards {
		players[i%3].Hand = append(players[i%3].Hand, card)
	}
	return players
}

// IsValidCombination 判断玩家出牌的牌型是否合法
func IsValidCombination(cards []card.Card) bool {
	// TODO
	// 这里需要根据斗地主的牌型规则进行复杂的判断
	// 例如：单张、对子、顺子、炸弹等
	return true // 示例代码，实际需要详细实现
}

// 游戏的主流程控制，包括叫地主、出牌、判定胜负等。
func StartGame(players []*player.Player) {
	// 叫地主逻辑
	diZhu := players[0] // 假设第一个玩家为地主
	diZhu.IsDiZhu = true

	// 游戏主循环
	for {
		for _, p := range players {
			fmt.Printf("%s 的手牌: %v\n", p.Name, p.Hand)
			// TODO
			// 这里需要实现玩家出牌逻辑
			// 如果玩家出完牌，则结束游戏
		}
	}

	// 判定胜负
	fmt.Println("游戏结束，胜利者为: ", diZhu.Name)
}
