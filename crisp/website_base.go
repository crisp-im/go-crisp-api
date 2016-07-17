// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteData mapping
type WebsiteData struct {
  Data  *Website  `json:"data,omitempty"`
}

// Website mapping
type Website struct {
  WebsiteID  *string  `json:"website_id,omitempty"`
}

// WebsiteCreate mapping
type WebsiteCreate struct {
  Name    *string  `json:"name,omitempty"`
  Domain  *string  `json:"domain,omitempty"`
}


// CreateWebsite creates a new website.
// Reference: https://docs.crisp.im/api/v1/#website-website-base-post
func (service *WebsiteService) CreateWebsite(name string, domain string) (*Website, *Response, error) {
  url := "website"
  req, _ := service.client.NewRequest("POST", url, WebsiteCreate{Name: &name, Domain: &domain})

  website := new(WebsiteData)
  resp, err := service.client.Do(req, website)
  if err != nil {
    return nil, resp, err
  }

  return website.Data, resp, err
}


// DeleteWebsite deletes an existing website.
// Reference: https://docs.crisp.im/api/v1/#website-website-base-delete
func (service *WebsiteService) DeleteWebsite(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s", websiteID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}
