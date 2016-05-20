package heapanalytics

import (
	"encoding/json"
	// "log"
	"net/http"
	"net/http/httptest"
	// "reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestURLOptionWorks(t *testing.T) {
	appID := "test_appid"
	expectedURL := "http://testing.com"

	c := NewClient(appID, URL(expectedURL))
	assert.NotEqual(t, expectedURL, DefaultURL)
	assert.Equal(t, expectedURL, c.url)
}

func TestHttpClientOptionWorks(t *testing.T) {
	appID := "test_appid"
	expectedClient := &http.Client{Timeout: time.Duration(1 * time.Second)}

	c := NewClient(appID, HttpClient(expectedClient))
	assert.NotEqual(t, expectedClient, http.DefaultClient)
	assert.Equal(t, expectedClient, c.httpClient)
}

func TestNewClientSetsDefauls(t *testing.T) {
	expectedAppID := "test_appid"
	expectedHTTPClient := http.DefaultClient

	c := NewClient(expectedAppID)
	assert.Equal(t, expectedHTTPClient, c.httpClient)
	assert.Equal(t, expectedAppID, c.appId)
	assert.Equal(t, DefaultURL, c.url)
	assert.Equal(t, DefaultPathTrack, c.pathTrack)
	assert.Equal(t, DefaultPathUserProperties, c.pathUserProperties)
}

func TestTrackSendsCorrectRequest(t *testing.T) {
	type expectedBodyFormat struct {
		AppId      string                 `json:"app_id"`
		Identity   string                 `json:"identity"`
		Event      string                 `json:"event"`
		Properties map[string]interface{} `json:"properties,omitempty"`
	}

	expectedPath := "/api/track"
	expectedContentType := "application/json"
	expectedAppID := "test_appid"
	expectedIdentity := "test_identity"
	expectedEvent := "test_event"
	expectedProperties := map[string]interface{}{"TestString": "This value", "TestNumber": float64(10)}

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, expectedPath, r.URL.Path)
			assert.Equal(t, expectedContentType, r.Header.Get("Content-Type"))

			d := json.NewDecoder(r.Body)
			var resp expectedBodyFormat
			assert.NoError(t, d.Decode(&resp))

			assert.Equal(t, expectedAppID, resp.AppId)
			assert.Equal(t, expectedIdentity, resp.Identity)
			assert.Equal(t, expectedEvent, resp.Event)
			assert.Equal(t, expectedProperties, resp.Properties)
		}))

	defer ts.Close()

	c := NewClient(expectedAppID)
	c.url = ts.URL

	err := c.Track(expectedIdentity, expectedEvent, expectedProperties)
	assert.NoError(t, err)
}

func TestUserPropertiesSendsCorrectRequest(t *testing.T) {
	type expectedBodyFormat struct {
		AppId      string                 `json:"app_id"`
		Identity   string                 `json:"identity"`
		Properties map[string]interface{} `json:"properties,omitempty"`
	}

	expectedPath := "/api/add_user_properties"
	expectedContentType := "application/json"
	expectedAppID := "test_appid"
	expectedIdentity := "test_identity"
	expectedProperties := map[string]interface{}{"TestString": "This value", "TestNumber": float64(10)}

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, expectedPath, r.URL.Path)
			assert.Equal(t, expectedContentType, r.Header.Get("Content-Type"))

			d := json.NewDecoder(r.Body)
			var resp expectedBodyFormat
			assert.NoError(t, d.Decode(&resp))

			assert.Equal(t, expectedAppID, resp.AppId)
			assert.Equal(t, expectedIdentity, resp.Identity)
			assert.Equal(t, expectedProperties, resp.Properties)
		}))

	defer ts.Close()

	c := NewClient(expectedAppID)
	c.url = ts.URL

	err := c.UserProperties(expectedIdentity, expectedProperties)
	assert.NoError(t, err)
}