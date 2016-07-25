// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// ResolveAllConversations resolves all conversations in website.
func (service *WebsiteService) ResolveAllConversations(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/batch/resolve", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, nil)

  return service.client.Do(req, nil)
}


// ReadAllConversations marks all conversations as read in website.
func (service *WebsiteService) ReadAllConversations(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/batch/read", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, nil)

  return service.client.Do(req, nil)
}
