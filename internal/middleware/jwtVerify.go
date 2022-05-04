package middleware

import (
	"github.com/gin-gonic/gin"
	err2 "github.com/wujunyi792/gin-template-new/internal/dto/err"
	"github.com/wujunyi792/gin-template-new/internal/service/jwtTokenGen"
)

func JwtVerify(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "" {
		entry, err := jwtTokenGen.ParseToken(token)
		if err == nil {
			//c.Set("token", token)
			c.Set("uid", entry.Info.UID)
			c.Next()
			return
		} else {
			Fail(c, err2.JWTErr)
			return
		}
	}
	Fail(c, err2.VerifyErr)
	return
}
