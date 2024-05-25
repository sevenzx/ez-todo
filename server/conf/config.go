package conf

type _Config struct {
	Server Server `mapstructure:"server" json:"server" yaml:"server"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

var Config = new(_Config)
