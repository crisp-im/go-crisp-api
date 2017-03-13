// Copyright 2016 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/url"
)


// PluginListData mapping
type PluginListData struct {
  Data  *[]PluginInformation  `json:"data,omitempty"`
}


// ListAllPlugins lists all available plugins on the marketplace.
func (service *PluginService) ListAllPlugins(pageNumber uint) (*[]PluginInformation, *Response, error) {
  url := fmt.Sprintf("plugins/list/all/%d", pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  plugins := new(PluginListData)
  resp, err := service.client.Do(req, plugins)
  if err != nil {
    return nil, resp, err
  }

  return plugins.Data, resp, err
}


// SearchPlugins searches for plugins in the marketplace, given a search term.
func (service *PluginService) SearchPlugins(query string, pageNumber uint) (*[]PluginInformation, *Response, error) {
  url := fmt.Sprintf("plugins/list/search/%d?query=%s", pageNumber, url.QueryEscape(query))
  req, _ := service.client.NewRequest("GET", url, nil)

  plugins := new(PluginListData)
  resp, err := service.client.Do(req, plugins)
  if err != nil {
    return nil, resp, err
  }

  return plugins.Data, resp, err
}
