package config

type Captcha struct {
	KeyLong   int `mapstructure:"key_long"`
	ImgWidth  int `mapstructure:"img_width"`
	ImgHeight int `mapstructure:"img_height"`
}
