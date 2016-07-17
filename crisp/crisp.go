// Copyright 2016 The go-crisp-api AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp

import (
  "encoding/json"
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

type ErrorResponse struct {
  Response *http.Response
}


// Returns a new API client
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


// Creates an API request
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


// Sends an API request
func (client *Client) Do(req *http.Request, v interface{}) (*Response, error) {
  rateLimitCategory := category(req.URL.Path)

  // If we've hit rate limit, don't make further requests before Reset time.
  if err := client.checkRateLimitBeforeDo(req, rateLimitCategory); err != nil {
    return nil, err
  }

  resp, err := client.client.Do(req)
  if err != nil {
    return nil, err
  }

  defer func() {
    // Drain up to 512 bytes and close the body to let the Transport reuse the connection
    io.CopyN(ioutil.Discard, resp.Body, 512)
    resp.Body.Close()
  }()

  response := newResponse(resp)

  client.rateMu.Lock()
  client.rateLimits[rateLimitCategory] = response.Rate
  client.mostRecent = rateLimitCategory
  client.rateMu.Unlock()

  err = CheckResponse(resp)
  if err != nil {
    // even though there was an error, we still return the response
    // in case the caller wants to inspect it further
    return response, err
  }

  if v != nil {
    if w, ok := v.(io.Writer); ok {
      io.Copy(w, resp.Body)
    } else {
      err = json.NewDecoder(resp.Body).Decode(v)
      if err == io.EOF {
        err = nil // ignore EOF errors caused by empty response body
      }
    }
  }

  return response, err
}


// Check response for errors
func CheckResponse(response *http.Response) error {
  if code := response.StatusCode; 200 <= code && code <= 299 {
    return nil
  }
  errorResponse := &ErrorResponse{Response: response}

  // TODO: map to segmented error returns w/ types / structs

  return errorResponse
}
