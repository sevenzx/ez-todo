package config

type _Config struct {
	Server     Server     `mapstructure:"server" json:"server" yaml:"server"`
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	JWT        JWT        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Lumberjack Lumberjack `mapstructure:"lumberjack" json:"lumberjack" yaml:"lumberjack"`
}

var Config = new(_Config)
