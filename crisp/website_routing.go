// Copyright 2017 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteRoutingSettingsData mapping
type WebsiteRoutingSettingsData struct {
  Data  *WebsiteRoutingSettings  `json:"data,omitempty"`
}

// WebsiteRoutingSettings mapping
type WebsiteRoutingSettings struct {
  Assign          *bool  `json:"assign,omitempty"`
  ReprocessRules  *bool  `json:"reprocess_rules,omitempty"`
}

// WebsiteRoutingRulesData mapping
type WebsiteRoutingRulesData struct {
  Data  *WebsiteRoutingRules  `json:"data,omitempty"`
}

// WebsiteRoutingRules mapping
type WebsiteRoutingRules struct {
  Rules  *[]WebsiteRoutingRulesItem  `json:"rules,omitempty"`
}

// WebsiteRoutingRulesItem mapping
type WebsiteRoutingRulesItem struct {
  Conditions  *[]UserFilter  `json:"conditions,omitempty"`
  Assign      *[]string      `json:"assign,omitempty"`
}

// WebsiteRoutingSettingsUpdate mapping
type WebsiteRoutingSettingsUpdate struct {
  Assign          bool  `json:"assign,omitempty"`
  ReprocessRules  bool  `json:"reprocess_rules,omitempty"`
}

// WebsiteRoutingRulesUpdate mapping
type WebsiteRoutingRulesUpdate struct {
  Conditions  *[]UserFilter  `json:"conditions,omitempty"`
  Assign      []string       `json:"assign,omitempty"`
}


// String returns the string representation of WebsiteRoutingSettings
func (instance WebsiteRoutingSettings) String() string {
  return Stringify(instance)
}


// String returns the string representation of WebsiteRoutingRules
func (instance WebsiteRoutingRules) String() string {
  return Stringify(instance)
}


// GetWebsiteRoutingSettings resolves the routing settings for a website.
func (service *WebsiteService) GetWebsiteRoutingSettings(websiteID string) (*WebsiteRoutingSettings, *Response, error) {
  url := fmt.Sprintf("website/%s/routing/settings", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  settings := new(WebsiteRoutingSettingsData)
  resp, err := service.client.Do(req, settings)
  if err != nil {
    return nil, resp, err
  }

  return settings.Data, resp, err
}


// UpdateWebsiteRoutingSettings updates the routing settings for a website.
func (service *WebsiteService) UpdateWebsiteRoutingSettings(websiteID string, settings WebsiteRoutingSettingsUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/routing/settings", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, settings)

  return service.client.Do(req, nil)
}


// ListWebsiteRoutingRules lists the routing rules for a website.
func (service *WebsiteService) ListWebsiteRoutingRules(websiteID string) (*WebsiteRoutingRules, *Response, error) {
  url := fmt.Sprintf("website/%s/routing/rules", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  rules := new(WebsiteRoutingRulesData)
  resp, err := service.client.Do(req, rules)
  if err != nil {
    return nil, resp, err
  }

  return rules.Data, resp, err
}


// SaveWebsiteRoutingRules saves the routing rules for a website.
func (service *WebsiteService) SaveWebsiteRoutingRules(websiteID string, rules WebsiteRoutingRulesUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/routing/rules", websiteID)
  req, _ := service.client.NewRequest("PUT", url, rules)

  return service.client.Do(req, nil)
}
