package corn

import (
	"github.com/robfig/cron"
	"github.com/wujunyi792/gin-template-new/internal/logger"
)

func init() {
	c := cron.New()
	err := c.AddFunc("0 0/10 * * * *", func() {})
	if err != nil {
		logger.Error.Fatalln(err)
	}
	c.Start()
	logger.Info.Println("corn init SUCCESS ")
}
