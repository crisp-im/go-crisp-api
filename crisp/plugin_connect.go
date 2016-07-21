// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// PluginConnectWebsitesData mapping
type PluginConnectWebsitesData struct {
  Data  *[]PluginConnectWebsites  `json:"data,omitempty"`
}

// PluginConnectWebsites mapping
type PluginConnectWebsites struct {
  WebsiteID  *string       `json:"website_id,omitempty"`
  Settings   *interface{}  `json:"settings,omitempty"`
}


// String returns the string representation of PluginConnectWebsites
func (instance PluginConnectWebsites) String() string {
  return Stringify(instance)
}


// CheckConnectSessionValidity checks whether the connected plugin session is valid or not.
// Reference: https://docs.crisp.im/api/v1/#plugin-plugin-connect-session-head
func (service *PluginService) CheckConnectSessionValidity() (*Response, error) {
  url := "plugin/connect/session"
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ListConnectWebsites lists the websites linked to connected plugin.
// Reference: https://docs.crisp.im/api/v1/#plugin-plugin-connect-websites-get
func (service *PluginService) ListConnectWebsites(pageNumber uint) (*[]PluginConnectWebsites, *Response, error) {
  url := fmt.Sprintf("plugin/connect/websites/%d", pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  websites := new(PluginConnectWebsitesData)
  resp, err := service.client.Do(req, websites)
  if err != nil {
    return nil, resp, err
  }

  return websites.Data, resp, err
}
