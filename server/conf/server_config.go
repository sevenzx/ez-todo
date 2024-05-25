package conf

type Server struct {
	HostPort   string `mapstructure:"host-port" json:"host-port" yaml:"host-port"`
	BaseRouter string `mapstructure:"base-router" json:"base-router" yaml:"base-router"`
}
