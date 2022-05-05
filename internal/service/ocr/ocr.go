package ocr

import (
	"github.com/parnurzeal/gorequest"
	"github.com/wujunyi792/where-it-go/config"
	"github.com/wujunyi792/where-it-go/internal/logger"
	"strings"
)

func OCR(base64String string) (text string) {
	var req GetOcrRequest
	var res GetOcrResponse
	req.Images = []string{base64String}
	resp, body, err := gorequest.New().Post(config.GetConfig().OCR.Url).
		SendStruct(&req).EndStruct(&res)
	if err != nil {
		logger.Error.Println(resp)
		logger.Error.Println(body)
		logger.Error.Println(err)
		return ""
	}
	if res.Status != "000" {
		logger.Error.Println("ocr检测失败: ", res.Msg)
		return ""
	}
	for i := 0; i < len(res.Results); i++ {
		for j := 0; j < len(res.Results[i]); j++ {
			text += res.Results[i][j].Text
		}
	}
	strings.ReplaceAll(text, " ", "")
	return
}
