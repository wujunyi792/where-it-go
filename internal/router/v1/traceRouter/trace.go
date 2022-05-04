package traceRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/wujunyi792/where-it-go/internal/handle/traceHandle"
)

func InitTraceRouter(e *gin.Engine) {
	trace := e.Group("/trace")
	{
		trace.GET("/cms/:phone", traceHandle.HandleSendSMC)
		trace.GET("/:phone/:token", traceHandle.HandleGetTrace)
	}
}
