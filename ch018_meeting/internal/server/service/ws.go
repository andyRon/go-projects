package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var wsP2PConnMap = sync.Map{}

func Wsp2PConnection(c *gin.Context) {
	// 1 获取房间和用户的信息
	in := new(WsP2PConnectionRequest)
	err := c.ShouldBindUri(in)
	if err != nil {
		log.Println("ShouldBindUri err:", err)
		return
	}
	// 2 升级协议
	var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	// 3 存储当前的链接信息
	userConnMap := new(sync.Map)
	value, ok := wsP2PConnMap.Load(in.RoomIdentity)
	if ok {
		userConnMap = value.(*sync.Map)
	}
	userConnMap.Store(in.UserIdentity, conn)
	wsP2PConnMap.Store(in.RoomIdentity, userConnMap)
	// 4 监听发过来的消息
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			return
		}
		value_, ok_ := wsP2PConnMap.Load(in.RoomIdentity)
		if ok_ {
			value_.(*sync.Map).Range(func(key, value any) bool {
				err := value.(*websocket.Conn).WriteMessage(websocket.TextMessage, data)
				if err != nil {
					log.Println("WriteMessage err:", err)
				}
				return true
			})
		}
	}
}
