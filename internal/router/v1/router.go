package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wujunyi792/where-it-go/config"
	"github.com/wujunyi792/where-it-go/internal/middleware"
	"github.com/wujunyi792/where-it-go/internal/router/v1/traceRouter"
	"github.com/wujunyi792/where-it-go/internal/router/v1/websocketRouter"
)

func MainRouter(e *gin.Engine) {
	e.Any("", func(c *gin.Context) {
		data := struct {
			UA         string
			Host       string
			Method     string
			Proto      string
			RemoteAddr string
			Message    string
		}{
			UA:         c.Request.Header.Get("User-Agent"),
			Host:       c.Request.Host,
			Method:     c.Request.Method,
			Proto:      c.Request.Proto,
			RemoteAddr: c.Request.RemoteAddr,
			Message:    fmt.Sprintf("Welcome to %s, version %s.", config.GetConfig().ProgramName, config.GetConfig().VERSION),
		}
		middleware.Success(c, data)
	})
	websocketRouter.InitWebSocketRouter(e)
	traceRouter.InitTraceRouter(e)
}
