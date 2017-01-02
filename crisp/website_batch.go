// Copyright 2016 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteBatchOperation mapping
type WebsiteBatchOperation struct {
  Sessions  *[]string  `json:"sessions,omitempty"`
}


// ResolveGivenConversations resolves given (or all) conversations in website.
func (service *WebsiteService) ResolveGivenConversations(websiteID string, sessions []string) (*Response, error) {
  url := fmt.Sprintf("website/%s/batch/resolve", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, WebsiteBatchOperation{Sessions: &sessions})

  return service.client.Do(req, nil)
}


// ReadGivenConversations marks given (or all) conversations as read in website.
func (service *WebsiteService) ReadGivenConversations(websiteID string, sessions []string) (*Response, error) {
  url := fmt.Sprintf("website/%s/batch/read", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, WebsiteBatchOperation{Sessions: &sessions})

  return service.client.Do(req, nil)
}


// RemoveGivenConversations removes given conversations in website.
func (service *WebsiteService) RemoveGivenConversations(websiteID string, sessions []string) (*Response, error) {
  url := fmt.Sprintf("website/%s/batch/remove", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, WebsiteBatchOperation{Sessions: &sessions})

  return service.client.Do(req, nil)
}
