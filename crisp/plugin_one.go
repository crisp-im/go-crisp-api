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
  ID           *string    `json:"id,omitempty"`
  URN          *string    `json:"urn,omitempty"`
  Type         *string    `json:"type,omitempty"`
  Name         *string    `json:"name,omitempty"`
  Description  *string    `json:"description,omitempty"`
  Features     *[]string  `json:"features,omitempty"`
  Showcase     *[]string  `json:"showcase,omitempty"`
  Price        *uint      `json:"price,omitempty"`
  Color        *string    `json:"color,omitempty"`
  Icon         *string    `json:"icon,omitempty"`
  Banner       *string    `json:"banner,omitempty"`
  Since        *string    `json:"since,omitempty"`
}

// PluginStarsData mapping
type PluginStarsData struct {
  Data  *PluginStarsObject  `json:"data,omitempty"`
}

// PluginStarsObject mapping
type PluginStarsObject struct {
  Object  *PluginStars  `json:"stars,omitempty"`
}

// PluginStars mapping
type PluginStars struct {
  Mean   *uint  `json:"mean,omitempty"`
  Total  *uint  `json:"total,omitempty"`
}

// PluginPersonalPluginRankData mapping
type PluginPersonalPluginRankData struct {
  Data  *PluginPersonalPluginRank  `json:"data,omitempty"`
}

// PluginPersonalPluginRank mapping
type PluginPersonalPluginRank struct {
  Rank  *uint  `json:"rank,omitempty"`
}


// String returns the string representation of PluginInformation
func (instance PluginInformation) String() string {
  return Stringify(instance)
}


// String returns the string representation of PluginStars
func (instance PluginStars) String() string {
  return Stringify(instance)
}


// String returns the string representation of PluginPersonalPluginRank
func (instance PluginPersonalPluginRank) String() string {
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


// GetPluginStars resolves plugin stars. This gives some stats about user rating of the plugin.
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
func (service *PluginService) RankPlugin(pluginID string, rank uint) (*Response, error) {
  url := fmt.Sprintf("plugin/%s/stars/self", pluginID)
  req, _ := service.client.NewRequest("PATCH", url, PluginPersonalPluginRank{&rank})

  return service.client.Do(req, nil)
}


// DeletePluginRank deletes personal rank of the plugin (as current user).
func (service *PluginService) DeletePluginRank(pluginID string) (*Response, error) {
  url := fmt.Sprintf("plugin/%s/stars/self", pluginID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}
