package di

import (
	"github.com/armineyvazi/jsonmap/config"
	"github.com/armineyvazi/jsonmap/constant"
	"github.com/armineyvazi/jsonmap/internal/http/rest/controller/v1/gpt"
	"github.com/armineyvazi/jsonmap/internal/router"
	"github.com/armineyvazi/jsonmap/internal/service"
	"github.com/armineyvazi/jsonmap/pkg/framework/adapters"
	"github.com/armineyvazi/jsonmap/pkg/framework/ports"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/fx"
	"os"
	"path"
)

func InitializeApp() error {
	app := fx.New(
		fx.Provide(
			NewConfig,
			NewLogger,
			NewHTTPServer,
			NewGptService,
			gpt.NewGptHandler,
		),
		fx.Invoke(router.Router),
	)

	if err := app.Err(); err != nil {
		return err
	}

	app.Run()

	return nil
}

func NewLogger(config ports.Config[config.Config]) ports.Logger {
	return adapters.NewLogger(config.GetConfig().ZapLogLevel, false)
}

func NewConfig() (ports.Config[config.Config], error) {
	var configs config.Config
	pwd, err := os.Getwd()
	if err != nil {
		return configs.GetConfig(), err
	}
	configPath := path.Join(pwd, constant.DirPath)
	err = adapters.NewViper(&configs, configPath)
	if err != nil {
		return configs.GetConfig(), err
	}

	return configs.GetConfig(), nil
}

func NewHTTPServer(config ports.Config[config.Config]) ports.Server {
	return adapters.NewFiber(config.GetConfig().AppName, config.GetConfig().HTTPPServer.Address)
}

func NewGptService(log ports.Logger) service.GptService {
	return service.NewGptService(openai.NewClient(constant.GptToken), log)
}
