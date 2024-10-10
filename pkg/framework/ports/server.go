package ports

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Server interface {
	Listen() error
	Test(req *http.Request, msTimeout ...int) (*http.Response, error)
	ActiveRecover()
	GET(path string, handlers ...fiber.Handler)
	POST(path string, handlers ...fiber.Handler)
	PUT(path string, handlers ...fiber.Handler)
	DELETE(path string, handlers ...fiber.Handler)
	OPTIONS(path string, handlers ...fiber.Handler)
	PATCH(path string, handlers ...fiber.Handler)
	ShutDownWithContext(ctx context.Context) error
	Use(args ...interface{})
	ActiveSwagger()
}
