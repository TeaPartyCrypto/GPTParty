package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Get the Discord bot token from the environment variable.
	discordToken := os.Getenv("DISCORD_BOT_TOKEN")
	if discordToken == "" {
		fmt.Println("Discord bot token not found. Set DISCORD_BOT_TOKEN environment variable.")
		return
	}

	// Create a new Discord session.
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// Register the messageCreate function as a callback for the MessageCreate event.
	dg.AddHandler(messageCreate)

	// Open the websocket connection to Discord.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	// Wait for a CTRL+C or other termination signal.
	fmt.Println("Bot is running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close the Discord session.
	_ = dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}

	fmt.Printf("Message received: %s", m.Content)

	//

	if isQuestion(m.Content) {
		// if the message is a question, send it to the GPT model
		// and get a response
		response := getGPTResponse(m.Content)
		sResponse := &gPTResponse{}
		err := json.Unmarshal([]byte(response), sResponse)
		if err != nil {
			fmt.Println("Error unmarshalling response:", err)
			return
		}
		// if the response is relevant, send it to the channel
		if isThisPromptRelevant(sResponse.Choices[0].Text) {
			_, _ = s.ChannelMessageSend(m.ChannelID, sResponse.Choices[0].Text)
		} else {
			// if the response is not relevant, ignore it
			return
		}
	} else {
		// if the message is not a question, ignore it
		return
	}

}

func getGPTResponse(m string) string {

	// get the message from the user
	message := m

	fmt.Println("sending message to openai: " + message)

	url := "https://api.openai.com/v1/completions"

	data := requestBody{
		Model:            "text-davinci-003",
		Prompt:           message,
		Temperature:      0,
		MaxTokens:        150,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0.6,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_API_KEY")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
	return string(body)
}

func isThisPromptRelevant(promptResponse string) bool {
	irrelevantPhrases := []string{
		"I'm not sure",
		"I don't know",
		"Sorry, I cannot",
		"Apologies, I cannot",
		"I am unable to",
		"Unable to provide",
		"I do not have enough information",
		"Please provide more information",
		"Could you please clarify",
		"Can you please provide more context",
		"Please clarify",
	}

	lowercaseResponse := strings.ToLower(promptResponse)

	for _, phrase := range irrelevantPhrases {
		if strings.Contains(lowercaseResponse, strings.ToLower(phrase)) {
			return false
		}
	}

	return true
}

func isQuestion(s string) bool {
	questionWords := []string{"who", "what", "when", "where", "why", "how", "which", "whose", "whom", "is", "are", "am", "do", "does", "did", "can", "could", "will", "would", "should", "shall", "might", "may", "ought", "have", "has", "had"}
	containsQuestionWord := false

	// Check if the string contains a question mark.
	if strings.Contains(s, "?") {
		return true
	}

	// Check if the string contains any common question words.
	words := strings.Split(strings.ToLower(s), " ")
	for _, word := range words {
		for _, questionWord := range questionWords {
			if word == questionWord {
				containsQuestionWord = true
				break
			}
		}
		if containsQuestionWord {
			break
		}
	}

	return containsQuestionWord
}

type requestBody struct {
	Model            string   `json:"model"`
	Prompt           string   `json:"prompt"`
	Temperature      float64  `json:"temperature"`
	MaxTokens        int      `json:"max_tokens"`
	TopP             int      `json:"top_p"`
	FrequencyPenalty int      `json:"frequency_penalty"`
	PresencePenalty  float64  `json:"presence_penalty"`
	Stop             []string `json:"stop"`
}

type gPTResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     any    `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
