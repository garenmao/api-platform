package config

type Captcha struct {
	Height          int `yaml:"height"`
	Width           int `yaml:"width"`
	NoiseCount      int `yaml:"noiseCount"`
	ShowLineOptions int `yaml:"showLineOptions"`
	Length          int `yaml:"length"`
}
