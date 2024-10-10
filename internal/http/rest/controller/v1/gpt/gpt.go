package gpt

import (
	"github.com/armineyvazi/jsonmap/internal/service"
	"github.com/armineyvazi/jsonmap/pkg/framework/ports"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// ErrorResponse represents the structure for error messages.
type ErrorResponse struct {
	Error string `json:"error"`
}

// GptHandler defines the interface for handling GPT-related HTTP requests.
type GptHandler interface {
	ReturnJsonResponse(ctx *fiber.Ctx) error
}

// gptHandler implements the GptHandler interface.
type gptHandler struct {
	laptopService service.GptService
	log           ports.Logger
}

// NewGptHandler creates a new instance of gptHandler.
func NewGptHandler(laptopService service.GptService, log ports.Logger) GptHandler {
	return &gptHandler{
		laptopService: laptopService,
		log:           log,
	}
}

// LaptopRequest represents the request structure for getting laptop details.
type LaptopRequest struct {
	Data string `json:"data"`
}

// ReturnJsonResponse handles the POST request to get laptop details.
//
// @Summary Get laptop details from GPT model
// @Description Processes a request and returns laptop details based on the input data.
// @Description Example curl command:
// @Description ```bash
// @Description curl -X POST "http://localhost:3000/api/v1/gpt" -H "Content-Type: application/json" -d '{"data": "example request data"}'
// @Description ```
// @Tags Laptop
// @Accept json
// @Produce json
// @Param request body LaptopRequest true "Request body containing data"
// @Success 200 {object} []dto.LaptopDetail
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/gpt [post]
func (g *gptHandler) ReturnJsonResponse(ctx *fiber.Ctx) error {
	var request LaptopRequest

	// Bind the JSON body to the LaptopRequest struct.
	if err := ctx.BodyParser(&request); err != nil {
		g.log.Error(ctx.Context(), "error from request", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Failed to parse JSON",
		})
	}

	// Fetch laptop details using the GptService.
	laptopDetails, err := g.laptopService.GetLaptopDetails(ctx.Context(), request.Data)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	g.log.Info(ctx.Context(), "request successfully",
		zap.Any("request", request),
		zap.Any("response", laptopDetails))

	return ctx.JSON(laptopDetails)
}
