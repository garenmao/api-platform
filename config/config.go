package config

type Config struct {
	Server  Server  `yaml:"server"`
	Redis   Redis   `yaml:"redis"`
	Captcha Captcha `yaml:"captcha"`
	Mysql   Mysql   `yaml:"mysql"`
	Zap     Zap     `yaml:"zap"`
}
