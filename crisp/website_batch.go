// Copyright 2018 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteBatchConversationsOperation mapping
type WebsiteBatchConversationsOperation struct {
  Sessions  *[]string  `json:"sessions,omitempty"`
}

// WebsiteBatchPeopleOperation mapping
type WebsiteBatchPeopleOperation struct {
  People  *WebsiteBatchPeopleOperationInner  `json:"people,omitempty"`
}

// WebsiteBatchPeopleOperationInner mapping
type WebsiteBatchPeopleOperationInner struct {
  Profiles  *[]string                                `json:"profiles,omitempty"`
  Search    *WebsiteBatchPeopleOperationInnerSearch  `json:"search,omitempty"`
}

// WebsiteBatchPeopleOperationInnerSearch mapping
type WebsiteBatchPeopleOperationInnerSearch struct {
  Filter    []WebsiteFilter  `json:"filter,omitempty"`
  Operator  string           `json:"operator,omitempty"`
}


// BatchResolveConversations resolves given (or all) items in website (conversation variant).
func (service *WebsiteService) BatchResolveConversations(websiteID string, sessions []string) (*Response, error) {
  url := fmt.Sprintf("website/%s/batch/resolve", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, WebsiteBatchConversationsOperation{Sessions: &sessions})

  return service.client.Do(req, nil)
}


// BatchReadConversations marks given (or all) items as read in website (conversation variant).
func (service *WebsiteService) BatchReadConversations(websiteID string, sessions []string) (*Response, error) {
  url := fmt.Sprintf("website/%s/batch/read", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, WebsiteBatchConversationsOperation{Sessions: &sessions})

  return service.client.Do(req, nil)
}


// BatchRemoveConversations removes given items in website (conversation variant).
func (service *WebsiteService) BatchRemoveConversations(websiteID string, sessions []string) (*Response, error) {
  url := fmt.Sprintf("website/%s/batch/remove", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, WebsiteBatchConversationsOperation{Sessions: &sessions})

  return service.client.Do(req, nil)
}


// BatchRemovePeople removes given items in website (people variant).
func (service *WebsiteService) BatchRemovePeople(websiteID string, people WebsiteBatchPeopleOperationInner) (*Response, error) {
  url := fmt.Sprintf("website/%s/batch/remove", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, WebsiteBatchPeopleOperation{People: &people})

  return service.client.Do(req, nil)
}
