package websocketRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/wujunyi792/where-it-go/internal/handle/websocketHandle"
)

func InitWebSocketRouter(e *gin.Engine) {
	websocket := e.Group("/websocket")
	{
		websocket.GET("/connect", websocketHandle.HandleConnectWebSocket)
	}
}
