package gpt4

type Response struct {
	Text     string
	Relevant bool
}

type GPT4Client struct {
	apiKey string
}

func NewClient(apiKey string) *GPT4Client {
	return &GPT4Client{apiKey: apiKey}
}

func (c *GPT4Client) GetGPT4Response(input string) (*Response, error) {
	// Call the GPT-4 API with the input text, and process the API response.

	// You need to define your logic for determining if the response is relevant or not.
	// This could be based on confidence score or other factors depending on the GPT model and API you are using.
	isRelevant := determineRelevance(apiResponse)

	return &Response{
		Text:     apiResponse.Text,
		Relevant: isRelevant,
	}, nil
}

func determineRelevance(apiResponse APIResponse) bool {
	// Implement your logic to determine if the GPT model response is relevant.
	// This could be based on confidence score, the presence of certain keywords, or any other factors.

	// For example, if your API response contains a confidence score:
	if apiResponse.ConfidenceScore > 0.8 {
		return true
	}

	return false
}
