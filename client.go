package heapanalytics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// DefaultHost is the default host to use
	DefaultURL = "https://heapanalytics.com"
	//DefaultPathTrack is the default path to use for track events
	DefaultPathTrack = "/api/track"
	//DefaultPathUserProperties is the default path to use for add user properties
	DefaultPathUserProperties = "/api/add_user_properties"
	//ContentType is the type of data to send the API
	ContentType = "application/json"
)

// ClientOption defines the format of an optional parameter for NewClient
type ClientOption func(*Client)

// Host is an ClientOption that can be passed into NewClient
// to change the hostname from default
func URL(url string) ClientOption {
	return func(c *Client) {
		c.url = url
	}
}

// HttpClient is a ClientOption that can be passed into NewClient
// to change the httpclient used from default
func HttpClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

// Client represents a client for the heap API
type Client struct {
	appId              string
	httpClient         *http.Client
	url                string
	pathTrack          string
	pathUserProperties string
}

// NewClient returns a pointer to a new API client
func NewClient(appID string, options ...ClientOption) *Client {
	// create a client with default settings
	client := &Client{
		appId:              appID,
		httpClient:         http.DefaultClient,
		url:                DefaultURL,
		pathTrack:          DefaultPathTrack,
		pathUserProperties: DefaultPathUserProperties,
	}

	// apply any user configuration if applicable
	for _, option := range options {
		option(client)
	}

	return client
}

// Track posts an API request to the "track" API
func (c *Client) Track(identity, event string, properties map[string]interface{}) error {
	e := NewEvent(c.appId, identity, event, properties)
	return c.send(e, c.pathTrack)
}

// Identify posts an API request to the "user properties" API
func (c *Client) UserProperties(identity string, properties map[string]interface{}) error {
	e := NewEvent(c.appId, identity, "", properties)
	return c.send(e, c.pathUserProperties)
}

// send sends POSTS to the api at heap with event e at path p
func (c *Client) send(e *Event, p string) error {
	b, err := json.Marshal(e)
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)
	resp, err := c.httpClient.Post(
		c.url+p,
		ContentType,
		r,
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		msg, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Expected 200 (OK) but got %d. Could not parse response body: %s", resp.StatusCode, err.Error())
		}
		return fmt.Errorf("Expected 200 (OK) but got %d (%s)", resp.StatusCode, msg)

	} else {
		return err
	}
}
