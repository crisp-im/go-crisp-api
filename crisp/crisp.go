// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "bytes"
  "encoding/json"
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "net/url"
)


const (
  libraryVersion = "1.0.0"
  defaultEndpointURL = "https://api.crisp.im/v1/"
  userAgent = "go-crisp-api/" + libraryVersion
  acceptContentType = "application/json"
)

// BasicAuth maps basic auth credientials
type BasicAuth struct {
  Available bool
  Username string
  Password string
}

// Client maps an API client
type Client struct {
  client *http.Client
  basicAuth *BasicAuth

  BaseURL *url.URL
  UserAgent string

  common service

  Email   *EmailService
  Plugin  *PluginService
  User    *UserService
  Website *WebsiteService
}

type service struct {
  client *Client
}

// Response maps an API HTTP response
type Response struct {
  *http.Response
}

// ErrorResponse maps an API HTTP error response
type ErrorResponse struct {
  Response *http.Response
  Reason   string  `json:"reason"`
}


// Error prints an error response
func (response *ErrorResponse) Error() string {
  return fmt.Sprintf("%v %v: %d %v",
    response.Response.Request.Method, response.Response.Request.URL,
    response.Response.StatusCode, response.Reason)
}


// NewClient returns a new API client
func NewClient(httpClient *http.Client, endpointURL *string) *Client {
  targetEndpointURL := defaultEndpointURL

  if httpClient == nil {
    httpClient = http.DefaultClient
  }
  if endpointURL != nil {
    targetEndpointURL = *endpointURL
  }

  baseURL, _ := url.Parse(targetEndpointURL)

  client := &Client{client: httpClient, basicAuth: &BasicAuth{}, BaseURL: baseURL, UserAgent: userAgent}
  client.common.client = client

  client.Email = (*EmailService)(&client.common)
  client.Plugin = (*PluginService)(&client.common)
  client.User = (*UserService)(&client.common)
  client.Website = (*WebsiteService)(&client.common)

  return client
}


// Authenticate saves authentication parameters
func (client *Client) Authenticate(username string, password string) {
  client.basicAuth.Username = username
  client.basicAuth.Password = password
  client.basicAuth.Available = true
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

  if client.basicAuth.Available == true {
    req.SetBasicAuth(client.basicAuth.Username, client.basicAuth.Password)
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
  errorResponse := &ErrorResponse{Response: response}

  data, err := ioutil.ReadAll(response.Body)
  if err == nil && data != nil {
    json.Unmarshal(data, errorResponse)
  }

  return errorResponse
}
