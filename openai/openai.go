package openai

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var apiKey string;

func SetAPIKey(key string) {
	apiKey = key
}

func NewRequest(model string, messages []*Message) *request {
	return &request{
		Model:    model,
		Messages: messages,
	}
}

func NewMessage(role string, content string) *Message {
	return &Message{
		Role:    role,
		Content: content,
	}
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type request struct {
	Model    string      `json:"model"`
	Messages []*Message `json:"messages"`
}

type response struct {
	ID             string    `json:"id"`
	Object         string    `json:"object"`
	Created        int       `json:"created"`
	Model          string    `json:"model"`
	Choices        []choices `json:"choices"`
	Usege          usage     `json:"usage"`
	SysFingerprint string    `json:"system_fingerprint"`
}

type choices struct {
	Index        int      `json:"index"`
	Message      Message `json:"message"`
	Logprobs     string   `json:"logprobs"`
	FinishReason string   `json:"finish_reason"`
}

type usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func (r *request) Call() response {
	pl, _ := json.Marshal(r)
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(pl))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	ctx, _ := ioutil.ReadAll(resp.Body)

	var data response
	json.Unmarshal(ctx, &data)

	return data
}