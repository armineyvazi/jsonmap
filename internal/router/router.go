package router

import (
	"context"
	"github.com/armineyvazi/jsonmap/internal/http/rest/controller/v1/gpt"
	"github.com/armineyvazi/jsonmap/pkg/framework/ports"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/fx"
)

func Router(
	lc fx.Lifecycle,
	gptHandler gpt.GptHandler,
	httpService ports.Server) {

	httpService.ActiveRecover()
	httpService.ActiveSwagger()
	httpService.POST("api/v1/gpt", gptHandler.ReturnJsonResponse)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := httpService.Listen(); err != nil {
					log.Error(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := httpService.ShutDownWithContext(ctx); err != nil {
				log.Error(err)
			}
			return nil
		},
	})
}
