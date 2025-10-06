package openai

import (
	"encoding/json"
	"fmt"
	"io"
	"krillin-ai/config"
	"krillin-ai/log"
	"net/http"
	"os"
	"strings"

	"go.uber.org/zap"
)

func (c *Client) ChatCompletion(query string) (string, error) {
	baseUrl := config.Conf.Llm.BaseUrl
	if baseUrl == "" {
		baseUrl = "https://api.openai.com/v1"
	}
	url := baseUrl + "/chat/completions"

	requestBody := map[string]interface{}{
		"model": config.Conf.Llm.Model,
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "You are an assistant that helps with subtitle translation.",
			},
			{
				"role":    "user",
				"content": query,
			},
		},
		"temperature": 0.9,
		"max_tokens":  4096,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	curlCmd := fmt.Sprintf("curl -X POST '%s' \\\n  -H 'Content-Type: application/json' \\\n  -H 'Authorization: Bearer %s' \\\n  -d '%s'",
		url,
		config.Conf.Llm.ApiKey,
		string(jsonData))
	log.GetLogger().Info("curl command", zap.String("curl", curlCmd))

	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Conf.Llm.ApiKey))

	client := &http.Client{}
	if config.Conf.App.Proxy != "" {
		transport := &http.Transport{
			Proxy: http.ProxyURL(config.Conf.App.ParsedProxy),
		}
		client.Transport = transport
		log.GetLogger().Info("using proxy", zap.String("proxy", config.Conf.App.Proxy))
	}

	resp, err := client.Do(req)
	if err != nil {
		log.GetLogger().Error("http request failed", zap.Error(err))
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.GetLogger().Info("response status", zap.Int("status_code", resp.StatusCode))
	log.GetLogger().Info("response body", zap.String("body", string(body)))

	if resp.StatusCode != http.StatusOK {
		log.GetLogger().Error("non-200 status code", zap.Int("status_code", resp.StatusCode), zap.String("body", string(body)))
		return "", fmt.Errorf("non-200 status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	if len(result.Choices) == 0 {
		log.GetLogger().Error("no choices in response")
		return "", fmt.Errorf("no choices in response")
	}

	return result.Choices[0].Message.Content, nil
}

func (c *Client) Text2Speech(text, voice string, outputFile string) error {
	baseUrl := config.Conf.Tts.Openai.BaseUrl
	if baseUrl == "" {
		baseUrl = "https://api.openai.com/v1"
	}
	url := baseUrl + "/audio/speech"

	// 创建HTTP请求
	reqBody := fmt.Sprintf(`{
		"model": "tts-1",
		"input": "%s",
		"voice":"%s",
		"response_format": "wav"
	}`, text, voice)
	req, err := http.NewRequest("POST", url, strings.NewReader(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Conf.Tts.Openai.ApiKey))

	// 发送HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.GetLogger().Error("openai tts failed", zap.Int("status_code", resp.StatusCode), zap.String("body", string(body)))
		return fmt.Errorf("openai tts none-200 status code: %d", resp.StatusCode)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func parseJSONResponse(jsonStr string) (string, error) {
	var response struct {
		Translations []struct {
			Original   string `json:"original_sentence"`
			Translated string `json:"translated_sentence"`
		} `json:"translations"`
	}

	err := json.Unmarshal([]byte(jsonStr), &response)
	if err != nil {
		return "", fmt.Errorf("failed to parse JSON: %v", err)
	}

	var result strings.Builder
	for i, item := range response.Translations {
		result.WriteString(fmt.Sprintf("%d\n%s\n%s\n\n",
			i+1,
			item.Translated,
			item.Original))
	}

	return result.String(), nil
}
