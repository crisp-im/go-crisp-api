// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


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
  Price        *uint      `json:"price"`
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
  Mean   *uint  `json:"mean"`
  Total  *uint  `json:"total"`
}

// PluginPersonalPluginRankData mapping
type PluginPersonalPluginRankData struct {
  Data  *PluginPersonalPluginRank  `json:"data"`
}

// PluginPersonalPluginRank mapping
type PluginPersonalPluginRank struct {
  Rank  *uint  `json:"rank"`
}


// GetPluginInformation resolves plugin information.
// Reference: https://docs.crisp.im/api/v1/#plugin-one-plugin-get
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


// GetPluginStars resolves plugin stars. This gives some stats about user rating of the plugin.
// Reference: https://docs.crisp.im/api/v1/#plugin-one-plugin-get-1
func (service *PluginService) GetPluginStars(pluginID string) (*PluginStars, *Response, error) {
  url := fmt.Sprintf("plugin/%s/stars", pluginID)
  req, _ := service.client.NewRequest("GET", url, nil)

  stars := new(PluginStarsData)
  resp, err := service.client.Do(req, stars)
  if err != nil {
    return nil, resp, err
  }

  return stars.Data.Object, resp, err
}


// GetPersonalPluginRank resolves our own ranking of the plugin (if we ever ranked it).
// Reference: https://docs.crisp.im/api/v1/#plugin-one-plugin-get-2
func (service *PluginService) GetPersonalPluginRank(pluginID string) (*PluginPersonalPluginRank, *Response, error) {
  url := fmt.Sprintf("plugin/%s/stars/self", pluginID)
  req, _ := service.client.NewRequest("GET", url, nil)

  stars := new(PluginPersonalPluginRankData)
  resp, err := service.client.Do(req, stars)
  if err != nil {
    return nil, resp, err
  }

  return stars.Data, resp, err
}


// RankPlugin ranks the plugin (as current user).
// Reference: https://docs.crisp.im/api/v1/#plugin-one-plugin-get-2
func (service *PluginService) RankPlugin(pluginID string, rank uint) (*Response, error) {
  url := fmt.Sprintf("plugin/%s/stars/self", pluginID)
  req, _ := service.client.NewRequest("PATCH", url, PluginPersonalPluginRank{&rank})

  return service.client.Do(req, nil)
}


// DeletePluginRank deletes personal rank of the plugin (as current user).
// Reference: https://docs.crisp.im/api/v1/#plugin-one-plugin-delete
func (service *PluginService) DeletePluginRank(pluginID string) (*Response, error) {
  url := fmt.Sprintf("plugin/%s/stars/self", pluginID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}
