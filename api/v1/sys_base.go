package v1

import (
	"api-platform/cache"
	"api-platform/global"
	"api-platform/model/common"
	"api-platform/model/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var Store = cache.DefaultRedisStore

type BaseApi struct{}

const SourceString = "1234567890QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm"

func (b *BaseApi) Captcha(c *gin.Context) {
	// 验证码配置
	config := global.Config.Captcha
	driver := base64Captcha.NewDriverString(config.Height,
		config.Width, config.NoiseCount,
		config.ShowLineOptions, config.Length, SourceString,
		&color.RGBA{R: 254, G: 254, B: 254, A: 254},
		nil, []string{"Flim-Flam.ttf", "wqy-microhei.ttc"})
	cp := base64Captcha.NewCaptcha(driver, Store.UseWithCtx(c))
	if id, b64s, err := cp.Generate(); err != nil {
		common.FailWithMessage("验证码获取失败", c)
	} else {
		common.OkWithData(response.CaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: config.Length,
		}, c)
	}
}
