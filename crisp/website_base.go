// Copyright 2017 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/url"
)


// WebsiteData mapping
type WebsiteData struct {
  Data  *Website  `json:"data,omitempty"`
}

// Website mapping
type Website struct {
  WebsiteID  *string  `json:"website_id,omitempty"`
  Name       *string  `json:"name,omitempty"`
  Domain     *string  `json:"domain,omitempty"`
  Logo       *string  `json:"logo,omitempty"`
}

// WebsiteCreate mapping
type WebsiteCreate struct {
  Name    string  `json:"name,omitempty"`
  Domain  string  `json:"domain,omitempty"`
}


// String returns the string representation of Website
func (instance Website) String() string {
  return Stringify(instance)
}


// CheckWebsiteExists checks if given website exists (by domain).
func (service *WebsiteService) CheckWebsiteExists(domain string) (*Response, error) {
  url := fmt.Sprintf("website?domain=%s", url.QueryEscape(domain))
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// CreateWebsite creates a new website.
func (service *WebsiteService) CreateWebsite(websiteData WebsiteCreate) (*Website, *Response, error) {
  url := "website"
  req, _ := service.client.NewRequest("POST", url, websiteData)

  website := new(WebsiteData)
  resp, err := service.client.Do(req, website)
  if err != nil {
    return nil, resp, err
  }

  return website.Data, resp, err
}


// GetWebsite resolves an existing website information.
func (service *WebsiteService) GetWebsite(websiteID string) (*Website, *Response, error) {
  url := fmt.Sprintf("website/%s", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  website := new(WebsiteData)
  resp, err := service.client.Do(req, website)
  if err != nil {
    return nil, resp, err
  }

  return website.Data, resp, err
}


// DeleteWebsite deletes an existing website.
func (service *WebsiteService) DeleteWebsite(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s", websiteID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}
