package config

type _Config struct {
	Server     Server     `mapstructure:"server" json:"server" yaml:"server"`
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Lumberjack Lumberjack `mapstructure:"lumberjack" json:"lumberjack" yaml:"lumberjack"`
}

var Config = new(_Config)
