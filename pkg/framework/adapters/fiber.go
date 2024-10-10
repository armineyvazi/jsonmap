package adapters

import (
	"context"
	"github.com/armineyvazi/jsonmap/pkg/framework/ports"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	l "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"go.elastic.co/apm/module/apmfiber"
	"net/http"
)

type FiberHttpServer struct {
	app     *fiber.App
	debug   bool
	address string
}

func NewFiber(appName string, address string) ports.Server {

	return &FiberHttpServer{
		app: fiber.New(fiber.Config{
			AppName:                      appName,
			StreamRequestBody:            false,
			DisablePreParseMultipartForm: false,
			ReduceMemoryUsage:            false,
			JSONEncoder:                  json.Marshal,
			JSONDecoder:                  json.Unmarshal,
			XMLEncoder:                   nil,
			Network:                      "",
			EnableTrustedProxyCheck:      false,
			TrustedProxies:               nil,
			EnableIPValidation:           false,
			EnablePrintRoutes:            false,
			ColorScheme:                  fiber.Colors{},
			RequestMethods:               nil,
			EnableSplittingOnParsers:     false,
		}),
		debug:   false,
		address: address,
	}
}

func (fh *FiberHttpServer) ActiveRecover() {
	fh.app.Use(recover.New())
}

func (fh *FiberHttpServer) ActiveApm() {
	fh.app.Use(apmfiber.Middleware())
}

func (fh *FiberHttpServer) ActiveLogger() {
	fh.app.Use(l.New())
}

func (fh *FiberHttpServer) Use(args ...interface{}) {
	fh.app.Use(args...)
}

func (fh *FiberHttpServer) ActiveSwagger() {
	fh.app.Get("/swagger/*", swagger.HandlerDefault)
}

func (fh *FiberHttpServer) GET(path string, handlers ...fiber.Handler) {
	fh.app.Get(path, handlers...)
}

func (fh *FiberHttpServer) POST(path string, handlers ...fiber.Handler) {
	fh.app.Post(path, handlers...)
}

func (fh *FiberHttpServer) PUT(path string, handlers ...fiber.Handler) {
	fh.app.Put(path, handlers...)
}

func (fh *FiberHttpServer) DELETE(path string, handlers ...fiber.Handler) {
	fh.app.Delete(path, handlers...)
}

func (fh *FiberHttpServer) OPTIONS(path string, handlers ...fiber.Handler) {
	fh.app.Options(path, handlers...)
}

func (fh *FiberHttpServer) PATCH(path string, handlers ...fiber.Handler) {
	fh.app.Patch(path, handlers...)
}

func (fh *FiberHttpServer) Listen() error {
	if err := fh.app.Listen(fh.address); err != nil {
		return err
	}
	return nil
}

// ShutDownWithContext gracefully shuts down the http server
func (fh *FiberHttpServer) ShutDownWithContext(ctx context.Context) error {
	return fh.app.ShutdownWithContext(ctx)
}

func (fh *FiberHttpServer) Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	return fh.app.Test(req, msTimeout...)
}
