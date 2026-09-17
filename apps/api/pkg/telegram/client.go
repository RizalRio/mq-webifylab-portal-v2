package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	token  string
	chatID string
	client *http.Client
}

func NewClient(token, chatID string) *Client {
	return &Client{
		token:  token,
		chatID: chatID,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type sendMessageReq struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func (c *Client) SendMessage(message string) error {
	if c.token == "" || c.chatID == "" {
		return fmt.Errorf("telegram token or chat ID is empty")
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", c.token)

	reqBody := sendMessageReq{
		ChatID:    c.chatID,
		Text:      message,
		ParseMode: "Markdown",
	}

	jsonValue, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code from telegram: %d", resp.StatusCode)
	}

	return nil
}
