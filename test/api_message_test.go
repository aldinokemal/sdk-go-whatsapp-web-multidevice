/*
WhatsApp API MultiDevice

Testing MessageAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package SdkWhatsappWebMultiDevice

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
)

func Test_SdkWhatsappWebMultiDevice_MessageAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MessageAPIService DeleteMessage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var messageId string

		resp, httpRes, err := apiClient.MessageAPI.DeleteMessage(context.Background(), messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MessageAPIService ReactMessage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var messageId string

		resp, httpRes, err := apiClient.MessageAPI.ReactMessage(context.Background(), messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MessageAPIService ReadMessage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var messageId string

		resp, httpRes, err := apiClient.MessageAPI.ReadMessage(context.Background(), messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MessageAPIService RevokeMessage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var messageId string

		resp, httpRes, err := apiClient.MessageAPI.RevokeMessage(context.Background(), messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MessageAPIService UpdateMessage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var messageId string

		resp, httpRes, err := apiClient.MessageAPI.UpdateMessage(context.Background(), messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
