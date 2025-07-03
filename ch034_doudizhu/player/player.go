package player

import "doudizhu/card"

// Player 玩家
type Player struct {
	Name    string
	Hand    []card.Card // 手牌
	IsDiZhu bool        // 是否地主
}
