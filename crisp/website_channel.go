// Copyright 2018 Crisp IM SARL All rights reserved.
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
  Email  *string  `json:"email,omitempty"`
}

// WebsiteChannelEmailDomainData mapping
type WebsiteChannelEmailDomainData struct {
  Data  *WebsiteChannelEmailDomain  `json:"data,omitempty"`
}

// WebsiteChannelEmailDomain mapping
type WebsiteChannelEmailDomain struct {
  Root      *string  `json:"root,omitempty"`
  Basic     *string  `json:"basic,omitempty"`
  Custom    *string  `json:"custom,omitempty"`
  Verified  *bool    `json:"verified,omitempty"`
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
  Query  *string  `json:"query,omitempty"`
  Name   *string  `json:"name,omitempty"`
  Value  *string  `json:"value,omitempty"`
}

// WebsiteChannelEmailRelayData mapping
type WebsiteChannelEmailRelayData struct {
  Data  *WebsiteChannelEmailRelay  `json:"data,omitempty"`
}

// WebsiteChannelEmailRelay mapping
type WebsiteChannelEmailRelay struct {
  Status  *string                        `json:"status,omitempty"`
  SMTP    *WebsiteChannelEmailRelaySMTP  `json:"smtp,omitempty"`
}

// WebsiteChannelEmailRelaySMTP mapping
type WebsiteChannelEmailRelaySMTP struct {
  Username  *string  `json:"username,omitempty"`
  Server    *string  `json:"server,omitempty"`
  Port      *uint16  `json:"port,omitempty"`
}

// WebsiteChannelEmailDomainChange mapping
type WebsiteChannelEmailDomainChange struct {
  Basic   string  `json:"basic,omitempty"`
  Custom  string  `json:"custom,omitempty"`
}

// WebsiteChannelEmailRelayRequest mapping
type WebsiteChannelEmailRelayRequest struct {
  SMTP  *WebsiteChannelEmailRelayRequestSMTP  `json:"smtp"`
}

// WebsiteChannelEmailRelayRequestSMTP mapping
type WebsiteChannelEmailRelayRequestSMTP struct {
  Username  *string  `json:"username"`
  Password  *string  `json:"password"`
  Server    *string  `json:"server"`
  Port      *uint16  `json:"port"`
}


// String returns the string representation of WebsiteChannelEmail
func (instance WebsiteChannelEmail) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteChannelEmailSetupFlow
func (instance WebsiteChannelEmailSetupFlow) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteChannelEmailRelay
func (instance WebsiteChannelEmailRelay) String() string {
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


// GetWebsiteEmailChannelDomain resolves domain for website email channel.
func (service *WebsiteService) GetWebsiteEmailChannelDomain(websiteID string) (*WebsiteChannelEmailDomain, *Response, error) {
  url := fmt.Sprintf("website/%s/channel/email/domain", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  domain := new(WebsiteChannelEmailDomainData)
  resp, err := service.client.Do(req, domain)
  if err != nil {
    return nil, resp, err
  }

  return domain.Data, resp, err
}


// RequestWebsiteEmailChannelDomainChange requests a change in the email channel domain used to send and receive emails.
func (service *WebsiteService) RequestWebsiteEmailChannelDomainChange(websiteID string, domain WebsiteChannelEmailDomainChange) (*Response, error) {
  url := fmt.Sprintf("website/%s/channel/email/domain", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, domain)

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


// GetWebsiteEmailChannelRelay resolves the website email channel relay store.
func (service *WebsiteService) GetWebsiteEmailChannelRelay(websiteID string) (*WebsiteChannelEmailRelay, *Response, error) {
  url := fmt.Sprintf("website/%s/channel/email/relay", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  channel := new(WebsiteChannelEmailRelayData)
  resp, err := service.client.Do(req, channel)
  if err != nil {
    return nil, resp, err
  }

  return channel.Data, resp, err
}


// RequestWebsiteEmailChannelRelayChange requests a change in the email channel relay store.
func (service *WebsiteService) RequestWebsiteEmailChannelRelayChange(websiteID string, relay WebsiteChannelEmailRelayRequest) (*Response, error) {
  url := fmt.Sprintf("website/%s/channel/email/relay", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, relay)

  return service.client.Do(req, nil)
}
