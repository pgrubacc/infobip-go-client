package examples

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"infobip-go-client/pkg/infobip"
	"infobip-go-client/pkg/infobip/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// The following examples can also be used to test the client against a real environment.
// Replace the apiKey and baseURL params along with the From, To and Content fields of the message, then run the test.
func TestSendTextExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	whatsApp := client.WhatsApp()
	message := models.TextMessage{
		MessageCommon: models.MessageCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.TextContent{Text: "This message was sent from the Infobip API using the Go API client."},
	}

	msgResp, respDetails, err := whatsApp.SendTextMessage(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.MessageResponse{}, msgResp)
	fmt.Printf("%+v", msgResp)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
}

func TestSendDocumentExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	whatsApp := client.WhatsApp()
	message := models.DocumentMessage{
		MessageCommon: models.MessageCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.DocumentContent{MediaURL: "https://myurl.com/doc1.doc"},
	}

	msgResp, respDetails, err := whatsApp.SendDocumentMessage(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MessageResponse{}, msgResp)
}

func TestSendImageExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	whatsApp := client.WhatsApp()
	message := models.ImageMessage{
		MessageCommon: models.MessageCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.ImageContent{MediaURL: "https://thumbs.dreamstime.com/z/example-red-tag-example-red-square-price-tag-117502755.jpg"},
	}

	msgResp, respDetails, err := whatsApp.SendImageMessage(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MessageResponse{}, msgResp)
}

func TestAudioExample(t *testing.T) {
	apiKey := "secret"
	baseURL := "https://myinfobipurl.com"
	client, err := infobip.NewClient(baseURL, apiKey)
	whatsApp := client.WhatsApp()
	message := models.AudioMessage{
		MessageCommon: models.MessageCommon{
			From: "111111111111",
			To:   "222222222222",
		},
		Content: models.AudioContent{MediaURL: "https://dl.espressif.com/dl/audio/ff-16b-2c-44100hz.aac"},
	}

	msgResp, respDetails, err := whatsApp.SendAudioMessage(context.Background(), message)
	fmt.Printf("%+v\n", msgResp)

	require.Nil(t, err)
	assert.NotEqual(t, models.ResponseDetails{}, respDetails)
	assert.NotEqual(t, models.MessageResponse{}, msgResp)
}
