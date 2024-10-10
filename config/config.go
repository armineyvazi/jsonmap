package config

type Config struct {
	AppName     string      `mapstructure:"app_name"`
	HTTPPServer HTTPPServer `mapstructure:"http_server"`
	ZapLogLevel string      `mapstructure:"zap_log_level"`
}

type HTTPPServer struct {
	Address string `mapstructure:"address"`
}

func (c Config) GetConfig() Config {
	return c
}
