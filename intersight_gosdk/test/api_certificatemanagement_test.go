/*
Cisco Intersight

Testing CertificatemanagementApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package intersight

import (
	"context"
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_intersight_CertificatemanagementApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CertificatemanagementApiService CreateCertificatemanagementPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CertificatemanagementApi.CreateCertificatemanagementPolicy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificatemanagementApiService DeleteCertificatemanagementPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		httpRes, err := apiClient.CertificatemanagementApi.DeleteCertificatemanagementPolicy(context.Background(), moid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificatemanagementApiService GetCertificatemanagementPolicyByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.CertificatemanagementApi.GetCertificatemanagementPolicyByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificatemanagementApiService GetCertificatemanagementPolicyInventoryByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.CertificatemanagementApi.GetCertificatemanagementPolicyInventoryByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificatemanagementApiService GetCertificatemanagementPolicyInventoryList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CertificatemanagementApi.GetCertificatemanagementPolicyInventoryList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificatemanagementApiService GetCertificatemanagementPolicyList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CertificatemanagementApi.GetCertificatemanagementPolicyList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificatemanagementApiService PatchCertificatemanagementPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.CertificatemanagementApi.PatchCertificatemanagementPolicy(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificatemanagementApiService UpdateCertificatemanagementPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.CertificatemanagementApi.UpdateCertificatemanagementPolicy(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}