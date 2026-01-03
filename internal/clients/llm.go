package clients

import (
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/packages/param"
	"github.com/openai/openai-go/v3/responses"
	"github.com/openai/openai-go/v3/shared/constant"
)

type LLM interface {
	NewChatWithFile(ctx context.Context, question, fileData, filename string) (string, error)
}

type openAi struct {
	client openai.Client
}

func NewLLM() LLM {

	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
	)

	return &openAi{
		client: client,
	}
}

// NewChatWithFile creates a chat with system prompt, user question, and optional PDF file input
// fileData should be base64-encoded PDF file content (only PDF files are supported)
// filename should be the PDF filename (e.g., "resume.pdf")
// Returns the response text and an error if the request fails
func (llm *openAi) NewChatWithFile(ctx context.Context, url, fileData, filename string) (string, error) {
	// Build user message content - can include both text and file
	userContent := responses.ResponseInputMessageContentListParam{}

	// Add file if provided
	if fileData != "" || filename != "" {
		// Validate filename is PDF if provided
		if filename != "" && !strings.HasSuffix(strings.ToLower(filename), ".pdf") {
			return "", fmt.Errorf("invalid file type: only PDF files are supported")
		}

		fileParam := responses.ResponseInputFileParam{
			Type: constant.InputFile("input_file"),
		}

		// Format file_data as data URL if it's raw base64
		// OpenAI expects: data:application/pdf;base64,<base64-data>
		if fileData != "" {
			// Check if it's already a data URL
			if !strings.HasPrefix(fileData, "data:") {
				// Only support PDF files
				fileData = fmt.Sprintf("data:application/pdf;base64,%s", fileData)
			}
			fileParam.FileData = param.NewOpt(fileData)
		}
		if filename != "" {
			fileParam.Filename = param.NewOpt(filename)
		}
		userContent = append(userContent, responses.ResponseInputContentUnionParam{
			OfInputFile: &fileParam,
		})
	}

	// Add text question
	if url != "" {
		userContent = append(userContent, responses.ResponseInputContentParamOfInputText(url))
	}

	// Build input items
	inputItems := responses.ResponseInputParam{}

	// Define JSON schema for structured output
	jsonSchema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"jd": map[string]any{
				"type":        "string",
				"description": "Concise summary of the job description. Use 'Unable to retrieve JD' if the URL cannot be accessed.",
			},
			"elevator_pitch": map[string]any{
				"type":        "string",
				"description": "A compelling 2-3 sentence elevator pitch tailored to the job description and resume.",
			},
			"questions": map[string]any{
				"type":        "array",
				"description": "Exactly 3 meaningful questions to ask the hiring manager or recruiter",
				"items": map[string]any{
					"type": "string",
				},
				"minItems": 3,
				"maxItems": 3,
			},
		},
		"required":             []string{"jd", "elevator_pitch", "questions"},
		"additionalProperties": false,
	}

	// Add system prompt
	inputItems = append(inputItems, responses.ResponseInputItemParamOfMessage(
		`You are a job copilot assistant. Your task is to analyze the provided resume and job description URL to create a personalized elevator pitch and interview questions.

Instructions:
1. Visit the provided URL to access the job description. If the URL is inaccessible or requires login, use 'null' for the jd field.
2. Analyze the resume (PDF) to understand the candidate's background, skills, and experience.
3. Create a compelling elevator pitch (2-3 sentences) that highlights the candidate's most relevant qualifications for this specific role.
4. Generate exactly 3 thoughtful questions that demonstrate the candidate's interest and help them evaluate if the role is a good fit.

Focus on making the elevator pitch specific to the job requirements and the candidate's unique strengths.`,
		"system",
	))

	// Add user message with content list if we have file, otherwise use simple string
	if len(userContent) > 0 {
		inputItems = append(inputItems, responses.ResponseInputItemParamOfMessage(
			userContent,
			"user",
		))
	} else if url != "" {
		inputItems = append(inputItems, responses.ResponseInputItemParamOfMessage(
			url,
			"user",
		))
	}

	resp, err := llm.client.Responses.New(ctx, responses.ResponseNewParams{
		Model: "gpt-4o-mini", // Fast, cost-effective model that supports PDF input, web search, and structured outputs
		Input: responses.ResponseNewParamsInputUnion{
			OfInputItemList: inputItems,
		},
		Text: responses.ResponseTextConfigParam{
			Format: responses.ResponseFormatTextConfigParamOfJSONSchema("job_copilot_response", jsonSchema),
		},
		Tools: []responses.ToolUnionParam{
			{
				OfWebSearch: &responses.WebSearchToolParam{
					Type: responses.WebSearchToolTypeWebSearch, // Use the stable web search tool type
				},
			},
		},
	})
	if err != nil {
		// Return error instead of panicking - allows graceful error handling
		return "", fmt.Errorf("openai api error: %w", err)
	}

	return resp.OutputText(), nil
}
