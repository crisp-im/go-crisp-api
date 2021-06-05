// Copyright 2018 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "bytes"
  "encoding/json"
  "fmt"
  "time"
  "io"
  "io/ioutil"
  "net/http"
  "net/url"
)


const (
  libraryVersion = "3.10.15"
  defaultRestEndpointURL = "https://api.crisp.chat/v1/"
  defaultRealtimeEndpointURL = "https://app.relay.crisp.chat:443/"
  userAgent = "go-crisp-api/" + libraryVersion
  acceptContentType = "application/json"
  clientTimeout = 10
)

// ClientConfig mapping
type ClientConfig struct {
  HTTPClient *http.Client
  RestEndpointURL string
  RealtimeEndpointURL string
}

type auth struct {
  Available bool
  Username string
  Password string
  Tier string
}

// Client maps an API client
type Client struct {
  config *ClientConfig
  client *http.Client
  auth *auth

  BaseURL *url.URL
  UserAgent string

  common service

  Bucket  *BucketService
  Media   *MediaService
  Plugin  *PluginService
  Website *WebsiteService

  Events  *EventsService
}

type service struct {
  client *Client
}

// Response maps an API HTTP response
type Response struct {
  *http.Response
}

type errorResponse struct {
  Response *http.Response
  Reason   string  `json:"reason,omitempty"`
}


// Error prints an error response
func (response *errorResponse) Error() string {
  return fmt.Sprintf("%v %v: %d %v",
    response.Response.Request.Method, response.Response.Request.URL,
    response.Response.StatusCode, response.Reason)
}


// NewWithConfig returns a new API client
func NewWithConfig(config ClientConfig) *Client {
  // Defaults
  if config.HTTPClient == nil {
    config.HTTPClient = http.DefaultClient
    config.HTTPClient.Timeout = time.Duration(clientTimeout * time.Second)
  }
  if config.RestEndpointURL == "" {
    config.RestEndpointURL = defaultRestEndpointURL
  }
  if config.RealtimeEndpointURL == "" {
    config.RealtimeEndpointURL = defaultRealtimeEndpointURL
  }

  // Create client
  baseURL, _ := url.Parse(config.RestEndpointURL)

  client := &Client{config: &config, client: config.HTTPClient, auth: &auth{}, BaseURL: baseURL, UserAgent: userAgent}
  client.common.client = client

  // Map services
  client.Bucket = (*BucketService)(&client.common)
  client.Media = (*MediaService)(&client.common)
  client.Plugin = (*PluginService)(&client.common)
  client.Website = (*WebsiteService)(&client.common)
  client.Events = (*EventsService)(&client.common)

  return client
}


// New returns a new API client
func New() *Client {
  return NewWithConfig(ClientConfig{})
}


// AuthenticateTier saves authentication parameters for tier
func (client *Client) AuthenticateTier(tier string, username string, password string) {
  client.auth.Tier = tier
  client.auth.Username = username
  client.auth.Password = password
  client.auth.Available = true
}


// Authenticate saves authentication parameters for user (default)
func (client *Client) Authenticate(username string, password string) {
  client.AuthenticateTier("user", username, password)
}


// NewRequest creates an API request
func (client *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
  rel, err := url.Parse(urlStr)
  if err != nil {
    return nil, err
  }

  url := client.BaseURL.ResolveReference(rel)

  var buf io.ReadWriter
  if body != nil {
    buf = new(bytes.Buffer)
    err := json.NewEncoder(buf).Encode(body)
    if err != nil {
      return nil, err
    }
  }

  req, err := http.NewRequest(method, url.String(), buf)
  if err != nil {
    return nil, err
  }

  if client.auth.Available == true {
    req.SetBasicAuth(client.auth.Username, client.auth.Password)
    req.Header.Add("X-Crisp-Tier", client.auth.Tier)
  }

  req.Header.Add("Accept", acceptContentType)
  req.Header.Add("Content-Type", acceptContentType)

  if client.UserAgent != "" {
    req.Header.Add("User-Agent", client.UserAgent)
  }

  return req, nil
}


// Do sends an API request
func (client *Client) Do(req *http.Request, v interface{}) (*Response, error) {
  resp, err := client.client.Do(req)
  if err != nil {
    return nil, err
  }

  defer func() {
    io.CopyN(ioutil.Discard, resp.Body, 512)
    resp.Body.Close()
  }()

  response := newResponse(resp)

  err = CheckResponse(resp)
  if err != nil {
    return response, err
  }

  if v != nil {
    if w, ok := v.(io.Writer); ok {
      io.Copy(w, resp.Body)
    } else {
      err = json.NewDecoder(resp.Body).Decode(v)
      if err == io.EOF {
        err = nil
      }
    }
  }

  return response, err
}


// newResponse creates an HTTP response
func newResponse(httpResponse *http.Response) *Response {
  response := &Response{Response: httpResponse}

  return response
}


// CheckResponse checks response for errors
func CheckResponse(response *http.Response) error {
  if code := response.StatusCode; 200 <= code && code <= 299 {
    return nil
  }
  errorResponse := &errorResponse{Response: response}

  data, err := ioutil.ReadAll(response.Body)
  if err == nil && data != nil {
    json.Unmarshal(data, errorResponse)
  }

  return errorResponse
}
