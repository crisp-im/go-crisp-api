// Copyright 2021 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteVerifySettingsData mapping
type WebsiteVerifySettingsData struct {
  Data  *WebsiteVerifySettings  `json:"data,omitempty"`
}

// WebsiteVerifyKeyData mapping
type WebsiteVerifyKeyData struct {
  Data  *WebsiteVerifyKey  `json:"data,omitempty"`
}

// WebsiteVerifySettings mapping
type WebsiteVerifySettings struct {
  Enabled  *bool  `json:"enabled,omitempty"`
}

// WebsiteVerifyKey mapping
type WebsiteVerifyKey struct {
 	Secret  *string  `json:"secret,omitempty"`
}

// WebsiteVerifySettingsUpdate mapping
type WebsiteVerifySettingsUpdate struct {
  Enabled  bool  `json:"enabled,omitempty"`
}


// String returns the string representation of WebsiteVerifySettings
func (instance WebsiteVerifySettings) String() string {
  return Stringify(instance)
}


// String returns the string representation of WebsiteVerifyKey
func (instance WebsiteVerifyKey) String() string {
  return Stringify(instance)
}


// GetVerifySettings resolves verify settings.
func (service *WebsiteService) GetVerifySettings(websiteID string) (*WebsiteVerifySettings, *Response, error) {
  url := fmt.Sprintf("website/%s/verify/settings", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  settings := new(WebsiteVerifySettingsData)
  resp, err := service.client.Do(req, settings)
  if err != nil {
    return nil, resp, err
  }

  return settings.Data, resp, err
}


// UpdateVerifySettings updates verify settings.
func (service *WebsiteService) UpdateVerifySettings(websiteID string, settings WebsiteVerifySettingsUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/verify/settings", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, settings)

  return service.client.Do(req, nil)
}


// GetVerifyKey resolves verify key.
func (service *WebsiteService) GetVerifyKey(websiteID string) (*WebsiteVerifyKey, *Response, error) {
  url := fmt.Sprintf("website/%s/verify/key", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  key := new(WebsiteVerifyKeyData)
  resp, err := service.client.Do(req, key)
  if err != nil {
    return nil, resp, err
  }

  return key.Data, resp, err
}


// RollVerifyKey rolls verify key.
func (service *WebsiteService) RollVerifyKey(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/verify/key", websiteID)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}
