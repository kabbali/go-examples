package http_calls

import (
	"errors"
	"github.com/kabbali/go-httpclient/gohttp"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	gohttp.StartMockServer()
	os.Exit(m.Run())
}

func TestGetEndpointsGetFromApiError(t *testing.T) {
	// Initialization
	gohttp.FlushMocks()
	gohttp.AddMock(gohttp.Mock{
		Method: http.MethodGet,
		Url:    "https://api.github.com",
		Error:  errors.New("timeout getting response from api"),
	})
	// Execution
	endpoints, err := GetEndpoints()

	// Validation
	assert.Nil(t, endpoints)
	assert.NotNil(t, err)
	assert.EqualValues(t, "timeout getting response from api", err.Error())
}

func TestGetEndpointsNotFoundError(t *testing.T) {
	// Initialization
	gohttp.FlushMocks()
	gohttp.AddMock(gohttp.Mock{
		Method:             http.MethodGet,
		Url:                "https://api.github.com",
		RequestBody:        "",
		Error:              nil,
		ResponseBody:       `{"message": "https://api.github.com not found"}`,
		ResponseStatusCode: http.StatusNotFound,
	})
	// Execution
	endpoints, err := GetEndpoints()

	// Validation
	assert.Nil(t, endpoints)
	assert.NotNil(t, err)
	assert.EqualValues(t, "error when trying to fetch \"https://api.github.com\"", err.Error())

}

func TestGetEndpointsUnmarshalJsonError(t *testing.T) {
	// Initialization
	gohttp.FlushMocks()
	gohttp.AddMock(gohttp.Mock{
		Method:             http.MethodGet,
		Url:                "https://api.github.com",
		RequestBody:        "",
		Error:              nil,
		ResponseBody:       `{"events_url": "https://api.github.com/events"`,
		ResponseStatusCode: http.StatusOK,
	})
	// Execution
	endpoints, err := GetEndpoints()

	// Validation
	assert.Nil(t, endpoints)
	assert.NotNil(t, err)
	assert.EqualValues(t, "unexpected end of JSON input", err.Error())
}

func TestGetEndpointsNoError(t *testing.T) {
	// Initialization
	gohttp.FlushMocks()
	gohttp.AddMock(gohttp.Mock{
		Method:             http.MethodGet,
		Url:                "https://api.github.com",
		RequestBody:        "",
		Error:              nil,
		ResponseBody:       `{"events_url": "https://api.github.com/events"}`,
		ResponseStatusCode: http.StatusOK,
	})
	// Execution
	endpoints, err := GetEndpoints()

	// Validation
	assert.Nil(t, err)
	assert.NotNil(t, endpoints)
	assert.EqualValues(t, "https://api.github.com/events", endpoints.EventsUrl)
}
