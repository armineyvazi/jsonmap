package adapters

import (
	"fmt"
	"github.com/armineyvazi/jsonmap/pkg/framework/ports"
	"github.com/spf13/viper"
)

func NewViper[T ports.Config[T]](config *T, configAddress string) error {
	viper.SetConfigFile(configAddress)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}
	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("error unmarshalling configuration: %w", err)
	}
	return nil
}
