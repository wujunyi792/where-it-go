package trace

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/parnurzeal/gorequest"
	uuid "github.com/satori/go.uuid"
	"github.com/wujunyi792/where-it-go/config"
	"github.com/wujunyi792/where-it-go/internal/logger"
	"github.com/wujunyi792/where-it-go/internal/service/oss"
	"github.com/wujunyi792/where-it-go/pkg/utils/crypto"
	"strings"
	"time"
)

func SendCmsCode(phone string) (string, error) {
	var req SendCMSRequest
	req.UUID = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	req.Phone = phone
	now := time.Now().Format("2006-01-02 15:04:05")
	req.TimeString = now
	req.TimeMD5 = crypto.Md5Crypto("MOFXTCJq8bOhlSi" + now)

	var res SendCMSResponse
	resp, body, err := gorequest.New().Post("https://xcweb02.caict.ac.cn:8088/dYvFMYL8/h8A2xuUsHKoMz").
		Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36").
		Set("Referer", "https://xc.caict.ac.cn/").
		SendStruct(&req).
		EndStruct(&res)
	if err != nil {
		logger.Error.Println(resp)
		logger.Error.Println(body)
		logger.Error.Println(err)
		return "", err[0]
	}
	logger.Debug.Println(res)
	if res.Code != "00" {
		return "", errors.New(res.ErrorDesc)
	}
	return res.QueryId, nil
}

func GetTrace(phone string, token string, queryId string) (*GetTraceInfoResponse, error) {
	var req GetTraceInfoRequest
	req.Crypto = crypto.Md5Crypto("OcpqZSOIZOxr0" + phone)
	req.Phone = phone
	req.Code = ""
	req.QueryId = queryId
	req.Token = token

	var res GetTraceInfoResponse
	resp, body, err := gorequest.New().Post("https://xcweb02.caict.ac.cn:8088/PMfdZtmQM6PIu/jKDFMlURRsN6D9").
		Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36").
		Set("Referer", "https://xc.caict.ac.cn/").
		SendStruct(&req).
		EndStruct(&res)
	if err != nil {
		logger.Error.Println(resp)
		logger.Error.Println(body)
		logger.Error.Println(err)
		return nil, err[0]
	}

	logger.Debug.Println(res)
	if res.Code != "00" {
		return nil, errors.New(res.ErrorDesc)
	}
	if config.GetConfig().OSS.Use {
		imageData, _ := base64.StdEncoding.DecodeString(res.Result.Message)
		imageReader := bytes.NewReader(imageData)
		url := oss.UploadFileToOss("a.jpg", imageReader)
		res.Result.Message = url
	} else {
		res.Result.MessageBase64 = res.Result.Message
		res.Result.Message = "data:image/png;base64," + res.Result.Message
	}

	return &res, nil
}
