package request

type Register struct {
	Username    string `json:"userName"`
	Password    string `json:"passWord"`
	Age         int    `json:"age"`
	Gender      bool   `json:"gender"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	CaptchaId   string `json:"captchaId"`
	Captcha     string `json:"captcha"`
}
