// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/url"
)


// PluginConnectAllWebsitesData mapping
type PluginConnectAllWebsitesData struct {
  Data  *[]PluginConnectAllWebsites  `json:"data,omitempty"`
}

// PluginConnectAllWebsites mapping
type PluginConnectAllWebsites struct {
  WebsiteID  *string       `json:"website_id,omitempty"`
  Settings   *interface{}  `json:"settings,omitempty"`
}

// PluginConnectWebsitesSinceData mapping
type PluginConnectWebsitesSinceData struct {
  Data  *[]PluginConnectWebsitesSince  `json:"data,omitempty"`
}

// PluginConnectWebsitesSince mapping
type PluginConnectWebsitesSince struct {
  WebsiteID   *string       `json:"website_id,omitempty"`
  Settings    *interface{}  `json:"settings,omitempty"`
  Difference  *string       `json:"difference,omitempty"`
}


// String returns the string representation of PluginConnectAllWebsites
func (instance PluginConnectAllWebsites) String() string {
  return Stringify(instance)
}


// String returns the string representation of PluginConnectWebsitesSince
func (instance PluginConnectWebsitesSince) String() string {
  return Stringify(instance)
}


// CheckConnectSessionValidity checks whether the connected plugin session is valid or not.
// Reference: https://docs.crisp.im/api/v1/#plugin-plugin-connect-session-head
func (service *PluginService) CheckConnectSessionValidity() (*Response, error) {
  url := "plugin/connect/session"
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ListAllConnectWebsites lists all websites linked to connected plugin.
// Reference: https://docs.crisp.im/api/v1/#plugin-plugin-connect-websites-get
func (service *PluginService) ListAllConnectWebsites(pageNumber uint) (*[]PluginConnectAllWebsites, *Response, error) {
  url := fmt.Sprintf("plugin/connect/websites/all/%d", pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  websites := new(PluginConnectAllWebsitesData)
  resp, err := service.client.Do(req, websites)
  if err != nil {
    return nil, resp, err
  }

  return websites.Data, resp, err
}


// ListConnectWebsitesSince lists the websites linked or unlinked or updated for connected plugin, since given date.
// Reference: https://docs.crisp.im/api/v1/#plugin-plugin-connect-websites-get-1
func (service *PluginService) ListConnectWebsitesSince(dateSince string) (*[]PluginConnectWebsitesSince, *Response, error) {
  url := fmt.Sprintf("plugin/connect/websites/since?date_since=%s", url.QueryEscape(dateSince))
  req, _ := service.client.NewRequest("GET", url, nil)

  websites := new(PluginConnectWebsitesSinceData)
  resp, err := service.client.Do(req, websites)
  if err != nil {
    return nil, resp, err
  }

  return websites.Data, resp, err
}
