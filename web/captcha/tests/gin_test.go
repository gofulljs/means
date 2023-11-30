package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/gofulljs/means/web/captcha"
	"testing"
)

func TestCaptchaGin(t *testing.T) {
	r := gin.Default()

	mgr := captcha.NewMgr(captcha.WithPrefix("/"))
	g := r.Group("/captcha")
	{
		mRouter := mgr.GetRouters()
		for k, v := range mRouter {
			handerFunc := v.ServeHTTP
			g.GET(k, func(ctx *gin.Context) {
				handerFunc(ctx.Writer, ctx.Request)
			})
		}
	}

	r.Run()
}
