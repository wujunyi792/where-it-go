package websocketHandle

import (
	"github.com/gin-gonic/gin"
	"github.com/wujunyi792/gin-template-new/internal/service/websocket"
	"github.com/wujunyi792/gin-template-new/pkg/utils/gen/xrandom"
)

func HandleConnectWebSocket(c *gin.Context) {
	websocket.SocketServer(c.Writer, c.Request, xrandom.GetRandom(7, xrandom.RAND_ALL)) // 授予一个唯一id
}
