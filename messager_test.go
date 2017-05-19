package mattergo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	w := Webhook{
		URL:      "http://example.com/hooks/replaceme",
		Username: "bot",
		Channel:  "twitter",
	}

	m := Message{
		Webhook: w,
		Text:    "Hallo",
	}

	assert.Equal(t, m.URL, w.URL, "Message should contain same URL as webhook")
	assert.Equal(t, m.Username, w.Username, "Message should contain same Username as webhook")
	assert.Equal(t, m.Channel, w.Channel, "Message should contain same Channel as webhook")
}
