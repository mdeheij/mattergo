package mattergo

import (
	"encoding/json"
	"net/http"
	"strings"
)

//Webhook is a definition of Mattermost's incoming Webhook
type Webhook struct {
	URL      string
	IconURL  string `json:"icon_url"`
	Username string `json:"username"`
	Channel  string `json:"channel"`
}

//Message is a chat message to be sent using a webhook
type Message struct {
	Webhook
	Text string `json:"text"`
}

//Send a message to a Mattermost chat channel
func (m *Message) Send() error {
	messageJSON, err := json.Marshal(m)
	if err != nil {
		return err
	}

	body := strings.NewReader(string(messageJSON))
	req, err := http.NewRequest("POST", m.URL, body)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}
