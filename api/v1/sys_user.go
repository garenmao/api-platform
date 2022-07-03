package v1

import (
	"api-platform/cache"
	"api-platform/global"
	"api-platform/model/common"
	"api-platform/model/domain"
	"api-platform/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type UserApi struct{}

var store = cache.DefaultRedisStore

func (u *UserApi) Register(c *gin.Context) {
	var r request.Register
	err := c.ShouldBindJSON(&r)
	global.Logger.Info("Register info", zap.Any("Body", r))
	if err != nil {
		global.Logger.Error("Register error:%s", zap.Error(err))
		common.FailWithMessage("struct bind error", c)
		return
	}
	// 校验验证码是否正确
	verify := store.UseWithCtx(c).Verify(r.CaptchaId, r.Captcha, false)
	if !verify {
		common.FailWithMessage("captcha error", c)
		return
	}

	now := time.Now()
	user := &domain.User{
		Model: domain.Model{
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		Name:         r.Username,
		Age:          r.Age,
		PassWord:     r.Password,
		Email:        r.Email,
		PhoneNumber:  r.PhoneNumber,
		Gender:       r.Gender,
		Disable:      false,
		RegisterDate: &now,
	}

	userService.CreateUser(user)
	common.OkWithDetailed(nil, "success", c)
}

func (u *UserApi) Login(c *gin.Context) {

}
