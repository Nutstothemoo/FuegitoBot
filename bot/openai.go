package bot

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)
 
 const COMPLETIONS_URL = "https://api.openai.com/v1/chat/completions"
 
 type CreateCompletionsRequest struct {
	Model            string            `json:"model,omitempty"`
	Messages         []Message         `json:"messages,omitempty"`
	Prompt           []string          `json:"prompt,omitempty"`
	Suffix           string            `json:"suffix,omitempty"`
	MaxTokens        int               `json:"max_tokens,omitempty"`
	Temperature      float64           `json:"temperature,omitempty"`
	TopP             float64           `json:"top_p,omitempty"`
	N                int               `json:"n,omitempty"`
	Stream           bool              `json:"stream,omitempty"`
	LogProbs         int               `json:"logprobs,omitempty"`
	Echo             bool              `json:"echo,omitempty"`
	Stop             []string          `json:"stop,omitempty"`
	PresencePenalty  float64           `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64           `json:"frequency_penalty,omitempty"`
	BestOf           int               `json:"best_of,omitempty"`
	LogitBias        map[string]string `json:"logit_bias,omitempty"`
	User             string            `json:"user,omitempty"`
 }


 type Client struct {
	apiKey       string
	Organization string
 }
 
 // NewClient creates a new client
 func NewClient(apiKey string, organization string) *Client {
	return &Client{
	 apiKey:       apiKey,
	 Organization: organization,
	}
 }
 
 // Post makes a post request
 func (c *Client) Post(url string, input any) (response []byte, err error) {
	response = make([]byte, 0)
 
	rJson, err := json.Marshal(input)
	if err != nil {
	 return response, err
	}
 
	resp, err := c.Call(http.MethodPost, url, bytes.NewReader(rJson))
	if err != nil {
	 return response, err
	}
	defer resp.Body.Close()
 
	response, err = io.ReadAll(resp.Body)
	return response, err
 }
 

 // Call makes a request
 func (c *Client) Call(method string, url string, body io.Reader) (response *http.Response, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
	 return response, err
	}
 
	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	req.Header.Add("Content-Type", "application/json")
	if c.Organization != "" {
	 req.Header.Add("OpenAI-Organization", c.Organization)
	}
 
	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
 }
 
 func (c *Client) CreateCompletionsRaw(r CreateCompletionsRequest) ([]byte, error) {
	return c.Post(COMPLETIONS_URL, r)
 }
 
 func (c *Client) CreateCompletions(r CreateCompletionsRequest) (response CreateCompletionsResponse, err error) {
	raw, err := c.CreateCompletionsRaw(r)
	if err != nil {
	 return response, err
	}
 
	err = json.Unmarshal(raw, &response)
	return response, err
 }
 
 type CreateCompletionsResponse struct {
	ID      string `json:"id,omitempty"`
	Object  string `json:"object,omitempty"`
	Created int    `json:"created,omitempty"`
	Model   string `json:"model,omitempty"`
	Choices []struct {
	 Message struct {
		 Role    string `json:"role,omitempty"`
		 Content string `json:"content,omitempty"`
	 } `json:"message"`
	 Text         string      `json:"text,omitempty"`
	 Index        int         `json:"index,omitempty"`
	 Logprobs     interface{} `json:"logprobs,omitempty"`
	 FinishReason string      `json:"finish_reason,omitempty"`
	} `json:"choices,omitempty"`
	Usage struct {
	 PromptTokens     int `json:"prompt_tokens,omitempty"`
	 CompletionTokens int `json:"completion_tokens,omitempty"`
	 TotalTokens      int `json:"total_tokens,omitempty"`
	} `json:"usage,omitempty"`
 
	Error Error `json:"error,omitempty"`
 }
 
 // Error is the error standard response from the API
 type Error struct {
	Message string      `json:"message,omitempty"`
	Type    string      `json:"type,omitempty"`
	Param   interface{} `json:"param,omitempty"`
	Code    interface{} `json:"code,omitempty"`
 }
 
 type Message struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
 }