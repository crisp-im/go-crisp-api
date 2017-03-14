// Copyright 2016 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// PluginInformationData mapping
type PluginInformationData struct {
  Data  *PluginInformation  `json:"data,omitempty"`
}

// PluginInformation mapping
type PluginInformation struct {
  ID            *string                   `json:"id,omitempty"`
  URN           *string                   `json:"urn,omitempty"`
  Type          *string                   `json:"type,omitempty"`
  Name          *string                   `json:"name,omitempty"`
  Description   *string                   `json:"description,omitempty"`
  Features      *[]string                 `json:"features,omitempty"`
  Showcase      *[]string                 `json:"showcase,omitempty"`
  Price         *uint                     `json:"price,omitempty"`
  Plans         *[]PluginInformationPlan  `json:"plans,omitempty"`
  Color         *string                   `json:"color,omitempty"`
  Icon          *string                   `json:"icon,omitempty"`
  Banner        *string                   `json:"banner,omitempty"`
  Configurable  *bool                     `json:"configurable,omitempty"`
  Since         *string                   `json:"since,omitempty"`
}

// PluginInformationPlan mapping
type PluginInformationPlan struct {
  ID     *string  `json:"id,omitempty"`
  Name   *string  `json:"name,omitempty"`
  Price  *uint    `json:"price,omitempty"`
}


// String returns the string representation of PluginInformation
func (instance PluginInformation) String() string {
  return Stringify(instance)
}


// GetPluginInformation resolves plugin information.
func (service *PluginService) GetPluginInformation(pluginID string) (*PluginInformation, *Response, error) {
  url := fmt.Sprintf("plugin/%s", pluginID)
  req, _ := service.client.NewRequest("GET", url, nil)

  plugin := new(PluginInformationData)
  resp, err := service.client.Do(req, plugin)
  if err != nil {
    return nil, resp, err
  }

  return plugin.Data, resp, err
}
