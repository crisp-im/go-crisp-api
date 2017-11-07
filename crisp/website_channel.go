// Copyright 2017 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/url"
)


// WebsiteChannelEmailData mapping
type WebsiteChannelEmailData struct {
  Data  *WebsiteChannelEmail  `json:"data,omitempty"`
}

// WebsiteChannelEmail mapping
type WebsiteChannelEmail struct {
  Domain  *string  `json:"domain,omitempty"`
  Email   *string  `json:"email,omitempty"`
}

// WebsiteChannelEmailSetupFlowData mapping
type WebsiteChannelEmailSetupFlowData struct {
  Data  *WebsiteChannelEmailSetupFlow  `json:"data,omitempty"`
}

// WebsiteChannelEmailSetupFlow mapping
type WebsiteChannelEmailSetupFlow struct {
  Domain  *string                             `json:"domain,omitempty"`
  Setup   *WebsiteChannelEmailSetupFlowSetup  `json:"setup,omitempty"`
}

// WebsiteChannelEmailSetupFlowSetup mapping
type WebsiteChannelEmailSetupFlowSetup struct {
  Records  *[]WebsiteChannelEmailSetupFlowSetupRecord  `json:"records,omitempty"`
}

// WebsiteChannelEmailSetupFlowSetupRecord mapping
type WebsiteChannelEmailSetupFlowSetupRecord struct {
  Type   *string  `json:"type,omitempty"`
  Name   *string  `json:"name,omitempty"`
  Value  *string  `json:"value,omitempty"`
}

// WebsiteChannelEmailRequest mapping
type WebsiteChannelEmailRequest struct {
  Domain  *string  `json:"domain,omitempty"`
}


// String returns the string representation of WebsiteChannelEmail
func (instance WebsiteChannelEmail) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteChannelEmailSetupFlow
func (instance WebsiteChannelEmailSetupFlow) String() string {
  return Stringify(instance)
}


// GetWebsiteEmailChannel resolves the website email channel value.
func (service *WebsiteService) GetWebsiteEmailChannel(websiteID string) (*WebsiteChannelEmail, *Response, error) {
  url := fmt.Sprintf("website/%s/channel/email", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  channel := new(WebsiteChannelEmailData)
  resp, err := service.client.Do(req, channel)
  if err != nil {
    return nil, resp, err
  }

  return channel.Data, resp, err
}


// RequestWebsiteEmailChannelChange requests a change in the email channel domain used to send and receive emails.
func (service *WebsiteService) RequestWebsiteEmailChannelChange(websiteID string, domain string) (*Response, error) {
  url := fmt.Sprintf("website/%s/channel/email", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, WebsiteChannelEmailRequest{Domain: &domain})

  return service.client.Do(req, nil)
}


// GenerateWebsiteEmailChannelSetupFlow retrieves the email channel setup flow.
func (service *WebsiteService) GenerateWebsiteEmailChannelSetupFlow(websiteID string, domain string) (*WebsiteChannelEmailSetupFlow, *Response, error) {
  url := fmt.Sprintf("website/%s/channel/email/setup?domain=%s", websiteID, url.QueryEscape(domain))
  req, _ := service.client.NewRequest("GET", url, nil)

  setup := new(WebsiteChannelEmailSetupFlowData)
  resp, err := service.client.Do(req, setup)
  if err != nil {
    return nil, resp, err
  }

  return setup.Data, resp, err
}
