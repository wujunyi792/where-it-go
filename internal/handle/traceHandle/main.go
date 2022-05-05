package traceHandle

import (
	"github.com/gin-gonic/gin"
	trace2 "github.com/wujunyi792/where-it-go/internal/dto/trace"
	"github.com/wujunyi792/where-it-go/internal/middleware"
	"github.com/wujunyi792/where-it-go/internal/redis"
	"github.com/wujunyi792/where-it-go/internal/service/trace"
	"github.com/wujunyi792/where-it-go/pkg/utils/check"
	"time"
)

var (
	RedisPrefix = "where-it-go:"
)

func HandleSendSMC(c *gin.Context) {
	phone := c.Param("phone")
	_, err := redis.GetRedis().Get(RedisPrefix + phone)
	if err == nil {
		middleware.FailWithCode(c, 40201, "查询过于频繁，请稍后再试")
		return
	}
	if !check.VerifyMobileFormat(phone) {
		middleware.FailWithCode(c, 40205, "手机号格式不正确")
		return
	}
	traceId, err := trace.SendCmsCode(phone)
	if err != nil {
		middleware.FailWithCode(c, 40202, err.Error())
		return
	}
	redis.GetRedis().Set(RedisPrefix+phone, traceId, 3*time.Minute)
	middleware.Success(c, &trace2.ServiceSendCMSResponse{
		TraceId: traceId,
		Phone:   phone,
	})
}

func HandleGetTrace(c *gin.Context) {
	phone := c.Param("phone")
	token := c.Param("token")
	traceId, err := redis.GetRedis().Get(RedisPrefix + phone)
	if err != nil {
		// TODO 应该一段时间还是能用的，没测过
		redis.GetRedis().RemoveKey(RedisPrefix+phone, false)
		middleware.FailWithCode(c, 40203, "验证码已过期或不存在")
		return
	}
	traceResult, err := trace.GetTrace(phone, token, traceId)
	if err != nil {
		middleware.FailWithCode(c, 40204, err.Error())
		return
	}
	_ = redis.GetRedis().RemoveKey(RedisPrefix+phone, false)
	middleware.Success(c, traceResult.Result)
}
