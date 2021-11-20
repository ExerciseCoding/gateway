package middleware

import (
	"errors"
	"fmt"
	"gateway/public"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		fmt.Println("session信息",session.Get(public.AdminSessionInfoKey), c.Request)
		if adminInfo,ok:=session.Get(public.AdminSessionInfoKey).(string);!ok||adminInfo==""{
			ResponseError(c, InternalErrorCode, errors.New("user not login"))
			c.Abort()
			return
		}
		c.Next()
	}
}
