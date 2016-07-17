// Copyright 2016 The go-crisp-api AUTHORS. All rights reserved.
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
  apiVersion = "v2"
  defaultBaseURL = "https://api.crisp.im"
  userAgent = "go-crisp-api/" + libraryVersion
  acceptContentType = "application/json"
)

type Client struct {
  client *http.Client

  BaseURL *url.URL
  UserAgent string

  common service

  Plugin *PluginService
}

type service struct {
  client *Client
}

type Response struct {
  *http.Response
}

type ErrorResponse struct {
  Response *http.Response
}


// Error prints an error response
func (response *ErrorResponse) Error() string {
  return fmt.Sprintf("%v %v: %d %v %+v",
    response.Response.Request.Method, response.Response.Request.URL,
    response.Response.StatusCode)
}


// NewClient returns a new API client
func NewClient(httpClient *http.Client) *Client {
  if httpClient == nil {
    httpClient = http.DefaultClient
  }

  baseURL, _ := url.Parse(defaultBaseURL)

  client := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
  client.common.client = client

  client.Plugin = (*PluginService)(&client.common)

  return client
}


// NewRequest creates an API request
func (client *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
  rel, err := url.Parse(urlStr)
  if err != nil {
    return nil, err
  }

  u := client.BaseURL.ResolveReference(rel)

  var buf io.ReadWriter
  if body != nil {
    buf = new(bytes.Buffer)
    err := json.NewEncoder(buf).Encode(body)
    if err != nil {
      return nil, err
    }
  }

  req, err := http.NewRequest(method, u.String(), buf)
  if err != nil {
    return nil, err
  }

  req.Header.Add("Accept", acceptContentType)
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

  // TODO: map to segmented error returns w/ types / structs

  return errorResponse
}
