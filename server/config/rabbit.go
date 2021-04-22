package config

type RabbitMq struct {
	Admin   string `mapstructure:"admin" json:"admin" yaml:"admin"`
	Pwd     string `mapstructure:"pwd" json:"pwd" yaml:"pwd"`
	Port    int    `mapstructure:"port" json:"port" yaml:"port"`
	Ip      string `mapstructure:"ip" json:"ip" yaml:"ip"`
	Verhost string `mapstructure:"verhost" json:"verhost" yaml:"verhost"`
}
