/*
WhatsApp API MultiDevice

Testing GroupApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sdk_go_whatsapp_web_multidevice

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
)

func Test_sdk_go_whatsapp_web_multidevice_GroupApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GroupApiService JoinGroupWithLink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GroupApi.JoinGroupWithLink(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
