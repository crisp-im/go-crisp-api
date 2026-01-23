// Copyright 2018 Crisp IM SAS All rights reserved.
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
  "errors"
  "io/ioutil"
  "net/http"
  "net/url"
)


const (
  libraryVersion = "3.72.1"
  defaultRestEndpointURL = "https://api.crisp.chat/v1/"
  userAgent = "go-crisp-api/" + libraryVersion
  acceptContentType = "application/json"
  clientTimeout = 10
  clientIdleConnTimeout = 45
  clientMaxIdleConns = 16
  clientMaxConnsPerHost = 64
  clientMaxIdleConnsPerHost = 4
  clientRequestRetryAttempts = 2
  clientRequestRetryHoldMillis = 1000
)

var errorDoAllAttemptsExhausted = errors.New("all request attempts were exhausted")
var errorDoAttemptNilRequest = errors.New("request could not be constructed")

// ClientConfig mapping
type ClientConfig struct {
  HTTPClient *http.Client
  HTTPHeaders *map[string]string
  RestEndpointURL string
}

type auth struct {
  Available bool
  Username string
  Password string
  Tier string
}

// Client maps an API client
type Client struct {
  client *http.Client
  headers *map[string]string
  auth *auth

  BaseURL *url.URL
  UserAgent string

  common service

  Bucket  *BucketService
  Media   *MediaService
  Plugin  *PluginService
  Plan    *PlanService
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
  Reason   string        `json:"reason,omitempty"`
  Data     *interface{}  `json:"data,omitempty"`
}


// Error prints an error response
func (response *errorResponse) Error() string {
  // Marshal error data to a JSON string (if any)
  var errorData []byte

  if response.Data != nil {
    errorData, _ = json.Marshal(*response.Data)
  } else {
    errorData = []byte("{}")
  }

  return fmt.Sprintf("HTTP %d %v\nPath → %v %v\nData → %s",
    response.Response.StatusCode, response.Reason,
    response.Response.Request.Method, response.Response.Request.URL, errorData)
}


// NewWithConfig returns a new API client
func NewWithConfig(config ClientConfig) *Client {
  // Defaults
  if config.HTTPClient == nil {
    // Build client transport
    clientTransport := http.DefaultTransport.(*http.Transport).Clone()
    clientTransport.IdleConnTimeout = time.Duration(clientIdleConnTimeout * time.Second)
    clientTransport.MaxIdleConns = clientMaxIdleConns
    clientTransport.MaxConnsPerHost = clientMaxConnsPerHost
    clientTransport.MaxIdleConnsPerHost = clientMaxIdleConnsPerHost

    // Create client
    config.HTTPClient = &http.Client{
      Timeout: time.Duration(clientTimeout * time.Second),
      Transport: clientTransport,
    }
  }
  if config.RestEndpointURL == "" {
    config.RestEndpointURL = defaultRestEndpointURL
  }

  // Create client
  baseURL, _ := url.Parse(config.RestEndpointURL)

  client := &Client{client: config.HTTPClient, headers: config.HTTPHeaders, auth: &auth{}, BaseURL: baseURL, UserAgent: userAgent}
  client.common.client = client

  // Map services
  client.Bucket = (*BucketService)(&client.common)
  client.Media = (*MediaService)(&client.common)
  client.Plugin = (*PluginService)(&client.common)
  client.Plan = (*PlanService)(&client.common)
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

  // Append custom headers? (if any)
  if client.headers != nil {
    for headerKey, headerValue := range *client.headers {
      req.Header.Add(headerKey, headerValue)
    }
  }

  // Append authorization header? (if authenticated)
  if client.auth.Available == true {
    req.SetBasicAuth(client.auth.Username, client.auth.Password)
    req.Header.Add("X-Crisp-Tier", client.auth.Tier)
  }

  // Append common headers
  req.Header.Add("Accept", acceptContentType)
  req.Header.Add("Content-Type", acceptContentType)

  // Append user agent header? (if any)
  if client.UserAgent != "" {
    req.Header.Add("User-Agent", client.UserAgent)
  }

  return req, nil
}


// Do sends an API request
func (client *Client) Do(req *http.Request, v interface{}) (*Response, error) {
  var lastErr error

  attempts := 0

  for attempts < clientRequestRetryAttempts {
    // Hold before this attempt? (ie. not first attempt)
    if attempts > 0 {
      time.Sleep(clientRequestRetryHoldMillis * time.Millisecond)
    }

    // Dispatch request attempt
    attempts++
    resp, shouldRetry, err := client.doAttempt(req, v)

    // Return response straight away? (we are done)
    if shouldRetry == false {
      return resp, err
    }

    // Should retry: store last error (we are not done)
    lastErr = err
  }

  // Set default error? (all attempts failed, but no error is set)
  if lastErr == nil {
    lastErr = errorDoAllAttemptsExhausted
  }

  // All attempts failed, return last attempt error
  return nil, lastErr
}


// doAttempt attempts an API request
func (client *Client) doAttempt(req *http.Request, v interface{}) (*Response, bool, error) {
  if req == nil {
    return nil, false, errorDoAttemptNilRequest
  }

  resp, err := client.client.Do(req)

  if checkRequestRetry(resp, err) == true {
    return nil, true, err
  }

  defer func() {
    io.CopyN(ioutil.Discard, resp.Body, 512)
    resp.Body.Close()
  }()

  response := newResponse(resp)

  err = checkResponse(resp)
  if err != nil {
    return response, false, err
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

  return response, false, err
}


// newResponse creates an HTTP response
func newResponse(httpResponse *http.Response) *Response {
  response := &Response{Response: httpResponse}

  return response
}


// checkRequestRetry checks if should retry request
func checkRequestRetry(response *http.Response, err error) bool {
  // Low-level error, or response status is a server error? (HTTP 5xx)
  if err != nil || response.StatusCode >= 500 {
    return true
  }

  // No low-level error (should not retry)
  return false
}


// checkResponse checks response for errors
func checkResponse(response *http.Response) error {
  // No error in response? (HTTP 2xx)
  if code := response.StatusCode; 200 <= code && code <= 299 {
    return nil
  }

  // Map response error data (eg. HTTP 4xx)
  errorResponse := &errorResponse{Response: response}

  data, err := ioutil.ReadAll(response.Body)
  if err == nil && data != nil {
    json.Unmarshal(data, errorResponse)
  }

  return errorResponse
}
