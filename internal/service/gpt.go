package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/armineyvazi/jsonmap/pkg/framework/ports"
	"go.uber.org/zap"
	"regexp"
	"sync"

	"github.com/armineyvazi/jsonmap/constant"
	"github.com/armineyvazi/jsonmap/dto"
	"github.com/armineyvazi/jsonmap/internal/utils"
	"github.com/sashabaranov/go-openai"
)

type GptService interface {
	GetLaptopDetails(ctx context.Context, requestContent string) ([]dto.LaptopDetail, error)
}

type gptService struct {
	client *openai.Client
	log    ports.Logger
}

func NewGptService(client *openai.Client, log ports.Logger) GptService {
	return &gptService{client: client, log: log}
}

// Simple caching for better performance.
var (
	laptopDetailCache = make(map[string][]dto.LaptopDetail)
	cacheMutex        sync.RWMutex
)

const maxRetryAttempts = 2

func (g *gptService) GetLaptopDetails(ctx context.Context, requestContent string) ([]dto.LaptopDetail, error) {
	// Check if the request data is already in the cache.
	cacheMutex.RLock()
	cachedDetails, foundInCache := laptopDetailCache[requestContent]
	cacheMutex.RUnlock()

	if foundInCache {
		g.log.Error(ctx, "Cache hit for request content:", zap.Any("cachedDetails", cachedDetails))
		return cachedDetails, nil
	}

	// Call the GPT model for a new response with retry logic.
	laptopDetails, err := g.fetchLaptopDetailsWithRetries(ctx, requestContent)
	if err != nil {
		g.log.Error(ctx, "error retries chat gpt", zap.Error(err))
		return nil, err
	}

	// Store the processed data in the cache.
	cacheMutex.Lock()
	laptopDetailCache[requestContent] = laptopDetails
	cacheMutex.Unlock()

	return laptopDetails, nil
}

func (g *gptService) fetchLaptopDetailsWithRetries(ctx context.Context, requestContent string) ([]dto.LaptopDetail, error) {
	var laptopDetails []dto.LaptopDetail
	var err error

	// Retry up to maxRetryAttempts if the laptop details are empty.
	for attempt := 0; attempt < maxRetryAttempts; attempt++ {
		laptopDetails, err = g.fetchLaptopDetails(ctx, requestContent)
		if err != nil {
			return nil, err
		}

		// Check if the response contains valid laptop details.
		if utils.CheckLaptopIsNotEmpty(laptopDetails) {
			return laptopDetails, nil
		}

		g.log.Info(ctx, "Attempt : Empty laptop details, retrying...",
			zap.Int("attempt", attempt+1),
			zap.Int("maxRetryAttempts", maxRetryAttempts))
	}

	// Final attempt after retries.
	laptopDetails, err = g.fetchLaptopDetails(ctx, requestContent)
	if err != nil {
		return nil, err
	}

	// Return the last attempt's result, even if it's empty.
	if !utils.CheckLaptopIsNotEmpty(laptopDetails) {
		g.log.Error(ctx,
			"Final attempt resulted in empty laptop details.",
			zap.Any("laptopDetails", laptopDetails))

		return nil, errors.New("failed to retrieve valid laptop details after retries")
	}

	return laptopDetails, nil
}

func (g *gptService) fetchLaptopDetails(ctx context.Context, requestContent string) ([]dto.LaptopDetail, error) {
	resp, err := g.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: constant.PromptToJSON + constant.BindToStruct + requestContent,
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	// Extract the JSON string using the regex.
	re := regexp.MustCompile(constant.Regex)
	jsonStr := re.FindString(resp.Choices[0].Message.Content)
	if jsonStr == "" {
		g.log.Error(ctx, "failed to extract JSON string from GPT response",
			zap.String("chatGptResponse", resp.Choices[0].Message.Content))
		return nil, errors.New("failed to extract JSON string from GPT response")
	}

	// Unmarshal the JSON data into a slice of LaptopDetail structs.
	var laptopDetails []dto.LaptopDetail
	if err := json.Unmarshal([]byte(jsonStr), &laptopDetails); err != nil {
		return nil, err
	}

	return laptopDetails, nil
}
