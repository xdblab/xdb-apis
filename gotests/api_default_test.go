/*
XDB APIs

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package gotests

import (
	"github.com/stretchr/testify/assert"
	//openapiclient "github.com/xdblab/xdb-apis/goapi/xdbapi"
	openapiclient "github.com/xcherryio/apis/goapi/xcapi"
	"testing"
)

func Test_xdbapi_build(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	assert.NotNil(t, configuration)
	assert.NotNil(t, apiClient)
}
