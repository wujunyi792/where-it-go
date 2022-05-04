package mail

import (
	"fmt"
	"github.com/wujunyi792/gin-template-new/config"
	"github.com/wujunyi792/gin-template-new/internal/logger"
	"gopkg.in/gomail.v2"
	"time"
)

var mailConfig struct {
	Instance *gomail.Dialer
}

func init() {
	if !config.GetConfig().Mail.Use {
		panic("mail not open, please check config")
	}

	conf := &config.GetConfig().Mail.Config

	mailConfig.Instance = gomail.NewDialer(
		conf.SMTP,
		conf.PORT,
		conf.ACCOUNT,
		conf.PASSWORD,
	)

	err := SendMail(conf.ACCOUNT, config.GetConfig().ProgramName+` Golang Program init`, fmt.Sprintf("Name: %s\nVERSION: %s\nAuthor: %s\nTime: %s", config.GetConfig().ProgramName, config.GetConfig().VERSION, config.GetConfig().AUTHOR, time.Now().Format("2006-01-02 15:04:05")))

	if err != nil {
		logger.Error.Fatalln(err)
	}
	logger.Info.Println("mail init SUCCESS ")
}

func SendMail(to, title, content string) error {
	conf := &config.GetConfig().Mail.Config
	m := gomail.NewMessage()
	m.SetHeader("From", "Golang Program Manager"+"<"+conf.ACCOUNT+">")
	m.SetHeader("To", to)
	m.SetHeader("Subject", title)
	m.SetBody("text/plain", content)

	return mailConfig.Instance.DialAndSend(m)
}
