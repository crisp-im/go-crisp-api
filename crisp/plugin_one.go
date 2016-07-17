// Copyright 2016 The go-crisp-api AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp

import (
  "fmt"
)

// PluginService service
type PluginService service

// PluginInformationData mapping
type PluginInformationData struct {
  Data  *PluginInformation  `json:"data"`
}

// PluginInformation mapping
type PluginInformation struct {
  ID           *string    `json:"id"`
  URN          *string    `json:"urn"`
  Type         *string    `json:"type"`
  Name         *string    `json:"name"`
  Description  *string    `json:"description"`
  Features     *[]string  `json:"features"`
  Showcase     *[]string  `json:"showcase"`
  Price        *int       `json:"price"`
  Color        *string    `json:"color"`
  Icon         *string    `json:"icon"`
  Banner       *string    `json:"banner"`
  Since        *string    `json:"since"`
}

// PluginStarsData mapping
type PluginStarsData struct {
  Data  *PluginStarsObject  `json:"data"`
}

// PluginStarsObject mapping
type PluginStarsObject struct {
  Object  *PluginStars  `json:"stars"`
}

// PluginStars mapping
type PluginStars struct {
  Mean   *int  `json:"mean"`
  Total  *int  `json:"total"`
}

// GetPluginInformation gets information about a plugin
// Reference: https://docs.crisp.im/api/v1/#plugin-one-plugin-get
func (service *PluginService) GetPluginInformation(pluginID string) (*PluginInformation, *Response, error) {
  url := fmt.Sprintf("plugin/%s", pluginID)
  req, err := service.client.NewRequest("GET", url, nil)

  plugin := new(PluginInformationData)
  resp, err := service.client.Do(req, plugin)
  if err != nil {
    return nil, resp, err
  }

  return plugin.Data, resp, err
}

// GetPluginStars gets stars for a plugin
// Reference: https://docs.crisp.im/api/v1/#plugin-one-plugin-get
func (service *PluginService) GetPluginStars(pluginID string) (*PluginStars, *Response, error) {
  url := fmt.Sprintf("plugin/%s/stars", pluginID)
  req, err := service.client.NewRequest("GET", url, nil)

  stars := new(PluginStarsData)
  resp, err := service.client.Do(req, stars)
  if err != nil {
    return nil, resp, err
  }

  return stars.Data.Object, resp, err
}
